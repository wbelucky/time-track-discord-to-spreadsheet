package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		log.Printf("failed to create discord go session: %v\n", err)
		panic(err)
	}
	discord.AddHandler(OnVoiceStateUpdate)
	if err := discord.Open(); err != nil {
		log.Printf("failed to open discord connection: %v\n", err)
		panic(err)
	}
	defer discord.Close()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

}

func OnVoiceStateUpdate(s *discordgo.Session, m *discordgo.VoiceStateUpdate) {
	log.Printf("voice state update event: %#v", m)
}
