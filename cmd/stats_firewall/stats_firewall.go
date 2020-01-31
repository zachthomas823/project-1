package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/project-1/config"
	"github.com/project-1/dataframe"
	"github.com/project-1/stats"
)

func main() {
	df := dataframe.ReadCSV(config.FILE) // Create a dataframe to be used
	df.DropCol(0)
	port := ":" + strconv.FormatInt(config.PORT, 10)
	port = strings.Replace(port, "\n", "", 1)
	proxyPort := "http://localhost:" + strconv.FormatInt(config.PROXYPORT, 10) + "/"
	proxyPort = strings.Replace(proxyPort, "\n", "", 1)
	http.Handle("/", http.FileServer(http.Dir("web"))) // Use the index.html for the landing page
	http.HandleFunc("/player", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header["Referer"][0])
		if r.Header["Referer"][0] == proxyPort {
			var player = r.FormValue("player_name")      // Take the response from the player name entry
			fmt.Fprint(w, stats.PlayerStats(df, player)) // Call the PlayerStats function with the df and player name
		} else {
			http.Error(w, "must use "+proxyPort, 400)
		}
	})
	http.HandleFunc("/sorted_stats", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header)
		if r.Header["Referer"][0] == proxyPort {
			var stat = r.FormValue("stat")            // Take the response from the stat selection
			fmt.Fprintln(w, stat, "\n")               // Print out the stat at the top of the page
			fmt.Fprint(w, stats.StatLeader(df, stat)) // Print out the sorted and formatted results
		} else {
			http.Error(w, "must use "+proxyPort, 400)
		}

	})
	fmt.Println("Listening on port " + port)
	http.ListenAndServe(port, nil)
}
