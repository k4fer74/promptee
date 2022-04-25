package prompt

import "encoding/json"

// TODO: Events

type BroadcastKind string

var (
	Prompt     = BroadcastKind("prompt")
	BibleText  = BroadcastKind("bible_text")
	SongLyrics = BroadcastKind("song_lyrics")
)

type Broadcast struct {
	Kind    BroadcastKind `json:"kind"`
	Message []byte        `json:"message"`
}

type BibleMessage struct {
	Kind          BroadcastKind
	BookName      string `json:"book_name"`
	BookNumber    int    `json:"book_number"`
	ChapterNumber int    `json:"chapter_number"`
	VerseNumber   int    `json:"verse_number"`
	VerseText     string `json:"verse_text"`
}

type LyricsMessage struct {
	Kind   BroadcastKind
	Author string `json:"author"`
	Song   string `json:"song"`
	Text   string `json:"text"`
}

func (k BroadcastKind) ToBroadcast(content interface{}) (*Broadcast, error) {
	b, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}
	bc := &Broadcast{
		Kind:    k,
		Message: b,
	}
	return bc, nil
}

func (k BroadcastKind) UnmarshalBroadcastMessage(payload []byte, broadcastMessage interface{}) error {
	return json.Unmarshal(payload, &broadcastMessage)
}
