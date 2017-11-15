package commands

import (
	"DiscordBotV2"
	"fmt"
	"strconv"
)

type Roles struct{}

func (Roles) GetRank() int {
	return 4
}

func (Roles) Process(args gobot.CommandArgs) {
	c, err := args.Session.Channel(args.Message.ChannelID)
	if err != nil {
		args.SendMessage("Could not get channel!")
		return
	}

	roles, err := args.Session.GuildRoles(c.GuildID)
	if err != nil {
		args.SendMessage(fmt.Sprintf("An error has occurred: %v\n", err))
		return
	}

	args.SendMessage("Check console.")

	for _, role := range roles {
		fmt.Println("Role: " + role.Name + " / " + role.ID + " // " + strconv.Itoa(role.Position))
	}

}

func (Roles) Help() [2]string {
	return [2]string{"Gets server roles", ""}
}
