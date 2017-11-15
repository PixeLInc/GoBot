package gobot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const (
	main_guild  = "241582286204567552"
	log_channel = "299653045699084289"

	chris         = "100430116991041536"
	exitiumServer = "306527847130857474"
)

func MessageCreate(commands map[string]Command) func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(s *discordgo.Session, msg *discordgo.MessageCreate) {

		defer func() { // just in case.
			if r := recover(); r != nil {
				fmt.Println("Recovered from a panic!!")
				fmt.Println(r)
			}
		}()

		if msg.Author.ID == s.State.User.ID {
			return
		}

		c, err := s.Channel(msg.ChannelID)

		if err != nil {
			fmt.Println("Error on get channel: ", err)
			return
		}

		gName := "PM"

		g, err := s.Guild(c.GuildID)
		if err == nil {
			gName = g.Name
		}

		content := msg.Content

		fmt.Printf("(%v{%v})%v > %v\n", gName, c.Name, msg.Author.Username, content)

		if strings.HasPrefix(content, DefaultConfig.Prefix) {

			//TODO: Remove dis
			// if msg.Author.ID != DefaultConfig.OwnerID {
			// 	s.ChannelMessageSend(msg.ChannelID, "| Sorry! The bot is currently under-going a full code re-write! |")
			// 	return
			// }

			sub := content[1:len(content)]
			split := strings.Split(sub, " ")

			fmt.Printf("Command message: %s\n", sub)

			cmd := strings.ToLower(split[0])

			if commands[cmd] == nil {
				//s.ChannelMessageSend(msg.ChannelID, "Sorry, that command does not exist! Type ~help for a list of commands")
				return
			}

			rank := commands[cmd].GetRank()

			if rank != 0 && Users[msg.Author.ID] == 0 { // Not registered
				s.ChannelMessageSend(msg.ChannelID, "Sorry, Guests cannot use that command!")
				return
			}

			if Users[msg.Author.ID] != 0 && rank > 1 {
				uRank := Users[msg.Author.ID]

				if rank > uRank {
					s.ChannelMessageSend(msg.ChannelID, "Hey, you don't have the correct permissions to execute that command!")
					return
				}
			}

			commands[cmd].Process(CommandArgs{
				Message: msg,
				Session: s,
				Split:   split,
			})
		}
	}
}

func GuildMemberAdd(s *discordgo.Session, event *discordgo.GuildMemberAdd) {
	if event.Member.GuildID == main_guild { // If they're joining our guild
		s.ChannelMessageSend(log_channel, fmt.Sprintf("Welcome to the server, <@!%s>", event.Member.User.ID))
	} else {
		oGuild, err := s.State.Guild(event.Member.GuildID)
		if err != nil {
			fmt.Println("Failed to log user join in another guild!")
			return
		}

		if oGuild.ID == exitiumServer {
			if event.Member.User.ID == chris {
				s.GuildBanCreate(oGuild.ID, chris, 0)
				fmt.Println("Chris was banned :D")
				return
			}
		}

		s.ChannelMessageSend(log_channel, fmt.Sprintf("%s joined %s", event.Member.User.Username, oGuild.Name))
	}
}

func GuildRoleCreate(s *discordgo.Session, event *discordgo.GuildRoleCreate) {
	oGuild, err := s.State.Guild(event.GuildRole.GuildID)
	if err != nil {
		fmt.Println("Failed to log role create.")
		return
	}

	fmt.Printf("Role Create in %s: %s\n", oGuild.Name, event.GuildRole.Role.Name)
}

func GuildRoleDelete(s *discordgo.Session, event *discordgo.GuildRoleDelete) {
	oGuild, err := s.State.Guild(event.GuildID)
	if err != nil {
		fmt.Println("Failed to log role create.")
		return
	}

	fmt.Printf("Role delete in %s: ID->%s\n", oGuild.Name, event.RoleID)
}

func Ready(s *discordgo.Session, event *discordgo.Ready) {
	if err := s.UpdateStatus(0, "Type ~help to use me!"); err != nil {
		fmt.Println("ERROR SETTING STATUS!!")
	}
	fmt.Println("Set update status")

	fmt.Printf("I am in %v guilds\n", len(event.Guilds))

	if s.Debug {
		for _, guild := range event.Guilds {
			fmt.Println(guild.Name)
		}
	}
}
