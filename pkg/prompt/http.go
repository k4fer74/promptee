package prompt

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"
	"path"
)

type RestAPI struct {
	Router *mux.Router
}

type RestConf struct {
	Router *mux.Router
}

var Upgrader = websocket.Upgrader{}

func NewRestAPI(conf RestConf) *RestAPI {
	r := &RestAPI{
		Router: conf.Router,
	}
	r.RegisterHandlers()
	return r
}

func (h *RestAPI) RegisterHandlers() {
	h.Router.HandleFunc("/prompter/broadcast", h.BroadcastHandler)
	h.Router.HandleFunc("/prompter", h.PrompterHandler)
}

func (h *RestAPI) BroadcastHandler(w http.ResponseWriter, r *http.Request) {
	c, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func (h *RestAPI) PrompterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	tmpl, err := template.ParseFiles(path.Join("webui/templates", "prompter.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]string{
		"ws_path": "ws://" + r.Host + "/api/prompter/broadcast",
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
