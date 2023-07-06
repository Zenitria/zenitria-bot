package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"zenitria-bot/coingecko"
	"zenitria-bot/commands"
	"zenitria-bot/database"
	"zenitria-bot/events"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	TOKEN   string
	MONGODB string
)

func main() {
	godotenv.Load()

	TOKEN = os.Getenv("TOKEN")
	MONGODB = os.Getenv("MONGODB")

	database.DiscordDB = database.Connect(MONGODB, "Discord")
	database.GetXNODB = database.Connect(MONGODB, "Get-XNO")

	coingecko.Init()

	s, err := discordgo.New("Bot " + TOKEN)

	if err != nil {
		fmt.Print(err.Error())
		return
	}

	s.AddHandler(events.Ready)
	s.AddHandler(events.InteractionCreate)
	s.AddHandler(events.MessageCreate)

	bot := s.Open()

	if bot != nil {
		fmt.Print(bot.Error())
		return
	}

	commands.RegisterCommands(s)

	defer func() {
		s.Close()
		database.Disconnect(database.DiscordDB)
		database.Disconnect(database.GetXNODB)
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGTERM)
	<-sc
}
