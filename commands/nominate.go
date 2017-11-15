package commands

import (
	"DiscordBotV2"
	"fmt"
)

type Nominate struct{}

func (Nominate) GetRank() int {
	return 0
}

const (
	exitium  = "128315785922871296"
	roleName = "Top Dev"
	roleID   = "306546714989428756"
	otherID  = "281558310082052097"
	//roleID = "281862529658126336"
)

func (Nominate) Process(args gobot.CommandArgs) {

	if len(args.Split) < 2 {
		args.SendMessage("Invalid Args: nominate <@user>")
		return
	}

	if len(args.Message.Mentions) != 1 {
		args.SendMessage("Invalid Args: You must mention 1 user!")
		return
	}

	if args.Message.Author.ID != exitium && args.Message.Author.ID != "117789813427535878" {
		args.SendMessage("Only Exitium can use this command!")
		return
	}

	user := args.Message.Mentions[0]
	if user == nil {
		args.SendMessage("User is null")
		return
	}

	c, err := args.Session.State.Channel(args.Message.ChannelID)
	if err != nil {
		args.SendMessage("Could not get channel!")
		return
	}

	// Let's grab the guild roles
	roles, err := args.Session.GuildRoles(c.GuildID)
	if err != nil {
		args.SendMessage(fmt.Sprintf("An error has occurred: %v\n", err))
		return
	}

	rID := "-1"

	// Let's grab them all :D
	st, err := args.Session.GuildMembers(c.GuildID, "", 1000)
	if err != nil {
		args.SendMessage("Failed to grab roles!")
		return //
	}

	for _, role := range roles {
		if role.ID == roleID {
			rID = role.ID
			break
		}
	}

	if rID == "-1" {
		args.SendMessage("Failed to find " + roleName + " role on server!")
		return
	}

	for _, m := range st {
		if len(m.Roles) == 0 {
			continue
		}
		for _, ro := range m.Roles {
			if ro == roleID {
				// Remove them from the rank xD
				args.Session.GuildMemberRoleRemove(c.GuildID, m.User.ID, roleID)
				args.Session.GuildMemberRoleAdd(c.GuildID, m.User.ID, otherID)
			}
		}
	}

	if roleID != "-1" { // just in case?

		if err := args.Session.GuildMemberRoleAdd(c.GuildID, user.ID, roleID); err != nil {
			args.SendMessage(fmt.Sprintf("Error giving user role: %v", err))
			return
		}

		args.SendMessage(fmt.Sprintf("<@!%s> has been given "+roleName, user.ID))

	}

}

func (Nominate) Help() [2]string {
	return [2]string{"Gives person Top Dev via Exitium's Choice.", "@user"}
}
