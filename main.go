package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

var player bool
var stat bool

func init() {
}

func main() {
	flag.BoolVar(&player, "player", false, "get player statistics for name passed")
	flag.BoolVar(&stat, "stat", false, "get the leaders for an individual statistic")
	flag.Parse()
	if player {
		playerStats()
	} else if stat {
		statLeader()
	}
}

func playerStats() {
	nbaStats, _ := os.Open("./nba_data.csv")
	csvReader := csv.NewReader(nbaStats)
	header, _ := csvReader.Read()
	fmt.Println(header)
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

func statLeader() {
	nbaStats, _ := os.Open("./nba_data.csv")
	csvReader := csv.NewReader(nbaStats)
	header, _ := csvReader.Read()
	fmt.Println(header)
	stat := os.Args[2]
	var statIdx int
	var found bool
	for i := 0; i < len(header); i++ {
		if header[i] == stat {
			statIdx = i
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Couldn't find that stat")
	} else {
		var allStat []string
		var names []string
		contin := true
		for contin {
			line, err := csvReader.Read()
			if err != nil {
				contin = false
			} else {
				allStat = append(allStat, line[statIdx])
				names = append(names, strings.Split(line[1], "\\")[0])
			}
		}
		printStatLeader(names, allStat)
	}
}

func printStatLeader(names []string, stats []string) {
	for i := 0; i < len(names); i++ {
		fmt.Println(i+1, ":", names[i], "\t", stats[i])
	}
}
