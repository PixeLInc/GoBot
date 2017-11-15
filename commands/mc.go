package commands

import (
	"DiscordBotV2"
	"DiscordBotV2/mc"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"net"
	"strconv"
	"strings"
	"time"
)

type Mc struct{}

func (Mc) GetRank() int {
	return 2
}

func (Mc) Process(args gobot.CommandArgs) {
	if len(args.Split) != 2 {
		args.SendMessage("Invalid Args: mc <ip:port> (Example: mc.hypixel.net:25565)")
		return
	}

	args.Session.ChannelTyping(args.Message.ChannelID)

	message, err := args.SendMessage("Gathering info...")
	if err != nil {
		fmt.Println("Can't send the message>?")
		return
	}

	host := args.Split[1]

	if !strings.Contains(host, ":") {
		args.Session.ChannelMessageEdit(args.Message.ChannelID, message.ID, "Invalid Server, ex: mc.hypixel.net:25565")
		return
	}

	servConn, err := net.DialTimeout("tcp", host, time.Duration(5)*time.Second)
	if err != nil {
		args.Session.ChannelMessageEdit(args.Message.ChannelID, message.ID, "Failed to connect: Is it up?")
		return
	}

	fmt.Println("Connected to the server")

	pong, err := mc.PingServer(servConn, host)
	if err != nil {
		fmt.Println(err)
		args.Session.ChannelMessageEdit(args.Message.ChannelID, message.ID, "Failed to ping server: Read Error?")
		return
	}

	//fmt.Println(svr.name, pong.Players.Online, pong.Players.Max, pong.Description)

	//fmt.Println(strings.Split(pong.FavIcon, ",")[1])

	//BaseToImage(strings.Split(pong.FavIcon, ",")[1]) (SOON)

	serverDescription := "- Cant Get Description Yet! (Its full of colours which break things) -"
	switch v := pong.Description.(type) {
	case string:
		serverDescription = v
	}

	args.SendEmbedded(&discordgo.MessageEmbed{
		Title: "- Server Info -",
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:  "Server Description",
				Value: serverDescription,
			},
			&discordgo.MessageEmbedField{
				Name:  "Online Players",
				Value: strconv.Itoa(pong.Players.Online),
			},
			&discordgo.MessageEmbedField{
				Name:  "Max Players",
				Value: strconv.Itoa(pong.Players.Max),
			},
			&discordgo.MessageEmbedField{
				Name:  "Protocol",
				Value: pong.Version.Name,
			},
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://www.fruitiondayspa.ca/wp-content/uploads/profile-photo-coming-soon.jpg",
		},
		Color: 0x00b9ff,
	})

	if err := servConn.Close(); err != nil {
		fmt.Println("FAILED TO DISCONNNECT FROM THE SERVER")
	}
}

func (Mc) Help() [2]string {
	return [2]string{"Gives you info about a Minecraft Server", "<ip:port>"}
}
