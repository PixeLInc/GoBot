package commands

import (
	"DiscordBotV2"
	"github.com/bwmarrin/discordgo"
)

type DYK struct{}

func (DYK) GetRank() int {
	return 0
}

func (DYK) Process(args gobot.CommandArgs) {
	args.Session.ChannelTyping(args.Message.ChannelID)

	data := gobot.GetJSONObject("https://dev.pixelinc.tk/api/v1/dyk?key=" + gobot.DefaultConfig.ApiKey)

	if data == nil {
		args.SendMessage("PixeL's API must be down... fuckin pleb.")
		return
	}

	if data["status"].(string) == "error" {
		args.SendEmbedded(&discordgo.MessageEmbed {
			Title: "Woops, something went wrong!",
			Description: data["message"].(string),
			Color: 0xb10707,
		})
		return
	}

	text := data["message"].(string)

	if data["site"].(string) == "Did-You-Knows" {
		text = "Did you know " + text
	}

	args.SendEmbedded(&discordgo.MessageEmbed{
		Description: "**" + text + "**",
		Footer: &discordgo.MessageEmbedFooter {
			Text: data["site"].(string) + " | These facts may not be accurate",
		},
		Color: 0x00baec,
	});

	// args.SendMessage(string(data["value"].(string)))
}

func (DYK) Help() [2]string {
	return [2]string{"Grabs a random dyk fact", ""}
}
