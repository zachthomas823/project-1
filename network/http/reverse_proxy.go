package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"

	"github.com/project-1/config"
)

func main() {
	backendPort := config.PORT
	backendURL, err := url.Parse("http://localhost:" + strconv.FormatInt(backendPort, 10))
	if err != nil {
		log.Fatal(err)
	}

	logFile, err := os.OpenFile("activity.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	proxy := httputil.NewSingleHostReverseProxy(backendURL)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Routed request to server on port", backendPort)
		proxy.ServeHTTP(w, r)
	})
	port := ":" + strconv.FormatInt(config.PROXYPORT, 10)
	fmt.Println("Reverse proxy running on " + port)
	log.Println("Reverse Proxy running on " + port)
	http.ListenAndServe(port, nil)
}
