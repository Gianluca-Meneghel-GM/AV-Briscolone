package api

import (
	"../database"
	"../models"
	"../partita"
	"../store"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

func FinisciPartitaHandler(w http.ResponseWriter, req *http.Request, db *database.SqliteItemsAdapter) {
	giocatori, err := db.GetCurrentGiocatori()
	if err != nil {
		setResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = db.SvuotaCurrentGiocatori()
	if err != nil {
		setResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	store.AzzeraNumeroGiocatori()
	partita.ResetRound()
	partita.ResetMano()

	setResponse(w, giocatori, http.StatusOK)
}

func chiMancaHandler(mex messaggioDaClient) {
	iscritti := partita.GiocatoriIscritti()
	broadcast(setRegistratiResp{Azione: "setRegistrati", Iscritti: iscritti})
}

func chiamaValoreHandler(mex messaggioDaClient) {
	giocatore := mex.Mittente
	valoreStr := mex.Params[0]
	partita.AumentaAChiToccaChiamare()
	if valoreStr == "passo" {
		inChiamata := partita.PassaChiamata(giocatore)
		if len(inChiamata) == 1 {
			partita.SetChiamante(inChiamata[0])
		}
	} else {
		valore, err := strconv.Atoi(valoreStr)
		if err != nil {
			log.Println("errore leggendo val chiamato: " + err.Error())
			return
		}
		partita.ChiamaValore(valore)
	}
	chiamabili := partita.GetChiamabili()
	toccaA := partita.GetAChiToccaChiamare()
	chiamante := partita.GetChiamante()
	valChiamato := partita.GetValChiamato()
	broadcast(setChiamabiliResp{Azione: "setChiamabili", Chiamabili: chiamabili, ToccaA: toccaA, Chiamante: chiamante, ValChiamato: valChiamato})
}

func chiamaBotHandler() {
	toccaA := partita.GetAChiToccaChiamare()
	chiamaValoreHandler(messaggioDaClient{Mittente: toccaA, Params: []string{"passo"}})
}

func chiamaSemeHandler(mex messaggioDaClient) {
	partita.SetChiamata(mex.Params[0])
	toccaA := partita.GetAChiTocca()
	broadcast(iniziaRoundResp{Azione: "iniziaRound", ToccaA: toccaA, Briscola: mex.Params[0]})
}

func giocaCartaHandler(mex messaggioDaClient) {
	valoreStr := mex.Params[0]
	valore, err := strconv.Atoi(valoreStr)
	seme := mex.Params[1]
	if err != nil {
		log.Println("errore leggendo carta giocata: " + err.Error())
		return
	}
	partita.GiocaCarta(seme, valore, mex.Mittente)
	broadcast(getGiocata())
}

func getGiocata() setGiocataResp {
	carteInMano := partita.GetCarteInMano()
	cartePrese := partita.GetCartePrese()
	toccaA := partita.GetAChiTocca()
	carteGiocate := partita.CarteGiocate()
	if len(carteGiocate) == 5 {
		toccaA = -1
	}
	mano := partita.ManiGiocate()
	return setGiocataResp{Azione: "setGiocata", CartePrese: cartePrese, CarteInMano: carteInMano, CarteGiocate: carteGiocate,
		ToccaA: toccaA, Mano: mano}
}

func giocaCartaBotHandler() {
	toccaA := partita.GetAChiTocca()
	time.Sleep(1000)
	partita.GiocaCartaBot(toccaA)
	broadcast(getGiocata())
}

func altroRoundHandler(mex messaggioDaClient) {
	if mex.Mittente == 0{
		partita.ResetRound()
	}
	partita.SetGiocatorePronto(mex.Mittente)
	pronti := partita.GetGiocatoriPronti()
	for _,g := range pronti{
		if !g {
			return
		}
	}
	iniziaPartita()
}

func bastaCosiHandler(){
	partita.SalvaRecordPartita()
}

func TryStartPartitaHandler(w http.ResponseWriter, req *http.Request, db *database.SqliteItemsAdapter) {
	idStr := mux.Vars(req)["id"]
	if idStr == "" {
		setResponse(w, "Nessun id ricevuto", http.StatusInternalServerError)
		return
	}

	iscritti := partita.GiocatoriIscritti()
	if iscritti < 5 {
		setResponse(w, iscritti, http.StatusOK)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		setResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	partita.ResetRound()
	partita.ResetMano()
	partita.FaiIlMazzo()
	carte := partita.GetCarte(id)
	for i := range carte {
		carte[i].SemeStr = string(carte[i].Seme)
	}
	giocatori, err := db.GetCurrentGiocatori()
	if err != nil {
		setResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	toccaA := partita.GetChiInizia()

	setResponse(w, models.PrimaMano{Carte: carte, CurrentGiocatori: *giocatori, ToccaA: toccaA}, http.StatusOK)
}

func ChiamaCartaHandler(w http.ResponseWriter, req *http.Request, db *database.SqliteItemsAdapter) {
	valoreStr := mux.Vars(req)["valore"]
	giocatoreStr := mux.Vars(req)["id"]
	giocatore, err := strconv.Atoi(giocatoreStr)
	if err != nil {
		setResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	partita.AumentaAChiToccaChiamare()
	if valoreStr == "passo" {
		inChiamata := partita.PassaChiamata(giocatore)
		if len(inChiamata) == 1 {
			partita.SetChiamante(inChiamata[0])
		}
	} else {
		valore, err := strconv.Atoi(valoreStr)
		if err != nil {
			setResponse(w, err.Error(), http.StatusInternalServerError)
			return
		}
		partita.ChiamaValore(valore)
	}

	setResponse(w, "OK", http.StatusOK)
}

func ChiamaSemeHandler(w http.ResponseWriter, req *http.Request, db *database.SqliteItemsAdapter) {
	seme := mux.Vars(req)["seme"]
	partita.SetChiamata(seme)
	setResponse(w, "OK", http.StatusOK)
}

func GiocaCartaHandler(w http.ResponseWriter, req *http.Request, db *database.SqliteItemsAdapter) {
	idStr := mux.Vars(req)["id"]
	id, err := strconv.Atoi(idStr)
	valoreStr := mux.Vars(req)["val"]
	valore, err := strconv.Atoi(valoreStr)
	seme := mux.Vars(req)["sem"]
	if err != nil {
		setResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	partita.GiocaCarta(seme, valore, id)

	setResponse(w, "OK", http.StatusOK)

}

func GiocaCartaBotHandler(w http.ResponseWriter, req *http.Request, db *database.SqliteItemsAdapter) {
	idStr := mux.Vars(req)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		setResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	partita.GiocaCartaBot(id)

	setResponse(w, "OK", http.StatusOK)

}

func ChiamaCartaBotHandler(w http.ResponseWriter, req *http.Request, db *database.SqliteItemsAdapter) {
	giocatoreStr := mux.Vars(req)["id"]
	giocatore, err := strconv.Atoi(giocatoreStr)
	if err != nil {
		setResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	partita.AumentaAChiToccaChiamare()
	inChiamata := partita.PassaChiamata(giocatore)
	if len(inChiamata) == 1 {
		partita.SetChiamante(inChiamata[0])
	}
	setResponse(w, "OK", http.StatusOK)
}

func GetChiamabiliHandler(w http.ResponseWriter, req *http.Request, db *database.SqliteItemsAdapter) {
	chiamabili := partita.GetChiamabili()
	toccaA := partita.GetAChiToccaChiamare()
	chiamante := partita.GetChiamante()
	valChiamato := partita.GetValChiamato()
	setResponse(w, models.GetChiamabiliResp{Chiamabili: chiamabili, ToccaA: toccaA, Chiamante: chiamante, ValChiamato: valChiamato}, http.StatusOK)
}

func GetFineChiamataHandler(w http.ResponseWriter, req *http.Request, db *database.SqliteItemsAdapter) {
	cartaChiamata := partita.GetCartaChiamata()
	setResponse(w, cartaChiamata, http.StatusOK)
}

func GetGiocataHandler(w http.ResponseWriter, req *http.Request, db *database.SqliteItemsAdapter) {
	carteInMano := partita.GetCarteInMano()
	cartePrese := partita.GetCartePrese()
	toccaA := partita.GetAChiTocca()
	carteGiocate := partita.CarteGiocate()
	mano := partita.ManiGiocate()

	setResponse(w, models.GetGiocataResp{CarteInMano: carteInMano, CartePrese: cartePrese,
		ToccaA: toccaA, CarteGiocate: carteGiocate, Mano: mano}, http.StatusOK)
}

func MostraVittoriaHandler(w http.ResponseWriter, req *http.Request, db *database.SqliteItemsAdapter) {
	vittoria := partita.GetVittoria()
	setResponse(w, vittoria, http.StatusOK)
}

func GetFineRoundHandler(w http.ResponseWriter, req *http.Request, db *database.SqliteItemsAdapter) {
	idStr := mux.Vars(req)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		setResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	isLaFine := partita.TryFineRound(id)
	setResponse(w, isLaFine, http.StatusOK)
}
