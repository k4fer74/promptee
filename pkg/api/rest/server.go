package rest

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/k4fer74/promptee/pkg/bible"
	"github.com/k4fer74/promptee/pkg/lyrics"
	"github.com/k4fer74/promptee/pkg/prompt"
	"net/http"
	"time"
)

type Server struct {
	Router         *mux.Router
	BibleStore     bible.Store
	LyricsRegistry lyrics.Registry
	HTTPServer     *http.Server
}

func NewServer(bibleStore bible.Store, lyricsRegistry lyrics.Registry) *Server {
	return &Server{
		Router:         mux.NewRouter(),
		BibleStore:     bibleStore,
		LyricsRegistry: lyricsRegistry,
	}
}

func (s *Server) RegisterHandlers() {
	s.Router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			w.Header().Add("Access-Control-Allow-Origin", "*")
			next.ServeHTTP(w, r)
		})
	})

	apiRouter := s.Router.PathPrefix("/api").Subrouter()

	bibleAPI := bible.NewRestAPI(bible.RestConf{
		Store:  s.BibleStore,
		Router: apiRouter,
	})

	lyricsAPI := lyrics.NewLyricsAPI(lyrics.RestConf{
		Registry: s.LyricsRegistry,
		Router:   apiRouter,
	})

	promptAPI := prompt.NewRestAPI(prompt.RestConf{
		Router: apiRouter,
	})

	bibleAPI.RegisterHandlers()
	lyricsAPI.RegisterHandlers()
	promptAPI.RegisterHandlers()
}

func (s *Server) StartAPI() error {
	s.RegisterHandlers()

	hs := http.Server{
		Addr:              ":3160",
		Handler:           s.Router,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
	}

	s.HTTPServer = &hs

	return hs.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.HTTPServer.Shutdown(context.Background())
}
