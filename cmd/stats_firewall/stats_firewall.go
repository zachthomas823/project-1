package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/project-1/config"
	"github.com/project-1/dataframe"
	"github.com/project-1/stats"
)

func main() {
	df := dataframe.ReadCSV(config.FILE) // Create a dataframe to be used
	df.DropCol(0)                        // Dropping the first column which is arbitrary rankings

	year, month, day := time.Now().Date()
	logFileName := "activity-" + strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day) + ".log"
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	port := ":" + strconv.FormatInt(config.PORT, 10)
	port = strings.Replace(port, "\n", "", 1)

	proxyPort := "http://localhost:" + strconv.FormatInt(config.PROXYPORT, 10) + "/"
	proxyPort = strings.Replace(proxyPort, "\n", "", 1)

	http.Handle("/", http.FileServer(http.Dir("web"))) // Use the index.html for the landing page
	log.Println("Server running on " + port)

	http.HandleFunc("/player", func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Referer"][0] == proxyPort {
			var player = r.FormValue("player_name")      // Take the response from the player name entry
			fmt.Fprint(w, stats.PlayerStats(df, player)) // Call the PlayerStats function with the df and player name
			log.Println(player + " searched for on the player endpoint by " + proxyPort)
		} else {
			log.Println("Access denied to " + r.Header["Referer"][0] + " on player endpoint")
			http.Error(w, "must use "+proxyPort, 400)
		}
	})

	http.HandleFunc("/sorted_stats", func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Referer"][0] == proxyPort {
			var stat = r.FormValue("stat")            // Take the response from the stat selection
			fmt.Fprintln(w, stat, "\n")               // Print out the stat at the top of the page
			fmt.Fprint(w, stats.StatLeader(df, stat)) // Print out the sorted and formatted results
			log.Println(stat + " searched for on the sorted_stats endpoint by " + proxyPort)
		} else {
			log.Println("Access denied to " + r.Header["Referer"][0] + " on sorted_stats endpoint")
			http.Error(w, "must use "+proxyPort, 400)
		}

	})
	fmt.Println("Listening on port " + port)
	http.ListenAndServe(port, nil)
}
