package prompt

import "encoding/json"

type BroadcastKind string

var (
	BibleText  = BroadcastKind("bible_text")
	SongLyrics = BroadcastKind("song_lyrics")
)

type Broadcast struct {
	Kind    BroadcastKind `json:"kind"`
	Message string        `json:"message"`
}

type BibleMessage struct {
	BookName      string `json:"book_name"`
	ChapterNumber int    `json:"chapter_number"`
	VerseNumber   int    `json:"verse_number"`
	VerseText     string `json:"verse_text"`
}

type LyricsMessage struct {
	Author string
	Song   string
	Lyrics []string
}

func ToBroadcast(kind BroadcastKind, content interface{}) (*Broadcast, error) {
	b, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}
	bc := &Broadcast{
		Kind:    kind,
		Message: string(b),
	}
	return bc, nil
}
