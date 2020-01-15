package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/project-0/dataframe"
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
	df := dataframe.ReadCSV("./nba_data.csv")
	df.DropCol(0)
	header := df.Data[0]
	playerName := os.Args[2]
	found := false
	for i := 1; i < len(df.Data); i++ {
		line := df.Data[i]
		if strings.Split(line[0], "\\")[0] == playerName {
			printPlayer(header, line)
			found = true
			break
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
	df := dataframe.ReadCSV("./nba_data.csv")
	df.DropCol(0)
	for i := 1; i < len(df.Data); i++ {
		df.Data[i][0] = strings.Split(df.Data[i][0], "\\")[0]
	}
	header, _ := df.Data[0]
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
		df.Sort(statIdx)
		fmt.Println(df.Data[0])
		for i := 1; i < len(df.Data); i++ {
			fmt.Println(i, df.Data[i])
		}
	}
}
