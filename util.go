package gobot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	Users map[string]int
)

func GetUptime(startTime time.Time) string {
	uptime := time.Now().Sub(startTime)
	if uptime.Hours() > 200 {
		return "Too long.. 200+ hours"
	}

	return fmt.Sprintf(
		"%0.2d:%0.2d:%0.2d",
		int(uptime.Hours()),
		int(uptime.Minutes())%60,
		int(uptime.Seconds())%60,
	)
}

func GetUserCount(client *discordgo.Session) string {
	users := 0
	channels := 0

	servers, err := client.UserGuilds(0, "", "")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return "ERROR RETRIEVING!"
	}

	for _, server := range servers {
		m, err := client.GuildMembers(server.ID, "", 1000)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
			return "ERROR RETRIVING MEMBERS LIST!"
		}

		users += len(m)

		c, err := client.GuildChannels(server.ID)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
			return "ERROR RETRIEVING CHANNEL LIST!"
		}

		channels += len(c)
	}

	return fmt.Sprintf(
		"%d in %d channel(s) and %d server(s)",
		users,
		channels,
		len(servers),
	)
}

func GetJSONObject(url string) map[string]interface{} {
	response, err := http.Get(url)
	if err != nil {
		return nil
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	var data map[string]interface{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return nil
	}

	return data
}
