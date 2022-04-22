package promptee

import (
	"github.com/k4fer74/promptee/api/rest"
	"github.com/k4fer74/promptee/bible"
)

type App struct {
	Server *rest.Server
}

type AppConf struct {
	BibleStore bible.Store
}

func NewApp(conf AppConf) *App {
	a := &App{
		Server: rest.NewServer(conf.BibleStore),
	}
	return a
}
