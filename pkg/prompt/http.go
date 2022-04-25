package prompt

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/k4fer74/promptee/pkg/log"
	"net/http"
)

type RestAPI struct {
	Router *mux.Router
}

type RestConf struct {
	Router *mux.Router
}

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type BroadcastMessageRegistry struct {
	MessageType int
	Content     []byte
}

var HubConn = map[BroadcastKind][]*websocket.Conn{}
var LastBroadcastMessage *BroadcastMessageRegistry = nil

func NewRestAPI(conf RestConf) *RestAPI {
	r := &RestAPI{
		Router: conf.Router,
	}
	return r
}

func (h *RestAPI) RegisterHandlers() {
	h.Router.HandleFunc("/prompter/broadcast/{kind}", h.BroadcastHandler)
}

func (h *RestAPI) BroadcastHandler(w http.ResponseWriter, r *http.Request) {
	c, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Logger.Error("upgrade:", err)
		return
	}

	bKind := BroadcastKind(mux.Vars(r)["kind"])
	bWhiteList := map[BroadcastKind]bool{
		Prompt:     true,
		BibleText:  true,
		SongLyrics: true,
	}
	if _, ok := bWhiteList[bKind]; !ok {
		log.Logger.Error("broadcast client not supported")
		return
	}
	HubConn[bKind] = append(HubConn[bKind], c)

	if LastBroadcastMessage != nil {
		if err = c.WriteMessage(LastBroadcastMessage.MessageType, LastBroadcastMessage.Content); err != nil {
			log.Logger.Error(err)
		}
	}

	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			// TODO: if err equals to close event (1001, 1005) then remove connection from hub
			log.Logger.Errorf("read: %v", err)
			break
		}

		var b *Broadcast
		switch bKind {
		case BibleText:
			biblicalMessage := BibleMessage{
				Kind: BibleText,
			}
			if err = bKind.UnmarshalBroadcastMessage(message, &biblicalMessage); err != nil {
				log.Logger.Error(err)
				break
			}
			b, err = bKind.ToBroadcast(biblicalMessage)
			if err != nil {
				log.Logger.Error(err)
				break
			}
		case SongLyrics:
			songMessage := LyricsMessage{
				Kind: SongLyrics,
			}
			if err = bKind.UnmarshalBroadcastMessage(message, &songMessage); err != nil {
				log.Logger.Error(err)
				break
			}
			b, err = bKind.ToBroadcast(songMessage)
			if err != nil {
				log.Logger.Error(err)
				break
			}
		}

		if b == nil {
			break
		}

		LastBroadcastMessage = &BroadcastMessageRegistry{
			MessageType: mt,
			Content:     b.Message,
		}

		for _, prompt := range HubConn[Prompt] {
			err = prompt.WriteMessage(mt, b.Message)
			if err != nil {
				// TODO: if err equals to websocket.ErrCloseSent then remove connection from hub
				log.Logger.Errorf("write: %v", err)
			}
		}
	}
}
