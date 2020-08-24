package database

import (
	"../models"
	"github.com/jmoiron/sqlx"
)

type SqliteItemsAdapter struct {
	db *sqlx.DB
}

func NewSqliteItemsAdapter(db *sqlx.DB) *SqliteItemsAdapter {
	return &SqliteItemsAdapter{
		db: db,
	}
}

func (s SqliteItemsAdapter) GetGiocatori() (*[]models.Giocatore, error) {
	return models.GetGiocatori(s.db)
}

func (s SqliteItemsAdapter) GetNumGiocatori() (int, error) {
	return models.GetNumGiocatori(s.db)
}

func (s SqliteItemsAdapter) GetCurrentGiocatori() (*[]models.CurrentGiocatore, error) {
	return models.GetCurrentGiocatori(s.db)
}

func (s SqliteItemsAdapter) SvuotaCurrentGiocatori() error {
	return models.SvuotaCurrentGiocatori(s.db)
}

func (s SqliteItemsAdapter) AddGiocatore(nome string) error {
	return models.AddGiocatore(s.db, nome)
}

func (s SqliteItemsAdapter) SalvaRecordPartita() {
	models.SalvaRecordPartita(s.db)
}

func (s SqliteItemsAdapter) SelectGiocatore(nome string, isBot int) (int, error) {
	return models.SelectGiocatore(s.db, nome, isBot)
}

func (s SqliteItemsAdapter) AssegnaPuntiVittoria(puntivittoria []int) error {
	return models.AssegnaPuntiVittoria(s.db, puntivittoria)
}

func (s SqliteItemsAdapter) GetPuntiPartita(puntivittoria []int) error {
	return models.AssegnaPuntiVittoria(s.db, puntivittoria)
}

