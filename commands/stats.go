package commands

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"time"

	"DiscordBotV2"

	"github.com/bwmarrin/discordgo"
)

type Stats struct {
	StartTime time.Time
}

func (Stats) GetRank() int {
	return 0
}

func (s Stats) Process(args gobot.CommandArgs) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	rand.Seed(time.Now().Unix())

	args.SendEmbedded(&discordgo.MessageEmbed{
		Title: "-= Bot Stats -=",
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:  "Memory Used",
				Value: fmt.Sprintf("%.2f Mb", float64(m.Alloc)/1000000),
			},
			&discordgo.MessageEmbedField{
				Name:   "Users In Touch",
				Value:  gobot.GetUserCount(args.Session),
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:  "Up-Time",
				Value: gobot.GetUptime(s.StartTime),
			},
			&discordgo.MessageEmbedField{
				Name:  "Go Version",
				Value: runtime.Version(),
			},
			&discordgo.MessageEmbedField{
				Name:  "Concurrent Tasks",
				Value: strconv.Itoa(runtime.NumGoroutine()),
			},
			&discordgo.MessageEmbedField{
				Name:  "Developer",
				Value: "PixeL <3",
			},
			/*&discordgo.MessageEmbedField{ // From before the re-write. Still haven't implemented this yet.
				Name:  "Cached Messages",
				Value: fmt.Sprintf("%v/%v", len(message_cache), cache_limit),
			},
			&discordgo.MessageEmbedField{
				Name:  "Cache Reset",
				Value: GetTimeRemaining(),
			},*/
		},
		Color: 0xff00c1, //colors[rand.Intn(len(colors))],
	})
}

func (Stats) Help() [2]string {
	return [2]string{"Shows Bot Statistics", ""}
}
