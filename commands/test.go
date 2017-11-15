package commands

import (
	"DiscordBotV2"
)

type Test struct{}

func (Test) GetRank() int {
	return 5
}

func (Test) Process(args gobot.CommandArgs) {
    guilds, _ := args.Session.UserGuilds(100, "", "");

    message := "```";

    for _, guild := range guilds {
        message += guild.Name + ", ";
    }

    message += "```";

    args.SendMessage(message);
}

func (Test) Help() [2]string {
	return [2]string{"T T  Test", ""}
}
