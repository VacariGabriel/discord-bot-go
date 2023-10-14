package discord

import (
	"discord-bot/configs"
	"discord-bot/internal/handlers/proverbs"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var Session *discordgo.Session

func InitSession() {
	var err error
	token := configs.GetToken()
 	
	Session, err = discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		log.Fatal(err)
	}

	Session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	addHandlers()
}

func InitConnection() error {
	return Session.Open()
}

func CloseConnection() error {
	return Session.Close()
}

func addHandlers() {
	Session.AddHandler(proverbs.Proverbs)
}