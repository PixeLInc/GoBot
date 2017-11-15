package gobot

import (
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type CommandArgs struct {
	Session *discordgo.Session
	Message *discordgo.MessageCreate
	Split   []string
}

type Command interface {
	GetRank() int
	Process(CommandArgs)
	Help() [2]string
}

func (args CommandArgs) SendMessage(message string) (*discordgo.Message, error) {
	return args.Session.ChannelMessageSend(args.Message.ChannelID, message)
}

func (args CommandArgs) SendMessageDelete(message string, shouldDelete bool) (*discordgo.Message, error) {
	msg, error := args.Session.ChannelMessageSend(args.Message.ChannelID, message)

	if shouldDelete {
		for {
			<-time.After(3 * time.Second)
			go args.Session.ChannelMessageDelete(msg.ChannelID, msg.ID)
		}
	}

	return msg, error
}

func (args CommandArgs) SendEmbedded(msg *discordgo.MessageEmbed) {
	args.Session.ChannelMessageSendEmbed(args.Message.ChannelID, msg)
}

func (args CommandArgs) GetReason(cArgs []string, toSkip int) string {
	if len(cArgs) < (toSkip + 1) {
		return "No reason specified"
	}

	reason := append(cArgs[:0], cArgs[toSkip:]...)

	return strings.Join(reason, " ")
}
