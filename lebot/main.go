package main

import (
	"DiscordBotV2"
	"DiscordBotV2/commands"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	command_map map[string]gobot.Command
)

func CreateCommand(name string, command gobot.Command) {
	if _, ok := command_map[name]; ok {
		fmt.Printf("%s already exists in the current context\n", name)
		return
	}

	command_map[name] = command
	fmt.Printf("Loaded %s\n", name)
}

func main() {
	if err := gobot.ReadConfigFromFile("../config/conf.json"); err != nil {
		panic(err)
	}

	session, err := discordgo.New("Bot " + gobot.DefaultConfig.Token)
	if err != nil {
		fmt.Printf("Failed to create a discord session: %s\n", err)
		return
	}

	//session.Debug = true

	if err := session.Open(); err != nil {
		fmt.Printf("Failed to create a websocket: %s\n", err)
		return
	}

	usr := gobot.ReadJsonFile("../config/users.json")

	if usr == nil {
		gobot.Users = make(map[string]int)
	} else {
		gobot.Users = usr
	}

	fmt.Printf("Loaded %v users\n", len(gobot.Users))

	command_map = make(map[string]gobot.Command)

	//gobot.Users["117789813427535878"] = 999 // AYY das ME!! :D

	CreateCommand("test", commands.Test{})
	CreateCommand("pleb", commands.Pleb{})
	CreateCommand("stats", commands.Stats{time.Now()})
	//CreateCommand("nominate", commands.Nominate{})
	CreateCommand("roles", commands.Roles{})
	//CreateCommand("topdev", commands.TopDev{})
	CreateCommand("dyk", commands.DYK{})
	CreateCommand("urban", commands.Urban{})
	CreateCommand("related", commands.Related{})
	CreateCommand("kick", commands.Kick{})
	CreateCommand("ban", commands.Ban{})
	CreateCommand("mc", commands.Mc{})
	CreateCommand("special", commands.Special{})
	CreateCommand("timer", commands.Timer{command_map})
	CreateCommand("help", commands.Help{command_map})
	// CreateCommand("reload", commands.Reload{command_map})

	// add handlers nd such.
	session.AddHandler(gobot.Ready)
	session.AddHandler(gobot.MessageCreate(command_map))
	session.AddHandler(gobot.GuildMemberAdd)
	session.AddHandler(gobot.GuildRoleCreate)
	session.AddHandler(gobot.GuildRoleDelete)

	fmt.Println("I am now online! Press CTRL+C to exit.")

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-sigc

	fmt.Printf("Caught signal.. Shutting down!\n")
	if err := session.Close(); err != nil {
		fmt.Println("ERROR Closing Socket!")
		os.Exit(0)
		return
	}
	fmt.Println("Socket closed.")
	os.Exit(0)
}
