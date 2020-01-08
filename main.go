package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	nbaStats, _ := os.Open("./nba_data.csv")
	csvReader := csv.NewReader(nbaStats)
	playerName := os.Args[1]
	contin := true
	found := false
	for contin {
		line, error := csvReader.Read()
		if error != nil {
			contin = false
		} else if strings.Split(line[1], "\\")[0] == playerName {
			fmt.Println(line)
			found = true
			contin = false
		}
	}
	if !found {
		fmt.Println("Couldn't find that player")
	}
}
