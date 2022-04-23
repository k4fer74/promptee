package rest

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/k4fer74/promptee/pkg/bible"
	"github.com/k4fer74/promptee/pkg/prompt"
	"net/http"
	"strings"
	"time"
)

type Server struct {
	Router     *mux.Router
	BibleStore bible.Store
	HTTPServer *http.Server
}

func NewServer(bibleStore bible.Store) *Server {
	return &Server{
		Router:     mux.NewRouter(),
		BibleStore: bibleStore,
	}
}

func (s *Server) RegisterStaticFileServer() {
	fs := http.FileServer(http.Dir("webui/static"))
	s.Router.PathPrefix("/static/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, ".css") {
			w.Header().Set("Content-Type", "text/css")
		}

		if strings.Contains(r.URL.Path, ".png") || strings.Contains(r.URL.Path, ".svg") {
			w.Header().Set("Content-Type", "text/image")
		}

		if strings.Contains(r.URL.Path, ".js") {
			w.Header().Set("Content-Type", "text/javascript")
		}
	}).Handler(http.StripPrefix("/static/", fs))
}

func (s *Server) RegisterHandlers() {
	s.Router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	s.RegisterStaticFileServer()

	apiRouter := s.Router.PathPrefix("/api").Subrouter()

	bibleAPI := bible.NewRestAPI(bible.RestConf{
		Store:  s.BibleStore,
		Router: apiRouter,
	})

	promptAPI := prompt.NewRestAPI(prompt.RestConf{
		Router: apiRouter,
	})

	bibleAPI.RegisterHandlers()
	promptAPI.RegisterHandlers()

}

func (s *Server) StartAPI() error {
	s.RegisterHandlers()

	hs := http.Server{
		Addr:              "127.0.0.1:3160",
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
