package promptee

import (
	"github.com/k4fer74/promptee/pkg/api/rest"
	"github.com/k4fer74/promptee/pkg/bible"
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
