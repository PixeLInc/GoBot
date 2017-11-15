package commands

import (
	"DiscordBotV2"
)

type Ban struct{}

func (Ban) GetRank() int {
	return 5
}

func (Ban) Process(args gobot.CommandArgs) {
	if len(args.Split) <= 2 {
		args.SendMessage("Invalid Arguments: ban <@user> <reason>")
		return
	}

	if len(args.Message.Mentions) != 1 {
		args.SendMessage("Invalid Args: You must mention 1 user!")
		return
	}

	user := args.Message.Mentions[0]
	if user == nil {
		args.SendMessage("User is null")
		return
	}

	reason := args.GetReason(args.Split, 2)

	c, err := args.Session.State.Channel(args.Message.ChannelID)
	if err != nil {
		args.SendMessage("Could not get channel!")
		return
	}

	if err := args.Session.GuildBanCreateWithReason(c.GuildID, user.ID, reason, 0); err != nil {
		args.SendMessage("Could not ban user!")
		return
	}

	args.SendMessage(":ok_hand:")
}

func (Ban) Help() [2]string {
	return [2]string{"Ban someone", "@user <reason>"}
}
