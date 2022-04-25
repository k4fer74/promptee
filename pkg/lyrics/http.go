package lyrics

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type RestAPI struct {
	Registry Registry
	Router   *mux.Router
}

type RestConf struct {
	Registry Registry
	Router   *mux.Router
}

func NewLyricsAPI(conf RestConf) *RestAPI {
	a := &RestAPI{
		Registry: conf.Registry,
		Router:   conf.Router,
	}
	return a
}

func (a *RestAPI) RegisterHandlers() {
	a.Router.HandleFunc("/songs", a.GetSongsHandler).Methods("GET")
	a.Router.HandleFunc("/songs/{id:[0-9]+}", a.GetSongHandler).Methods("GET")
}

func (a *RestAPI) GetSongsHandler(w http.ResponseWriter, r *http.Request) {
	books, _ := a.Registry.GetSongs()
	b, _ := json.Marshal(books)
	fmt.Fprintf(w, "%s", b)
}

func (a *RestAPI) GetSongHandler(w http.ResponseWriter, r *http.Request) {
	songID := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(songID)
	song, _ := a.Registry.GetSong(id)
	s, _ := json.Marshal(song)
	fmt.Fprintf(w, "%s", s)
}
