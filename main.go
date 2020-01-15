package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
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
	header := df.Data[0]
	playerName := os.Args[2]
	found := false
	for i := 1; i < len(df.Data); i++ {
		line := df.Data[i]
		if strings.Split(line[1], "\\")[0] == playerName {
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

type statName struct {
	name string
	stat float64
}

func statLeader() {
	nbaStats, _ := os.Open("./nba_data.csv")
	csvReader := csv.NewReader(nbaStats)
	header, _ := csvReader.Read()
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
		var allStat []statName
		contin := true
		for contin {
			line, err := csvReader.Read()
			if err != nil {
				contin = false
			} else {
				name := strings.Split(line[1], "\\")[0]
				stat, _ := strconv.ParseFloat(line[statIdx], 64)
				allStat = append(allStat, statName{name, stat})
			}
		}
		sortedStats := statSorter(allStat)
		printStatLeader(sortedStats)
	}
}

func statSorter(statNames []statName) []statName {
	sortedNames := []statName{statNames[0]}
	for i := 1; i < len(statNames); i++ {
		stat := statNames[i].stat
		placed := false
		for j := 0; j < len(sortedNames); j++ {
			if stat > sortedNames[j].stat {
				sortedNames = append(sortedNames, sortedNames[len(sortedNames)-1])
				for k := len(sortedNames) - 1; k > j; k-- {
					sortedNames[k] = sortedNames[k-1]
				}
				sortedNames[j] = statNames[i]
				placed = true
				break
			}
		}
		if !placed {
			sortedNames = append(sortedNames, statNames[i])
		}
	}
	return sortedNames
}

func printStatLeader(sn []statName) {
	for i := 0; i < len(sn); i++ {
		fmt.Println(i+1, ":", sn[i].name, "\t", sn[i].stat)
	}
}
