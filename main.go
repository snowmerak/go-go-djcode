package main

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token := os.Getenv("MOCHA_TOKEN")
	if token == "" {
		panic("MOCHA_TOKEN not set")
	}

	session, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}

}
