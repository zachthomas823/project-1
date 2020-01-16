package stats

import (
	"fmt"
	"os"
	"strings"

	"github.com/project-0/config"
	"github.com/project-0/dataframe"
)

func StatLeader() {
	df := dataframe.ReadCSV(config.FILE)
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
		df.PrettyPrint()
	}
}
