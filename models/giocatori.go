package models

import (
	"briscolone/store"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Giocatore struct {
	Nome     string `db:"nome"`
	Vittorie int    `db:"vittorie"`
	Round    int    `db:"round"`
	Partite  int    `db:"partite"`
	PuntiTot int    `db:"puntitotali"`
}

type CurrentGiocatore struct {
	Id    int    `db:"id"`
	Nome  string `db:"nome"`
	Punti int    `db:"puntipartita"`
	IsBot int    `db:"isbot"`
}

func GetGiocatori(db *sqlx.DB) (*[]Giocatore, error) {

	giocatori := []Giocatore{}
	err := db.Select(&giocatori, `SELECT * FROM giocatori WHERE nome NOT IN (SELECT nome FROM currentGiocatori)`, nil)
	if err != nil {
		return nil, err
	}

	return &giocatori, nil
}

func AddGiocatore(db *sqlx.DB, nome string) error {

	_, err := db.Exec(`INSERT INTO giocatori VALUES ($1,0,0,0,0) `, nome)
	if err != nil {
		return err
	}

	return nil
}

func SalvaRecordPartita(db *sqlx.DB, contoRound int) {
	finePartita, err := GetCurrentGiocatori(db)
	if err != nil {
		log.Println("Errore prendendo currentGiocatori: " + err.Error())
		return
	}
	giocatori, err := GetGiocatori(db)
	if err != nil {
		log.Println("Errore prendendo giocatori: " + err.Error())
		return
	}
	puntiVincitore := getPuntiVincitore(finePartita)
	for _, g := range *giocatori {
		if id := getIdGiocatore(finePartita, g.Nome); id != -1 {
			g.Partite++
			punti := (*finePartita)[id].Punti
			g.PuntiTot += punti
			if punti == puntiVincitore {
				g.Vittorie++
			}
			g.Round += contoRound
			_, err := db.Exec(`UPDATE giocatori SET partite = $1, vittorie = $2, puntitotali = $3, round = $4 WHERE nome = $5`,
				g.Partite, g.Vittorie, g.PuntiTot, g.Round, g.Nome)
			if err != nil {
				log.Println("Errore salvando record: " + err.Error())
			}
		}
	}
}

func getPuntiVincitore(giocatori *[]CurrentGiocatore) int {
	max := 0
	for _, g := range *giocatori {
		if g.Punti > max {
			max = g.Punti
		}
	}
	return max
}

func getIdGiocatore(giocatori *[]CurrentGiocatore, nome string) int {
	for _, g := range *giocatori {
		if g.Nome == nome {
			return g.Id
		}
	}
	return -1
}

func SelectGiocatore(db *sqlx.DB, nome string, isBot int) (int, error) {
	id := store.AumentaNumeroGiocatori()
	_, err := db.Exec(`INSERT INTO currentgiocatori VALUES ($1,$2,0,$3) `, id, nome, isBot)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetNumGiocatori(db *sqlx.DB) (int, error) {

	var numGiocatori []int64
	err := db.Select(&numGiocatori, `SELECT COUNT(*) FROM currentgiocatori`)
	if err != nil {
		return 0, err
	}

	return int(numGiocatori[0]), nil
}

func GetCurrentGiocatori(db *sqlx.DB) (*[]CurrentGiocatore, error) {

	var giocatori []CurrentGiocatore
	err := db.Select(&giocatori, `SELECT * FROM currentgiocatori`)
	if err != nil {
		return nil, err
	}

	return &giocatori, nil
}

func SvuotaCurrentGiocatori(db *sqlx.DB) error {

	_, err := db.Exec(`DELETE FROM currentgiocatori`)
	if err != nil {
		return err
	}

	return nil
}

func AssegnaPuntiVittoria(db *sqlx.DB, puntiVittoria []int) error {
	for g := range puntiVittoria {
		var currentPunti []int
		err := db.Select(&currentPunti, `SELECT puntipartita FROM currentgiocatori WHERE id = $1`, g)
		if err != nil {
			return err
		}

		currentPunti[0] += puntiVittoria[g]
		_, err = db.Exec(`UPDATE currentgiocatori SET puntipartita = $1 WHERE id = $2`, currentPunti[0], g)
		if err != nil {
			return err
		}
	}
	return nil
}
