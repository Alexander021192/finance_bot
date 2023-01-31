package main

import (
	"log"

	"github.com/Alexander021192/finance_bot/internal/cients/tg"
	"github.com/Alexander021192/finance_bot/internal/config"
	"github.com/Alexander021192/finance_bot/internal/model/msg"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal("config init failed: ", err)
	}

	tgClient, err := tg.NewBotAPI(config)
	if err != nil {
		log.Fatal("tg client init failed: ", err)
	}

	msgModel := msg.NewModel(tgClient)
	tgClient.ListenUpdates(msgModel)
}
