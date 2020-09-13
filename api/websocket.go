package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var connessioni = map[int]*websocket.Conn{}

var mu sync.Mutex

func ascolta(conn *websocket.Conn) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Errore leggendo il messaggio: " + err.Error())
			return
		}

		var messaggio messaggioDaClient
		err = json.Unmarshal(p, &messaggio)
		if err != nil {
			log.Println("Messaggio indecifrabile: " + err.Error())
			return
		}

		switch messaggio.Azione {
		case "chiManca":
			chiMancaHandler(messaggio)
		case "sonoPronto":
			sonoProntoHandler(messaggio)
		case "addBots":
			addBotsHandler()
		case "chiamaValore":
			chiamaValoreHandler(messaggio)
		case "chiamaBot":
			chiamaBotHandler()
		case "chiamaSeme":
			chiamaSemeHandler(messaggio)
		case "giocaCarta":
			giocaCartaHandler(messaggio)
		case "giocaCartaBot":
			giocaCartaBotHandler(messaggio)
		case "iniziaNuovaMano":
			iniziaManoHandler()
		case "altroRound":
			altroRoundHandler(messaggio)
		case "altroRoundBot":
			idBot, _ := strconv.Atoi(messaggio.Params[0])
			altroRoundHandler(messaggioDaClient{Mittente: idBot})
		case "bastaCosì":
			bastaCosiHandler()
		}

	}
}

func registraWebSocket(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	connessioni[id] = ws

	go ascolta(ws)

	log.Println("Connesso col client " + idStr)
}

func scriviA(id int, messaggio interface{}) {
	if len(connessioni) < id+1 {
		log.Println("messaggio a giocatore non registrato: " + strconv.Itoa(id))
		return
	}
	if connessioni[id].RemoteAddr() == nil {
		return
	}
	data, err := json.Marshal(messaggio)
	if err != nil {
		log.Println("messaggio inconvertibile: " + err.Error())
		return
	}
	mu.Lock()
	defer mu.Unlock()
	if err := connessioni[id].WriteMessage(1, data); err != nil {
		log.Println("errore scrivendo messaggio a " + strconv.Itoa(id) + ": " + err.Error())
	}
}

func broadcast(messaggio interface{}) {
	for i := range connessioni {
		scriviA(i, messaggio)
	}
}
