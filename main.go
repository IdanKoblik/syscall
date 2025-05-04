package main

import (
	"fmt"
	"os/signal"
	"syscall"
	"os"
	"github.com/bwmarrin/discordgo"
)

func main() {
	cfg, err := LoadConfig(); if err != nil {
		panic(err)
	}

	discord, err := discordgo.New("Bot " + cfg.Token)
	discord.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsGuildMembers	

	discord.AddHandler(cfg.OnMemberJoin)
	discord.AddHandler(cfg.OnStartup)

	err = discord.Open(); if err != nil {
		panic(err)
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	discord.Close()
}