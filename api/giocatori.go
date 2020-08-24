package api

import (
	"../database"
	"../partita"
	"github.com/gorilla/mux"
	"net/http"
)

func GetGiocatoriHandler(w http.ResponseWriter, req *http.Request, db *database.SqliteItemsAdapter) {

	giocatori, err := db.GetGiocatori()
	if err != nil {
		setResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	setResponse(w, giocatori, http.StatusOK)
}

func AddGiocatoreHandler(w http.ResponseWriter, req *http.Request, db *database.SqliteItemsAdapter) {
	nome := mux.Vars(req)["nome"]
	if nome == "" {
		setResponse(w, "Non è arrivato nessun nome", http.StatusInternalServerError)
		return
	}

	err := db.AddGiocatore(nome)
	if err != nil {
		setResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := db.SelectGiocatore(nome, 0)
	if err != nil {
		setResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	setResponse(w, id, http.StatusOK)
}

func SelectGiocatoreHandler(w http.ResponseWriter, req *http.Request, db *database.SqliteItemsAdapter) {
	nome := mux.Vars(req)["nome"]
	if nome == "" {
		setResponse(w, "Non è arrivato nessun nome", http.StatusInternalServerError)
		return
	}

	id, err := db.SelectGiocatore(nome, 0)
	if err != nil {
		setResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	setResponse(w, id, http.StatusOK)
}

func iniziaPartita() {
	giocatori := partita.IniziaPartita()
	toccaA := partita.GetChiInizia()
	for i:=range connessioni{
		carte := partita.GetCarte(i)
		scriviA(i, IniziaPartitaResp{Azione: "iniziaPartita", Carte: carte, CurrentGiocatori: *giocatori, ToccaA: toccaA})
	}
	chiamabili := partita.GetChiamabili()
	broadcast(setChiamabiliResp{Azione: "setChiamabili", Chiamabili: chiamabili, ToccaA: toccaA, Chiamante: -1})
}

func sonoProntoHandler(mex messaggioDaClient){
	iscritti := partita.GiocatoriIscritti()
	scriviA(mex.Mittente, setRegistratiResp{Azione: "setRegistrati", Iscritti: iscritti})
	partita.SetGiocatorePronto(mex.Mittente)
	pronti := partita.GetGiocatoriPronti()
	for _,g := range pronti{
		if !g {
			return
		}
	}
	iniziaPartita()
}

func addBotsHandler() {
	iscritti := partita.GiocatoriIscritti()
	for i := iscritti; i<5; i++ {
		partita.IscriviBot(i)
	}
	iniziaPartita()
}

func AddBotHandler(w http.ResponseWriter, req *http.Request, db *database.SqliteItemsAdapter) {
	nome := mux.Vars(req)["nome"]
	if nome == "" {
		setResponse(w, "Non è arrivato nessun nome", http.StatusInternalServerError)
		return
	}

	id, err := db.SelectGiocatore(nome, 1)
	if err != nil {
		setResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	setResponse(w, id, http.StatusOK)
}
