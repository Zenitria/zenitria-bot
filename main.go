package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"zenitria-bot/coingecko"
	"zenitria-bot/commands"
	"zenitria-bot/config"
	"zenitria-bot/database"
	"zenitria-bot/events"

	"github.com/bwmarrin/discordgo"
)

func main() {
	database.DiscordDB = database.Connect(config.MONGODB_URI, "Discord")
	database.GetXNODB = database.Connect(config.MONGODB_URI, "Get-XNO")

	coingecko.Init()

	s, err := discordgo.New("Bot " + config.TOKEN)

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
