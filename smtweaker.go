package main

import (
	"github.com/flashmob/go-guerrilla"
	"github.com/flashmob/go-guerrilla/log"
	"fmt"
)

func main() {
	config := &guerrilla.AppConfig{
		Servers: []guerrilla.ServerConfig{},
		AllowedHosts: []string{},
	}
	backend := &SmtweakerBackend{}
	logger, _ := log.GetLogger("/tmp/smtweak.log")
	app, err := guerrilla.New(config, backend, logger)
	startErrors := app.Start()
	app.Shutdown()
	fmt.Print(err)
	fmt.Print(startErrors)
}
