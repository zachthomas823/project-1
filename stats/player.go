package stats

import (
	"strings"

	"github.com/project-0/dataframe"
)

// PlayerStats takes the given flag which is a player name and finds that player in
// the dataframe then returns a string of the entry found or returns "Couldn't find that player"
func PlayerStats(df *dataframe.Dataframe, flag string) string {
	header := df.Data[0]
	playerName := strings.ToLower(flag)
	for i := 1; i < len(df.Data); i++ {
		line := df.Data[i]
		if strings.ToLower(strings.Split(line[0], "\\")[0]) == playerName {
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
