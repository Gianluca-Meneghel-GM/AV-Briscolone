package api

import (
	"briscolone/database"
	"briscolone/partita"
	"briscolone/store"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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
	partita.ResetContoRound()
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
	puntiChiamatiStr := mex.Params[1]
	chiamanteProvvisorio := partita.GetChiamanteProvvisorio()
	partita.AumentaAChiToccaChiamare()
	haPassato := -1
	if valoreStr == "passo" {
		haPassato = giocatore
		inChiamata := partita.PassaChiamata(giocatore)
		if len(inChiamata) == 1 {
			partita.SetChiamante(inChiamata[0])
		}
	} else if valoreStr == "a monte" {
		broadcast(aMonteResp{Azione: "aMonte", Giocatore: giocatore})
		return
	} else {
		chiamanteProvvisorio = giocatore
		partita.SetChiamanteProvvisorio(giocatore)
		valore, err := strconv.Atoi(valoreStr)
		if err != nil {
			log.Println("errore leggendo val chiamato: " + err.Error())
			return
		}
		puntiChiamati, err := strconv.Atoi(puntiChiamatiStr)
		if err != nil {
			log.Println("errore leggendo punti chiamati: " + err.Error())
			return
		}
		partita.ChiamaValore(valore, puntiChiamati)
	}
	chiamabili := partita.GetChiamabili()
	toccaA := partita.GetAChiToccaChiamare()
	chiamante := partita.GetChiamante()
	valChiamato := partita.GetValChiamato()
	puntiPerVittoria := partita.GetPuntiVittoria()
	broadcast(setChiamabiliResp{Azione: "setChiamabili", Chiamabili: chiamabili, ToccaA: toccaA, Chiamante: chiamante,
		ValChiamato: valChiamato, PuntiVittoria: puntiPerVittoria, ChiamanteProvvisorio: chiamanteProvvisorio, HaPassato: haPassato})
}

func messaggioChatHandler(mex messaggioDaClient) {
	giocatore := mex.Mittente
	messaggioChat := mex.Params[0]
	broadcast(messaggioChatResp{Azione: "showMessaggioChat", Giocatore: giocatore, Messaggio: messaggioChat})
}

func chiamaBotHandler() {
	toccaA := partita.GetAChiToccaChiamare()
	chiamaValoreHandler(messaggioDaClient{Mittente: toccaA, Params: []string{"passo", "61"}})
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

func giocaCartaBotHandler(mex messaggioDaClient) {
	toccaA := partita.GetAChiTocca()
	if strconv.Itoa(toccaA) != mex.Params[0] {
		return
	}
	time.Sleep(1000)
	partita.GiocaCartaBot(toccaA)
	broadcast(getGiocata())
}

func iniziaManoHandler() {
	toccaA := partita.GetAChiTocca()
	carteInMano := partita.GetCarteInMano()
	cartePrese := partita.GetCartePrese()
	mano := partita.ManiGiocate()
	broadcast(setGiocataResp{Azione: "setGiocata", ToccaA: toccaA, CarteInMano: carteInMano, CartePrese: cartePrese, Mano: mano})
}

func altroRoundHandler(mex messaggioDaClient) {
	aMonte := mex.Params[0] == "true"
	if aMonte || mex.Mittente == 0 {
		partita.ResetRound()
	}
	for i := 0; i < 5; i++ {
		partita.SetGiocatorePronto(i)
	}
	pronti := partita.GetGiocatoriPronti()
	for _, g := range pronti {
		if !g {
			return
		}
	}
	iniziaPartita()
}

func bastaCosiHandler() {
	partita.SalvaRecordPartita()
	giocatori := partita.GetGiocatori()
	broadcast(setFinePartitaResp{Azione: "finePartita", Giocatori: *giocatori})
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
