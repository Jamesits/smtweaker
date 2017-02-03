package main

import (
	"fmt"
	"github.com/flashmob/go-guerrilla/envelope"
	"github.com/flashmob/go-guerrilla/backends"
)

type SmtweakerBackendConfig struct {
	LogReceivedMails bool `json:"log_received_mails"`
}

type SmtweakerBackend struct {
	config SmtweakerBackendConfig
	// embed functions form AbstractBackend so that DummyBackend satisfies the Backend interface
	backends.AbstractBackend
}

func (cb *SmtweakerBackend) Process(c *envelope.Envelope) backends.BackendResult {
	// err := saveSomewhere(c.Data)
	fmt.Print(c.Data)
	//if err != nil {
	//	return backends.NewBackendResult(fmt.Sprintf("554 Error: %s", err.Error()))
	//}
	return backends.NewBackendResult("250 OK")
}