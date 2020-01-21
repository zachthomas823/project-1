package main

import (
	"fmt"
	"net/http"

	"github.com/project-0/stats"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("web")))
	http.HandleFunc("/player", func(w http.ResponseWriter, r *http.Request) {
		var player = r.FormValue("player_name")
		fmt.Fprint(w, stats.PlayerStats(player))
	})
	http.HandleFunc("/pts", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, stats.StatLeader("PTS"))
	})
	http.HandleFunc("/ast", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, stats.StatLeaderJSON("AST"))
	})
	http.HandleFunc("/sorted_stats", func(w http.ResponseWriter, r *http.Request) {
		var stat = r.FormValue("stat")
		fmt.Fprintln(w, stat)
		fmt.Fprint(w, stats.StatLeader(stat))
	})
	http.ListenAndServe(":8080", nil)
}
