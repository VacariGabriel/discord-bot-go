package main

import (
	"discord-bot/configs"
	"discord-bot/internal/discord"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	configs.Load()

	discord.InitSession()
	if err := discord.InitConnection(); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := discord.CloseConnection(); err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Print("Bot Started")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}