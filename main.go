package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

var player bool

func init() {
}

func playerStats() {
	nbaStats, _ := os.Open("./nba_data.csv")
	csvReader := csv.NewReader(nbaStats)
	header, _ := csvReader.Read()
	playerName := os.Args[2]
	contin := true
	found := false
	for contin {
		line, error := csvReader.Read()
		if error != nil {
			contin = false
		} else if strings.Split(line[1], "\\")[0] == playerName {
			printPlayer(header, line)
			found = true
			contin = false
		}
	}
	if !found {
		fmt.Println("Couldn't find that player")
	}
}

func printPlayer(header []string, stats []string) {
	for i := 0; i < len(header); i++ {
		if header[i] == "Player" {
			stats[i] = strings.Split(stats[i], "\\")[0]
		}
		fmt.Println(header[i] + " - " + stats[i])
	}
}

func main() {
	flag.BoolVar(&player, "player", false, "get player statistics for name passed")
	flag.Parse()
	if player {
		playerStats()
	}
}
