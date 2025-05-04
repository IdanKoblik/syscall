package main

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func (config Config) OnMemberJoin(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	user := m.User

	description := fmt.Sprintf("- remember to read the <#%s>\n- [link to server](%s)",
	 config.RulesChannel,
	 config.ServerURL,
	)

	welcome := fmt.Sprintf("Welcome, <@%s>!", user.ID)
	_, err := s.ChannelMessageSend(config.WelcomeChannel, welcome); if err != nil {
		log.Println("Error sending welcome text:", err)
		return
	}

	title := fmt.Sprintf("Welcome to %s!", config.Name)
	avatar := fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.png", user.ID, user.Avatar)
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    user.GlobalName,
			IconURL: avatar,
		},
		Title: title,
		Description: description,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: avatar,
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	_, err = s.ChannelMessageSendEmbed(config.WelcomeChannel, embed); if err != nil {
		log.Println("Error sending welcome embed:", err)
	}
}
