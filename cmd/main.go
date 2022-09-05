package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/wbelucky/discord-time-track/handler"
	"github.com/wbelucky/discord-time-track/repository"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		log.Printf("failed to create discord go session: %v\n", err)
		panic(err)
	}

	// TODO: commandを追加 https://github.com/bwmarrin/discordgo/blob/master/examples/slash_commands/main.go

	s, err := repository.NewSpreadSheetRepository()
	if err != nil {
		panic(err)
	}
	d := handler.NewDiscordHandler(s)
	discord.AddHandler(d.OnVoiceStateUpdate)
	if err := discord.Open(); err != nil {
		log.Printf("failed to open discord connection: %v\n", err)
		panic(err)
	}
	log.Println("discord bot started successfully")

	defer discord.Close()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	log.Println("closing discord bot server")

}
