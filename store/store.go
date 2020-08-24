package store

var numeroGiocatori = -1

func GetNumeroGiocatori() int {
	return numeroGiocatori
}

func AzzeraNumeroGiocatori() {
	numeroGiocatori = -1
}

func AumentaNumeroGiocatori() int {
	numeroGiocatori++
	return numeroGiocatori
}
