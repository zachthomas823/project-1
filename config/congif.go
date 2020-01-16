package config

import "flag"

// FILE is where to find the data to be parsed
const FILE = "./nba_data.csv"

var Player bool
var Stat bool

func init() {
	flag.BoolVar(&Player, "player", false, "get player statistics for name passed")
	flag.BoolVar(&Stat, "stat", false, "get the leaders for an individual statistic")
	flag.Parse()
}
