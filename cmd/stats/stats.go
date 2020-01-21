package main

import (
	"fmt"
	"os"

	"github.com/project-0/dataframe"

	"github.com/project-0/config"
	"github.com/project-0/stats"
)

func main() {
	df := dataframe.ReadCSV(config.FILE)
	df.DropCol(0)
	if config.Player {
		fmt.Println(stats.PlayerStats(df, os.Args[2]))
	} else if config.Stat {
		fmt.Println(stats.StatLeader(df, os.Args[2]))
	} else {
		fmt.Println("Please use tag -player \"Player Name\" \nor tag -stat \"Stat abrev\"")
	}
}
