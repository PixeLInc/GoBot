package commands

import (
	"DiscordBotV2"
)

type Pleb struct{}

func (Pleb) GetRank() int {
	return 0
}

func (Pleb) Process(args gobot.CommandArgs) {
	args.SendMessage("Plebby is fucking gay.")
}

func (Pleb) Help() [2]string {
	return [2]string{"says Plebby is fucking gay", ""}
}
