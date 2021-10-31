package partita

import (
	"briscolone/models"
	"math/rand"
	"strconv"
	"time"
)

var cartePrese = []int{0, 0, 0, 0, 0}
var lastCartaGiocata models.Carta
var carteGiocate = make(map[int]models.Carta)
var briscola string
var briscolaTemp string
var puntiChiamante int
var puntiAltri int
var maniGiocate = 0
var puntiPerVittoria = 61
var haFinito = []bool{false, false, false, false, false}

func GetCartePrese() []int {
	return cartePrese
}

func GetCarteInMano() []int {
	inMano := []int{0, 0, 0, 0, 0}
	for i, g := range mano {
		inMano[i] = len(g)
	}
	return inMano
}

func SetLastCartaGiocata(semeStr string, val int) {
	seme := []rune(semeStr)
	lastCartaGiocata = models.Carta{Valore: val, Seme: seme[0], SemeStr: semeStr}
}

func CarteGiocate() map[int]models.Carta {
	return carteGiocate
}

func ManiGiocate() int {
	return maniGiocate
}

func GiocaCarta(seme string, valore, giocatore int) int {
	SetLastCartaGiocata(seme, valore)
	if len(carteGiocate) == 0 {
		briscolaTemp = seme
	}
	rimuoviCartaDaMano(giocatore, valore, seme)
	carteGiocate[giocatore] = models.CarteMap[strconv.Itoa(valore)+seme]
	if len(carteGiocate) == 5 {
		go finisciMano()
		return aChiTocca
	} else {
		AumentaAChiTocca()
		return aChiTocca
	}
}

func finisciMano() {
	time.Sleep(3000)
	prende := chiPrende()
	cartePrese[prende] += 5
	assegnaPunti(prende)
	ResetMano()
	maniGiocate++
	if maniGiocate == 8 {
		finisciRound()
	}
	aChiTocca = prende
}

func GiocaCartaBot(giocatore int) int {
	if len(mano[giocatore]) > 0 {
		idx := rand.Intn(len(mano[giocatore]))
		carta := mano[giocatore][idx]
		return GiocaCarta(carta.SemeStr, carta.Valore, giocatore)
	}
	return -1
}

func rimuoviCartaDaMano(giocatore, val int, semeStr string) {
	seme := []rune(semeStr)
	var nuovaMano []models.Carta
	for _, c := range mano[giocatore] {
		if c.Seme != seme[0] || c.Valore != val {
			nuovaMano = append(nuovaMano, c)
		}
	}
	mano[giocatore] = nuovaMano
}

func chiPrende() int {
	for _, c := range carteGiocate {
		if c.SemeStr == briscola {
			briscolaTemp = briscola
			break
		}
	}
	prende := 0
	comanda := models.Carta{Valore: 0}
	for g, c := range carteGiocate {
		if c.SemeStr == briscolaTemp {
			if c.Valore > comanda.Valore {
				comanda = c
				prende = g
			}
		}
	}
	return prende
}

func assegnaPunti(giocatore int) {
	punti := 0
	for _, c := range carteGiocate {
		punti += c.Punti
	}
	if giocatore == chiamante || giocatore == socio {
		puntiChiamante += punti
	} else {
		puntiAltri += punti
	}
}

func finisciRound() error {
	puntiGiocatore := getPuntiVittoria()
	err := db.AssegnaPuntiVittoria(puntiGiocatore)
	return err
}

func getPuntiVittoria() []int {
	puntigiocatore := []int{0, 0, 0, 0, 0}
	moltiplicatore := 1
	if puntiPerVittoria > 70 {
		moltiplicatore = 2
	} else if puntiPerVittoria > 80 {
		moltiplicatore = 3
	}
	if puntiChiamante == 120 {
		moltiplicatore *= 2
	}
	for g := range puntigiocatore {
		var pti int
		if g == chiamante {
			if puntiChiamante >= puntiPerVittoria {
				pti = 2
			} else {
				pti = -2
			}
		} else if g == socio {
			if puntiChiamante >= puntiPerVittoria {
				pti = 1
			} else {
				pti = -1
			}
		} else {
			if puntiChiamante < puntiPerVittoria {
				pti = 1
			} else {
				pti = -1
			}
		}
		puntigiocatore[g] += pti * moltiplicatore
	}
	return puntigiocatore
}

func GetVittoria() models.Vittoria {
	vincitori := []int{}
	puntiPartita := []int{0, 0, 0, 0, 0}
	puntiVincitori := 0
	if puntiChiamante >= puntiPerVittoria {
		vincitori = []int{chiamante, socio}
		puntiVincitori = puntiChiamante
	} else {
		for g := 0; g < 5; g++ {
			if g != chiamante && g != socio {
				vincitori = append(vincitori, g)
				puntiVincitori = puntiAltri
			}
		}
	}
	giocatori, _ := db.GetCurrentGiocatori()
	for _, g := range *giocatori {
		puntiPartita[g.Id] = g.Punti
	}
	return models.Vittoria{Vincitori: vincitori, PuntiVincitori: puntiVincitori, PuntiPartita: puntiPartita}
}

func TryFineRound(giocatore int) bool {
	haFinito[giocatore] = true
	isLaFine := true
	for _, g := range haFinito {
		if !g {
			isLaFine = false
			break
		}
	}
	return isLaFine
}

func ResetMano() {
	briscolaTemp = ""
	carteGiocate = map[int]models.Carta{}
	lastCartaGiocata = models.Carta{}
}
