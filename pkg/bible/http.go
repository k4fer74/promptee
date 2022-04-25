package bible

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type RestAPI struct {
	Store  Store
	Router *mux.Router
}

type RestConf struct {
	Store  Store
	Router *mux.Router
}

func NewRestAPI(conf RestConf) *RestAPI {
	r := &RestAPI{
		Store:  conf.Store,
		Router: conf.Router,
	}
	return r
}

func (h *RestAPI) RegisterHandlers() {
	h.Router.HandleFunc("/books", h.GetBooksHandler).Methods("GET")
	h.Router.HandleFunc("/books/{bookNumber:[0-9]+}/chapters", h.GetChaptersHandler).Methods("GET")
	h.Router.HandleFunc("/books/{bookNumber:[0-9]+}/chapters/{chapterNumber:[0-9]+}/verses", h.GetVersesHandler).Methods("GET")
	h.Router.HandleFunc("/books/{bookNumber:[0-9]+}/chapters/{chapterNumber:[0-9]+}/verses/{verseNumber:[0-9]+}", h.GetVerseHandler).Methods("GET")
}

func (h *RestAPI) GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, _ := h.Store.Books()
	b, _ := json.Marshal(books)
	fmt.Fprintf(w, "%s", b)
}

func (h *RestAPI) GetChaptersHandler(w http.ResponseWriter, r *http.Request) {
	bookNumber := mux.Vars(r)["bookNumber"]
	number, _ := strconv.Atoi(bookNumber)
	chapters, _ := h.Store.Chapters(number)
	c, _ := json.Marshal(chapters)
	fmt.Fprintf(w, "%s", c)
}

func (h *RestAPI) GetVersesHandler(w http.ResponseWriter, r *http.Request) {
	bNumber := mux.Vars(r)["bookNumber"]
	cNumber := mux.Vars(r)["chapterNumber"]
	bookNumber, _ := strconv.Atoi(bNumber)
	chapterNumber, _ := strconv.Atoi(cNumber)
	verses, _ := h.Store.Verses(bookNumber, chapterNumber)
	v, _ := json.Marshal(verses)
	fmt.Fprintf(w, "%s", v)
}

func (h *RestAPI) GetVerseHandler(w http.ResponseWriter, r *http.Request) {
	bNumber := mux.Vars(r)["bookNumber"]
	cNumber := mux.Vars(r)["chapterNumber"]
	vNumber := mux.Vars(r)["verseNumber"]
	bookNumber, _ := strconv.Atoi(bNumber)
	chapterNumber, _ := strconv.Atoi(cNumber)
	verseNumber, _ := strconv.Atoi(vNumber)
	verse, _ := h.Store.SingleVerse(bookNumber, chapterNumber, verseNumber)
	v, _ := json.Marshal(verse)
	fmt.Fprintf(w, "%s", v)
}
