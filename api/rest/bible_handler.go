package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (s *Server) GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, _ := s.BibleStore.Books()
	b, _ := json.Marshal(books)
	fmt.Fprintf(w, "%s", b)
}

func (s *Server) GetChaptersHandler(w http.ResponseWriter, r *http.Request) {
	bookNumber := mux.Vars(r)["bookNumber"]
	number, _ := strconv.Atoi(bookNumber)
	chapters, _ := s.BibleStore.Chapters(number)
	c, _ := json.Marshal(chapters)
	fmt.Fprintf(w, "%s", c)
}

func (s *Server) GetVersesHandler(w http.ResponseWriter, r *http.Request) {
	bNumber := mux.Vars(r)["bookNumber"]
	cNumber := mux.Vars(r)["chapterNumber"]
	bookNumber, _ := strconv.Atoi(bNumber)
	chapterNumber, _ := strconv.Atoi(cNumber)
	verses, _ := s.BibleStore.Verses(bookNumber, chapterNumber)
	v, _ := json.Marshal(verses)
	fmt.Fprintf(w, "%s", v)
}

func (s *Server) GetVerseHandler(w http.ResponseWriter, r *http.Request) {
	bNumber := mux.Vars(r)["bookNumber"]
	cNumber := mux.Vars(r)["chapterNumber"]
	vNumber := mux.Vars(r)["verseNumber"]
	bookNumber, _ := strconv.Atoi(bNumber)
	chapterNumber, _ := strconv.Atoi(cNumber)
	verseNumber, _ := strconv.Atoi(vNumber)
	verse, _ := s.BibleStore.SingleVerse(bookNumber, chapterNumber, verseNumber)
	v, _ := json.Marshal(verse)
	fmt.Fprintf(w, "%s", v)
}
