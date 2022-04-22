package osis

import (
	"encoding/xml"
	"github.com/k4fer74/promptee/bible"
	"io/ioutil"
	"os"
)

const (
	JFAA = bible.Version("João Ferreira de Almeida Atualizada")
)

// OSIS implements bible.Store
type OSIS struct {
	Bible
}

type XMLFile struct {
	XMLName xml.Name `xml:"osis"`
	Bible   Bible    `xml:"osisText"`
}

func New(fileName string) (bible.Store, error) {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var data XMLFile
	if err := xml.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	osis := &OSIS{
		Bible: data.Bible,
	}
	return osis, nil
}

func (o *OSIS) Version() bible.Version {
	return JFAA
}

func (o *OSIS) Books() ([]bible.Book, error) {
	var books []bible.Book
	for i, _ := range o.Bible.Books {
		name, err := bible.HumanReadableBookName(i + 1)
		if err != nil {
			return nil, err
		}
		books = append(books, bible.Book{
			Number: i + 1,
			Name:   name,
		})
	}
	return books, nil
}

func (o *OSIS) BookByOrder(order int) (*Book, error) {
	for i, b := range o.Bible.Books {
		if (i + 1) == order {
			return &b, nil
		}
	}
	return nil, ErrBookNotFound
}

func (o *OSIS) Chapters(bookNumber int) (c []bible.Chapter, err error) {
	b, err := o.BookByOrder(bookNumber)
	if err != nil {
		return nil, err
	}
	for i, _ := range b.Chapters {
		c = append(c, bible.Chapter{
			Number: i + 1,
		})
	}
	return
}

func (o *OSIS) ChapterByNumber(bookOrder, chapterNumber int) (*Chapter, error) {
	b, err := o.BookByOrder(bookOrder)
	if err != nil {
		return nil, err
	}
	for i, c := range b.Chapters {
		if (i + 1) == chapterNumber {
			return &c, nil
		}
	}
	return nil, ErrChapterNotFound
}

func (o *OSIS) Verses(bookNumber, chapterNumber int) (v []bible.Verse, err error) {
	c, err := o.ChapterByNumber(bookNumber, chapterNumber)
	if err != nil {
		return nil, err
	}
	for i, verse := range c.Verses {
		v = append(v, bible.Verse{
			Number: i + 1,
			Text:   verse.Text,
		})
	}
	return
}

func (o *OSIS) SingleVerse(bookNumber, chapterNumber, verseNumber int) (*bible.Verse, error) {
	c, err := o.ChapterByNumber(bookNumber, chapterNumber)
	if err != nil {
		return nil, err
	}
	for i, verse := range c.Verses {
		if (i + 1) == verseNumber {
			return &bible.Verse{
				Number: verseNumber,
				Text:   verse.Text,
			}, nil
		}
	}
	return nil, ErrVerseNotFound
}
