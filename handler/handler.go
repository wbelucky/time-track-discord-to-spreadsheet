package handler

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/wbelucky/discord-time-track/repository"
)

type DiscordHandler struct {
	spreadsheetRepository *repository.SpreadSheetRepository
}

func NewDiscordHandler(spreadsheetRepository *repository.SpreadSheetRepository) *DiscordHandler {
	return &DiscordHandler{
		spreadsheetRepository: spreadsheetRepository,
	}
}

// ref: https://techblog.cartaholdings.co.jp/entry/archives/6412
func (d *DiscordHandler) OnVoiceStateUpdate(s *discordgo.Session, m *discordgo.VoiceStateUpdate) {
	before := m.BeforeUpdate

	if before != nil && before.ChannelID == m.ChannelID {
		return
	}

	if before != nil && before.ChannelID != m.ChannelID {
		log.Printf("user %s left from channel %s", m.Member.User.String(), before.ChannelID)
		if err := d.spreadsheetRepository.WriteEndTime(m.Member.User.String(), time.Now()); err != nil {
			log.Printf("error: failed to append to spreadsheet: %v", err)
		}
	}

	if m.ChannelID != "" && (before == nil || before.ChannelID != m.ChannelID) {
		log.Printf("user %s entered to channel %s", m.Member.User.String(), m.ChannelID)

		if err := d.spreadsheetRepository.WriteStartTime(m.Member.User.String(), time.Now()); err != nil {
			log.Printf("error: failed to append to spreadsheet: %v", err)
		}
	}
}
