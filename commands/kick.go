package commands

import (
	"DiscordBotV2"
)

type Kick struct{}

func (Kick) GetRank() int {
	return 4
}

func (Kick) Process(args gobot.CommandArgs) {
	if len(args.Split) <= 2 {
		args.SendMessage("Invalid Arguments: kick <@user> <reason>")
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

	if err := args.Session.GuildMemberDeleteWithReason(c.GuildID, user.ID, reason); err != nil {
		args.SendMessage("Could not kick user!")
		return
	}

	args.SendMessage(":ok_hand:")
}

func (Kick) Help() [2]string {
	return [2]string{"Kick a user from the guild", "@user <reason>"}
}
