package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	nbaStats, _ := os.Open("./nba_data.csv")
	csvReader := csv.NewReader(nbaStats)
	for i := 0; i < 2; i++ {
		line, _ := csvReader.Read()
		// if line[1]
		fmt.Println(line[1])
	}
}
