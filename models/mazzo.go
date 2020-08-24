package models

type Carta struct {
	Valore     int
	Seme       rune
	SemeStr    string
	Punti      int
	Soprannome string
	isBriscola bool
}

var Carte = []Carta{{Valore: 2, SemeStr: "B", Seme: 'B'}, {Valore: 4, SemeStr: "B", Seme: 'B'}, {Valore: 5, SemeStr: "B", Seme: 'B'}, {Valore: 6, SemeStr: "B", Seme: 'B'}, {Valore: 7, SemeStr: "B", Seme: 'B'},
	{Valore: 8, SemeStr: "B", Seme: 'B', Punti: 2}, {Valore: 9, SemeStr: "B", Seme: 'B', Punti: 3}, {Valore: 10, SemeStr: "B", Seme: 'B', Punti: 4}, {Valore: 13, SemeStr: "B", Seme: 'B', Punti: 10}, {Valore: 21, SemeStr: "B", Seme: 'B', Punti: 11},
	{Valore: 2, SemeStr: "C", Seme: 'C'}, {Valore: 4, SemeStr: "C", Seme: 'C'}, {Valore: 5, SemeStr: "C", Seme: 'C'}, {Valore: 6, SemeStr: "C", Seme: 'C'}, {Valore: 7, SemeStr: "C", Seme: 'C'},
	{Valore: 8, SemeStr: "C", Seme: 'C', Punti: 2}, {Valore: 9, SemeStr: "C", Seme: 'C', Punti: 3}, {Valore: 10, SemeStr: "C", Seme: 'C', Punti: 4}, {Valore: 13, SemeStr: "C", Seme: 'C', Punti: 10}, {Valore: 21, SemeStr: "C", Seme: 'C', Punti: 11},
	{Valore: 2, SemeStr: "D", Seme: 'D'}, {Valore: 4, SemeStr: "D", Seme: 'D'}, {Valore: 5, SemeStr: "D", Seme: 'D'}, {Valore: 6, SemeStr: "D", Seme: 'D'}, {Valore: 7, SemeStr: "D", Seme: 'D'},
	{Valore: 8, SemeStr: "D", Seme: 'D', Punti: 2}, {Valore: 9, SemeStr: "D", Seme: 'D', Punti: 3}, {Valore: 10, SemeStr: "D", Seme: 'D', Punti: 4}, {Valore: 13, SemeStr: "D", Seme: 'D', Punti: 10}, {Valore: 21, SemeStr: "D", Seme: 'D', Punti: 11},
	{Valore: 2, SemeStr: "S", Seme: 'S', Soprannome: "La bisognosa"}, {Valore: 4, SemeStr: "S", Seme: 'S'}, {Valore: 5, SemeStr: "S", Seme: 'S'}, {Valore: 6, SemeStr: "S", Seme: 'S'}, {Valore: 7, SemeStr: "S", Seme: 'S'},
	{Valore: 8, SemeStr: "S", Seme: 'S', Punti: 2}, {Valore: 9, SemeStr: "S", Seme: 'S', Punti: 3}, {Valore: 10, SemeStr: "S", Seme: 'S', Punti: 4}, {Valore: 13, SemeStr: "S", Seme: 'S', Punti: 10}, {Valore: 21, SemeStr: "S", Seme: 'S', Punti: 11}}

var CarteMap = map[string]Carta{"2B": {Valore: 2, SemeStr: "B"}, "4B": {Valore: 4, SemeStr: "B"}, "5B": {Valore: 5, SemeStr: "B"}, "6B": {Valore: 6, SemeStr: "B"}, "7B": {Valore: 7, SemeStr: "B"},
	"8B": {Valore: 8, SemeStr: "B", Punti: 2}, "9B": {Valore: 9, SemeStr: "B", Punti: 3}, "10B": {Valore: 10, SemeStr: "B", Punti: 4}, "13B": {Valore: 13, SemeStr: "B", Punti: 10}, "21B": {Valore: 21, SemeStr: "B", Punti: 11},
	"2C": {Valore: 2, SemeStr: "C"}, "4C": {Valore: 4, SemeStr: "C"}, "5C": {Valore: 5, SemeStr: "C"}, "6C": {Valore: 6, SemeStr: "C"}, "7C": {Valore: 7, SemeStr: "C"},
	"8C": {Valore: 8, SemeStr: "C", Punti: 2}, "9C": {Valore: 9, SemeStr: "C", Punti: 3}, "10C": {Valore: 10, SemeStr: "C", Punti: 4}, "13C": {Valore: 13, SemeStr: "C", Punti: 10}, "21C": {Valore: 21, SemeStr: "C", Punti: 11},
	"2D": {Valore: 2, SemeStr: "D"}, "4D": {Valore: 4, SemeStr: "D"}, "5D": {Valore: 5, SemeStr: "D"}, "6D": {Valore: 6, SemeStr: "D"}, "7D": {Valore: 7, SemeStr: "D"},
	"8D": {Valore: 8, SemeStr: "D", Punti: 2}, "9D": {Valore: 9, SemeStr: "D", Punti: 3}, "10D": {Valore: 10, SemeStr: "D", Punti: 4}, "13D": {Valore: 13, SemeStr: "D", Punti: 10}, "21D": {Valore: 21, SemeStr: "D", Punti: 11},
	"2S": {Valore: 2, SemeStr: "S", Soprannome: "La bisognosa"}, "4S": {Valore: 4, SemeStr: "S"}, "5S": {Valore: 5, SemeStr: "S"}, "6S": {Valore: 6, SemeStr: "S"}, "7S": {Valore: 7, SemeStr: "S"},
	"8S": {Valore: 8, SemeStr: "S", Punti: 2}, "9S": {Valore: 9, SemeStr: "S", Punti: 3}, "10S": {Valore: 10, SemeStr: "S", Punti: 4}, "13S": {Valore: 13, SemeStr: "S", Punti: 10}, "21S": {Valore: 21, SemeStr: "S", Punti: 11}}
