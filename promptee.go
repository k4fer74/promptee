package promptee

import (
	"github.com/k4fer74/promptee/pkg/api/rest"
	"github.com/k4fer74/promptee/pkg/bible"
	"github.com/k4fer74/promptee/pkg/lyrics"
)

type App struct {
	Server *rest.Server
}

type AppConf struct {
	BibleStore     bible.Store
	LyricsRegistry lyrics.Registry
}

func NewApp(conf AppConf) *App {
	app := &App{
		Server: rest.NewServer(conf.BibleStore, conf.LyricsRegistry),
	}
	return app
}
