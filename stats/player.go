package stats

import (
	"fmt"
	"os"
	"strings"

	"github.com/project-0/config"
	"github.com/project-0/dataframe"
)

func PlayerStats() {
	df := dataframe.ReadCSV(config.FILE)
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
