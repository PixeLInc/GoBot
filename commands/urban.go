package commands

import (
	"DiscordBotV2"
	"strconv"
	"github.com/bwmarrin/discordgo"
)

type Urban struct{}

func (Urban) GetRank() int {
	return 0
}

func (Urban) Process(args gobot.CommandArgs) {
	if len(args.Split) != 2 {
		args.SendMessage("Invalid Usage, urban <word>, No spaces or underscores.")
		return
	}

	args.Session.ChannelTyping(args.Message.ChannelID)

	data := gobot.GetJSONObject("http://api.urbandictionary.com/v0/define?term=" + args.Split[1])

	if data == nil {
		args.SendMessage("Error!")
		return
	}

	if data["result_type"].(string) != "exact" {
		args.SendMessage("Failed to find that term!")
		return
	}

	mp := data["list"].([]interface{})[0].(map[string]interface{})

	args.SendEmbedded(&discordgo.MessageEmbed {
		Title: "Definition for " + args.Split[1],
		Description: mp["definition"].(string),
		Fields: []*discordgo.MessageEmbedField {
			&discordgo.MessageEmbedField {
				Name: "Example",
				Value: mp["example"].(string),
			},
			&discordgo.MessageEmbedField {
				Name:  "Rating",
				Value: strconv.FormatFloat(mp["thumbs_up"].(float64), 'f', -1, 32) + " 👍 | " + strconv.FormatFloat(mp["thumbs_down"].(float64), 'f', -1, 32) + " 👎",
			},
		},
		Footer: &discordgo.MessageEmbedFooter {
			Text: "Powered by Urban Dictionary",
		},
		Color: 0x00fff3,
	});

	// args.SendMessage("```" + mp["definition"].(string) + "\n\nExample: " + mp["example"].(string) + "\n\n" + strconv.FormatFloat(mp["thumbs_up"].(float64), 'f', -1, 32) + " 👍 | " + strconv.FormatFloat(mp["thumbs_down"].(float64), 'f', -1, 32) + " 👎```")
}

func (Urban) Help() [2]string {
	return [2]string{"Grabs the urban dictionary definition of a word", "word"}
}
