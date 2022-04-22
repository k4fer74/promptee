package osis

type Bible struct {
	Books []Book `xml:"div"`
}

type Book struct {
	Name     string    `xml:"osisID,attr"`
	Chapters []Chapter `xml:"chapter"`
}

type Chapter struct {
	Verses []Verse `xml:"verse"`
}

type Verse struct {
	Text string `xml:",chardata"`
}
