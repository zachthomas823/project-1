package main

import (
	"fmt"
	"net/http"

	"github.com/project-0/config"
	"github.com/project-0/stats"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if config.Player {
			fmt.Fprint(w, stats.PlayerStats())
		} else if config.Stat {
			fmt.Fprint(w, stats.StatLeader())
		} else {
			fmt.Fprint(w, "Please use tag -player \"Player Name\" \nor tag -stat \"Stat abrev\"")
		}
	})
	http.ListenAndServe(":8080", nil)
}
