package commands

import (
	"DiscordBotV2"
	"bytes"
	"fmt"
	"sort"
)

type Help struct {
	Commands map[string]gobot.Command
}

func (Help) GetRank() int {
	return 0
}

func (h Help) Process(args gobot.CommandArgs) {

	if len(args.Split) >= 2 {

		commandName := args.Split[1]

		if h.Commands[commandName] == nil {
			args.SendMessage("Invalid Command")
			return
		}

		command := h.Commands[commandName]

		usage := command.Help()[1]
		if usage == "" {
			usage = "No Usage"
		}

		args.SendMessage(fmt.Sprintf("{%v}%s - <%s> (%s)", command.GetRank(), commandName, usage, command.Help()[0]))
		return
	}

	var buffer bytes.Buffer

	uRank := 1

	if gobot.Users[args.Message.Author.ID] != 0 {
		uRank = gobot.Users[args.Message.Author.ID]
	}

	keys := make([]string, 0, len(h.Commands))
	for key := range h.Commands {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		command := h.Commands[key]
		if uRank >= command.GetRank() {
			buffer.WriteString(key + ", ")
		}
	}

	list := buffer.String()
	list = list[:len(list)-len(", ")]

	args.SendMessage("Commands: {" + list + "}")
	return
}

func (Help) Help() [2]string {
	return [2]string{"Shows the command listings", ""}
}
