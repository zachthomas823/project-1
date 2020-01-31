package config

import "flag"

// FILE is where to find the data to be parsed
const FILE = "./nba_data.csv"

// PORT is the port on which the server will run
const PORT int64 = 6969

const PROXYPORT int64 = 42069

// Player flag for if the user wants an individual player
var Player bool

// Stat flag for if the user wants Stats for every player
var Stat bool

func init() {
	flag.BoolVar(&Player, "player", false, "get player statistics for name passed")
	flag.BoolVar(&Stat, "stat", false, "get the leaders for an individual statistic")
	flag.Parse()
}
