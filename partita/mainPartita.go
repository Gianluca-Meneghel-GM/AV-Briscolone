package partita

import (
	"briscolone/database"
	"briscolone/models"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var db *database.SqliteItemsAdapter
var mazzo []models.Carta
var mano = [][]models.Carta{{}, {}, {}, {}, {}}
var giocatoriPronti = map[int]bool{0: false, 1: false, 2: false, 3: false, 4: false}
var chiInizia = 0
var aChiTocca int
var contoRound = 0

func Init(database *database.SqliteItemsAdapter) {
	db = database
	Purge()
	FaiIlMazzo()
}

func Purge() {
	err := db.SvuotaCurrentGiocatori()
	if err != nil {
		log.Println("Errore cancellando giocatori correnti: " + err.Error())
	}
}

func IscriviBot(id int) {
	nome := "Bot" + strconv.Itoa(id)
	_, err := db.SelectGiocatore(nome, 1)
	if err != nil {
		log.Println("Errore iscrivendo bot: " + err.Error())
		return
	}
}

func GetGiocatoriPronti() map[int]bool {
	return giocatoriPronti
}

func SetGiocatorePronto(id int) {
	giocatoriPronti[id] = true
	log.Println("è pronto " + strconv.Itoa(id))
}

func FaiIlMazzo() {
	mazzo = models.Carte
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(40, func(i, j int) { mazzo[i], mazzo[j] = mazzo[j], mazzo[i] })
	daiCarte()
}

func daiCarte() {
	for i, c := range mazzo {
		idx := i / 8
		c.SemeStr = string(c.Seme)
		mano[idx] = append(mano[idx], c)
	}
}

func GetCarte(id int) []models.Carta {
	return ordina(mano[id])
}

func ordina(data []models.Carta) []models.Carta {
	if len(data) <= 1 {
		return data
	}
	middle := len(data) / 2
	left := ordina(data[:middle])
	right := ordina(data[middle:])

	return merge(left, right)
}
func merge(left, right []models.Carta) []models.Carta {
	result := make([]models.Carta, len(left)+len(right))
	for i := 0; len(left) > 0 || len(right) > 0; i++ {
		if len(left) > 0 && len(right) > 0 {
			if left[0].Seme < right[0].Seme {
				result[i] = left[0]
				left = left[1:]
			} else if left[0].Seme > right[0].Seme {
				result[i] = right[0]
				right = right[1:]
			} else {
				if left[0].Valore > right[0].Valore {
					result[i] = left[0]
					left = left[1:]
				} else {
					result[i] = right[0]
					right = right[1:]
				}
			}
		} else if len(left) > 0 {
			result[i] = left[0]
			left = left[1:]
		} else if len(right) > 0 {
			result[i] = right[0]
			right = right[1:]
		}
	}
	return result
}

func GetChiInizia() int {
	aChiTocca = chiInizia
	aChiToccaChiamare = chiInizia
	return chiInizia
}

func AumentaChiInizia() {
	if chiInizia == 4 {
		chiInizia = 0
	} else {
		chiInizia++
	}
}

func GetAChiTocca() int {
	return aChiTocca
}

func AumentaAChiTocca() {
	if aChiTocca == 4 {
		aChiTocca = 0
	} else {
		aChiTocca++
	}
}

func GiocatoriIscritti() int {
	numIscritti, err := db.GetNumGiocatori()
	if err != nil {
		return 99
	}
	return numIscritti
}

func GetGiocatori() *[]models.Giocatore {
	gioc, err := db.GetGiocatori()
	if err != nil {
		log.Println("Errore prendendo giocatori: " + err.Error())
		return nil
	}
	return gioc
}

func IniziaPartita() *[]models.CurrentGiocatore {
	ResetMano()
	FaiIlMazzo()
	contoRound++
	giocatori, err := db.GetCurrentGiocatori()
	if err != nil {
		log.Println("errore leggendo giocatori: " + err.Error())
		return nil
	}
	return giocatori
}

func ResetContoRound() {
	contoRound = 0
}

func ResetRound() {
	inChiamata = []int{0, 1, 2, 3, 4}
	chiamante = -1
	aChiToccaChiamare = chiInizia
	currentValChiamato = 30
	chiamabili = []int{21, 13, 10, 9, 8, 7, 6, 5, 4, 2}
	haChiamato = false
	briscola = ""
	puntiChiamante = 0
	puntiAltri = 0
	cartePrese = []int{0, 0, 0, 0, 0}
	maniGiocate = 0
	puntiPerVittoria = 61
	giocatoriPronti = map[int]bool{0: false, 1: false, 2: false, 3: false, 4: false}
	haFinito = []bool{false, false, false, false, false}
	for i := range mano {
		mano[i] = []models.Carta{}
	}
	AumentaChiInizia()
}

func SalvaRecordPartita() {
	db.SalvaRecordPartita(contoRound)
}
