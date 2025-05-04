package main

import (
	"fmt"
	"encoding/json"
	"os"
	"github.com/bwmarrin/discordgo"
)

var raw map[string]interface{}

func createMessageComponents(item interface{}, name string) ([]discordgo.MessageComponent, *discordgo.MessageEmbed) {
	var components []discordgo.MessageComponent
	var embed *discordgo.MessageEmbed

	switch v := item.(type) {
	case Partner:
		if v.Website != "" {
			components = append(components, discordgo.Button{
				Label: "Visit Website",
				Style: discordgo.LinkButton,
				URL:   v.Website,
			})
		}

		if v.Discord != "" {
			components = append(components, discordgo.Button{
				Label:    "Discord",
				Style:    discordgo.PrimaryButton,
				CustomID: "more_info_" + v.Discord,
			})
		}

		embed = &discordgo.MessageEmbed{
			Title:       fmt.Sprintf("%s 🤝 %s", name,v.Name),
			Description: v.Description,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: v.Image,
			},
		}

	case Product:
		if v.Website != "" {
			components = append(components, discordgo.Button{
				Label: "Visit Website",
				Style: discordgo.LinkButton,
				URL:   v.Website,
			})
		}

		embed = &discordgo.MessageEmbed{
			Title:       fmt.Sprintf("Product: %s", v.Name),
			Description: v.Description,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: v.Image,
			},
		}
	}

	return components, embed
}

func handleRules(s *discordgo.Session, channel, rules string) {
	embed := &discordgo.MessageEmbed {
		Title: "Server Rules & Guidelines",
		Description: rules,
	}

	_, err := s.ChannelMessageSendEmbed(channel, embed); if err != nil {
		fmt.Println("Failed to send server rules: ", err)
		return 
	}
}

func (config *Config) OnStartup(s *discordgo.Session, m *discordgo.Ready) {
	data, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println("Failed to read config.json:", err)
		return
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		fmt.Println("Failed to parse config.json:", err)
		return
	}

	if config.Rules.Send {
		handleRules(s, config.Rules.Id, config.Rules.Rules)

		if rules, ok := raw["rules"].(map[string]interface{}); ok {
			rules["send"] = false
		}
	}

	partnersArray, ok := raw["partners"].(map[string]interface{})["partners"].([]interface{})
	if !ok {
		fmt.Println("Invalid partner structure in config.json")
		return
	}

	for i, partner := range config.Partners.Partners {
		if !partner.Send {
			continue
		}

		components, embed := createMessageComponents(partner, config.Guild)

		_, err := s.ChannelMessageSendComplex(config.Partners.Id, &discordgo.MessageSend{
			Embed: embed,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: components,
				},
			},
		})

		if err != nil {
			fmt.Println("Failed to send partner message:", err)
			continue
		}

		if i < len(partnersArray) {
			entry := partnersArray[i].(map[string]interface{})
			entry["send"] = false
		}
	}

	productsArray, ok := raw["products"].(map[string]interface{})["products"].([]interface{})
	if !ok {
		fmt.Println("Invalid product structure in config.json")
		return
	}

	for i, product := range config.Products.Products {
		if !product.Send {
			continue
		}

		components, embed := createMessageComponents(product, config.Name)

		_, err := s.ChannelMessageSendComplex(config.Products.Id, &discordgo.MessageSend{
			Embed: embed,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: components,
				},
			},
		})

		if err != nil {
			fmt.Println("Failed to send product message:", err)
			continue
		}

		if i < len(productsArray) {
			entry := productsArray[i].(map[string]interface{})
			entry["send"] = false
		}
	}

	updatedData, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		fmt.Println("Failed to marshal updated config.json:", err)
		return
	}

	if err := os.WriteFile("config.json", updatedData, 0644); err != nil {
		fmt.Println("Failed to write config.json:", err)
	}
}

