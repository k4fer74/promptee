package main

import (
	"github.com/k4fer74/promptee"
	"github.com/k4fer74/promptee/pkg/bible/osis"
	"github.com/k4fer74/promptee/pkg/log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal, 1)

	bible, err := osis.New("artifacts/bible/osis/pt.xml")
	if err != nil {
		panic(err)
	}

	app := promptee.NewApp(promptee.AppConf{
		BibleStore: bible,
	})

	go func() {
		log.Logger.Info("[*] Started REST API at 127.0.0.1:3160")
		if err := app.Server.StartAPI(); err != http.ErrServerClosed {
			log.Logger.Fatal(err)
		}
	}()

	signal.Notify(c, os.Interrupt)

	<-c

	if err := app.Server.Stop(); err != nil {
		log.Logger.Error(err)
	}

	os.Exit(0)
}
