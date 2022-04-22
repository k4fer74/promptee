package bible

import "errors"

var (
	ErrBookTranslationNotFound = errors.New("BookTranslationNotFound")
)

type Version string

type Book struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
}

func (m *Book) String() (name string) {
	name, _ = HumanReadableBookName(m.Number)
	return
}

type Chapter struct {
	Number int `json:"number"`
}

type Verse struct {
	Number int    `json:"number"`
	Text   string `json:"text"`
}

func HumanReadableBookName(bookNumber int) (string, error) {
	translation := map[int]string{
		1:  "Gênesis",
		2:  "Êxodo",
		3:  "Levítico",
		4:  "Números",
		5:  "Deuteronômio",
		6:  "Josue",
		7:  "Juízes",
		8:  "Rute",
		9:  "ISamuel",
		10: "2Samuel",
		11: "1Reis",
		12: "2Reis",
		13: "1Crônicas",
		14: "2Crônicas",
		15: "Esdras",
		16: "Neemias",
		17: "Ester",
		18: "Jó",
		19: "Salmos",
		20: "Provérbios",
		21: "Eclesiastes",
		22: "Cantares",
		23: "Isaías",
		24: "Jeremias",
		25: "Lamentações",
		26: "Ezequiel",
		27: "Daniel",
		28: "Oséias",
		29: "Joel",
		30: "Amós",
		31: "Obdias",
		32: "Jonas",
		33: "Miquéias",
		34: "Naum",
		35: "Habacuc",
		36: "Sofonias",
		37: "Ageu",
		38: "Zacarias",
		39: "Malaquias",
		40: "Mateus",
		41: "Marcos",
		42: "Lucas",
		43: "João",
		44: "Atos dos Apóstolos",
		45: "Romanos",
		46: "1Coríntios",
		47: "2Coríntios",
		48: "Gálatas",
		49: "Efésios",
		50: "Filipenses",
		51: "Colossenses",
		52: "1Tessalonicenses",
		53: "2Tessalonicenses",
		54: "1Timóteo",
		55: "2Timóteo",
		56: "Tito",
		57: "Filémon",
		58: "1Pedro",
		59: "2Pedro",
		60: "1João",
		61: "2João",
		62: "3João",
		63: "Hebreus",
		64: "Tiago",
		65: "Judas",
		66: "Apocalipse",
	}
	if _, ok := translation[bookNumber]; !ok {
		return "", ErrBookTranslationNotFound
	}
	return translation[bookNumber], nil
}
