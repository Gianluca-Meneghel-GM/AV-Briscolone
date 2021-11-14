package api

import (
	"briscolone/models"
	"encoding/json"
	"net/http"
	"time"
)

type Response struct {
	Message   string    `json:"message"`
	TimeStamp time.Time `json:"time_stamp"`
}

type setRegistratiResp struct {
	Azione   string
	Iscritti int
}

type IniziaPartitaResp struct {
	Azione           string
	Carte            []models.Carta
	CurrentGiocatori []models.CurrentGiocatore
	ToccaA           int
}

type setChiamabiliResp struct {
	Azione               string
	Chiamabili           []int
	ToccaA               int
	Chiamante            int
	ChiamanteProvvisorio int
	ValChiamato          int
	PuntiVittoria        int
	HaPassato            int
}

type setGiocataResp struct {
	Azione       string
	CarteInMano  []int
	CartePrese   []int
	Mano         int
	CarteGiocate map[int]models.Carta
	ToccaA       int
}

type setFinePartitaResp struct {
	Azione    string
	Giocatori []models.Giocatore
}

type iniziaRoundResp struct {
	Azione   string
	ToccaA   int
	Briscola string
}

type messaggioChatResp struct {
	Azione    string
	Giocatore int
	Messaggio string
}

func setResponse(w http.ResponseWriter, body interface{}, statusCode int) {
	if statusCode != http.StatusOK {
		w.WriteHeader(statusCode)
	}

	data, err := json.Marshal(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}
