package models

type PrimaMano struct {
	Carte            []Carta
	CurrentGiocatori []CurrentGiocatore
	ToccaA           int
}

type GetChiamabiliResp struct {
	Chiamabili  []int
	ToccaA      int
	Chiamante   int
	ValChiamato int
}

type GetGiocataResp struct {
	CarteInMano  []int
	CartePrese   []int
	CartaGiocata Carta
	CarteGiocate map[int]Carta
	ToccaA       int
	Mano         int
}

type Vittoria struct {
	Vincitori      []int
	PuntiVincitori int
	PuntiPartita   []int
}
