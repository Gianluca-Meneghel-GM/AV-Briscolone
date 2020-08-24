package partita

import "strconv"

var currentValChiamato = 30
var chiamabili = []int{21, 13, 10, 9, 8, 7, 6, 5, 4, 2}
var inChiamata = []int{0, 1, 2, 3, 4}
var aChiToccaChiamare int
var chiamante = -1
var socio = -1
var cartaChiamata string
var haChiamato = false

func ChiamaValore(val int) {
	currentValChiamato = val
}

func GetChiamabili() []int {
	var currChiamabili []int
	for _, c := range chiamabili {
		if c < currentValChiamato {
			currChiamabili = append(currChiamabili, c)
		}
	}
	return currChiamabili
}

func GetHaChiamato() bool {
	return haChiamato
}

func GetCartaChiamata() string {
	return cartaChiamata
}

func GetAChiToccaChiamare() int {
	return aChiToccaChiamare
}

func AumentaAChiToccaChiamare() {
	if aChiToccaChiamare == inChiamata[len(inChiamata)-1] {
		aChiToccaChiamare = inChiamata[0]
	} else {
		for i, g := range inChiamata {
			if aChiToccaChiamare == g {
				aChiToccaChiamare = inChiamata[i+1]
				return
			}
		}
	}
}

func PassaChiamata(giocatore int) []int {
	var newInChiam []int
	for _, g := range inChiamata {
		if g != giocatore {
			newInChiam = append(newInChiam, g)
		}
	}
	inChiamata = newInChiam
	return inChiamata
}

func SetChiamante(giocatore int) {
	chiamante = giocatore
}

func GetChiamante() int {
	return chiamante
}

func GetValChiamato() int {
	return currentValChiamato
}

func SetChiamata(seme string) {
	if currentValChiamato == 30 {
		currentValChiamato = 21
	}
	valStr := strconv.Itoa(currentValChiamato)
	cartaChiamata = valStr + seme
	socio = findSocio(currentValChiamato, seme)
	briscola = seme
	haChiamato = true
}

func findSocio(val int, semeStr string) int {
	seme := []rune(semeStr)
	for i, m := range mano {
		for _, c := range m {
			if c.Valore == val && c.Seme == seme[0] {
				return i
			}
		}
	}
	return 0
}
