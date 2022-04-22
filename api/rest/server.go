package rest

import (
	"github.com/gorilla/mux"
	"github.com/k4fer74/promptee/bible"
	"net/http"
	"time"
)

type Server struct {
	Router     *mux.Router
	BibleStore bible.Store
}

func NewServer(bibleStore bible.Store) *Server {
	return &Server{
		Router:     mux.NewRouter(),
		BibleStore: bibleStore,
	}
}

func (s *Server) RegisterHandlers() {
	s.Router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	s.Router.HandleFunc("/books", s.GetBooksHandler).Methods("GET")
	s.Router.HandleFunc("/books/{bookNumber}/chapters", s.GetChaptersHandler).Methods("GET")
	s.Router.HandleFunc("/books/{bookNumber}/chapters/{chapterNumber}/verses", s.GetVersesHandler).Methods("GET")
	s.Router.HandleFunc("/books/{bookNumber}/chapters/{chapterNumber}/verses/{verseNumber}", s.GetVerseHandler).Methods("GET")
}

func (s *Server) ListenAndServe() error {
	s.RegisterHandlers()

	hs := http.Server{
		Addr:              "127.0.0.1:3160",
		Handler:           s.Router,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
	}
	return hs.ListenAndServe()
}
