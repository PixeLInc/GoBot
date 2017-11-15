package commands

import (
	"DiscordBotV2"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Related struct{}

func (Related) GetRank() int {
	return 0
}

func (Related) Process(args gobot.CommandArgs) {
	if len(args.Split) != 2 {
		args.SendMessage("Invalid Usage: related <word>, Use underscores for spaces.")
		return
	}

	word := strings.ToLower(args.Split[1])

	args.Session.ChannelTyping(args.Message.ChannelID)

	data := gobot.GetJSONObject("http://api.conceptnet.io/related/c/en/" + word + "?filter=/c/en")
	if data == nil {
		args.SendMessage("Error!")
		return
	}

	terms := make(map[string]float64)

	limit := 10

	// Example -> K: 41 v: map[@id:/c/en/lie_detector weight:0.444]

	for k, _ := range data["related"].([]interface{}) {
		if k >= limit {
			break
		}

		stff := data["related"].([]interface{})[k].(map[string]interface{})

		term := strings.Replace(stff["@id"].(string), "/c/en/", "", -1)
		weight := stff["weight"].(float64)
		//fmt.Printf(term + ": " + weight + " \n");

		terms[term] = weight
	}

	messageTerms := ""

	n := map[float64][]string{}
	var a []float64
	for k, v := range terms {
		n[v] = append(n[v], k)
	}
	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(a)))

	for _, k := range a {
		for _, s := range n[k] {
			messageTerms += s + " has a " + strconv.FormatFloat(k*100, 'f', -1, 32) + "%% match\n"
		}
	}

	if messageTerms == "" {
		messageTerms = "No related terms found!"
	}

	message := fmt.Sprintf("```Related Terms:\n" +
		messageTerms + "```")

	args.SendMessage(message)
}

func (Related) Help() [2]string {
	return [2]string{"Trys to relate terms", "term"}
}
