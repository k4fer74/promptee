package bible

type Store interface {
	Books() ([]Book, error)
	Chapters(bookNumber int) ([]Chapter, error)
	Verses(bookNumber int, chapterNumber int) ([]Verse, error)
	SingleVerse(bookNumber int, chapterNumber int, verseNumber int) (*Verse, error)
}
