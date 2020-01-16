package stats

import (
	"os"
	"strings"

	"github.com/project-0/config"
	"github.com/project-0/dataframe"
)

func PlayerStats() string {
	df := dataframe.ReadCSV(config.FILE)
	df.DropCol(0)
	header := df.Data[0]
	playerName := os.Args[2]
	for i := 1; i < len(df.Data); i++ {
		line := df.Data[i]
		if strings.Split(line[0], "\\")[0] == playerName {
			return printPlayer(header, line)
		}
	}
	return ("Couldn't find that player")
}

func printPlayer(header []string, stats []string) string {
	var result string
	for i := 0; i < len(header); i++ {
		if header[i] == "Player" {
			stats[i] = strings.Split(stats[i], "\\")[0]
		}
		result = result + header[i] + " - " + stats[i] + "\n"
	}
	return result
}