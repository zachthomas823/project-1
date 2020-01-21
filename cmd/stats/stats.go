package main

import (
	"fmt"
	"os"

	"github.com/project-0/config"
	"github.com/project-0/stats"
)

func main() {
	if config.Player {
		fmt.Println(stats.PlayerStats(os.Args[2]))
	} else if config.Stat {
		fmt.Println(stats.StatLeader(os.Args[2]))
	} else {
		fmt.Println("Please use tag -player \"Player Name\" \nor tag -stat \"Stat abrev\"")
	}
}
