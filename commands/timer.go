package commands

import (
	"DiscordBotV2"
	"fmt"
	"strings"
	"time"
)

type Timer struct {
	Commands map[string]gobot.Command
}

func (Timer) GetRank() int {
	return 2
}

func (t Timer) Process(args gobot.CommandArgs) {
	if len(args.Split) < 2 {
		args.SendMessage("Invalid Usage: timer <command_name> {command args}")
		return
	}

	functionName := strings.ToLower(args.Split[1])

	if t.Commands[functionName] == nil {
		args.SendMessage("Invalid Command!")
		return
	}

	if functionName == "timer" { // Does a StackOverflow
		args.SendMessage("No.")
		return
	}

	command := t.Commands[functionName]

	args.Split = append(args.Split[:0], args.Split[1:]...) // removes the first split {Eg. Hey dude Hi, removes Hey (dude Hi)}

	start := time.Now()

	command.Process(gobot.CommandArgs{args.Session, args.Message, args.Split})

	endTime := time.Now()

	args.SendMessage(fmt.Sprintf("``Completed in %v (%v second(s))``", endTime.Sub(start), endTime.Sub(start).Seconds()))
}

func (Timer) Help() [2]string {
	return [2]string{"Times a command", ""}
}
