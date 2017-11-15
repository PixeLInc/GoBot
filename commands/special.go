package commands

import (
	"DiscordBotV2"
)

const myServer = "314933413847105537"

type Special struct{}

func (Special) GetRank() int {
	return 0
}

func (Special) Process(args gobot.CommandArgs) {
	c, err := args.Session.State.Channel(args.Message.ChannelID)
	if err != nil {
		args.SendMessage("Could not get channel!")
		return
	}

	if c.GuildID != myServer {
		args.SendMessage("Can't do that in this server ;3")
		return
	}

	args.Session.GuildMemberRoleAdd(c.GuildID, args.Message.Author.ID, "315559247863873536")
	args.SendMessage("You are now *SPECIALLLLLL*")
}

func (Special) Help() [2]string {
	return [2]string{"Wanna be speciallll ;)? Type ME!", ""}
}
