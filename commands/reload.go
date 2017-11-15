package commands

import (
    "DiscordBotV2"
    "strings"
)

type Reload struct {
    Commands map[string]gobot.Command
}

func (Reload) GetRank() int {
    return 6
}

func (r Reload) Process(args gobot.CommandArgs) {
    if len(args.Split) != 2 {
        args.SendMessage("Invalid Usage: reload <command_name>")
        return
    }

    functionName := strings.ToLower(args.Split[1])

    if r.Commands[functionName] == nil {
        args.SendMessage("Failed to reload command: The specified command does not exist.")
        return
    }

    r.Commands["test"] = Test{}
    args.SendMessage("Reloaded.")
}

func (Reload) Help() [2]string {
    return [2]string{"Reloads a command", "<command_name>"}
}
