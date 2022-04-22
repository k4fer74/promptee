package main

import (
	"github.com/k4fer74/promptee"
	"github.com/k4fer74/promptee/bible/osis"
)

func main() {
	bible, err := osis.New("artifacts/bible/osis/pt.xml")
	if err != nil {
		panic(err)
	}
	app := promptee.NewApp(promptee.AppConf{
		BibleStore: bible,
	})
	if err = app.Server.ListenAndServe(); err != nil {
		panic(err)
	}
}
