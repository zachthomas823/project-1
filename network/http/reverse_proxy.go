package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"

	"github.com/project-1/config"
)

func main() {
	backendPort := config.PORT
	backendURL, err := url.Parse("http://localhost:" + strconv.FormatInt(backendPort, 10))
	if err != nil {
		log.Fatal(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(backendURL)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})
	port := ":" + strconv.FormatInt(config.PROXYPORT, 10)
	fmt.Println("Running reverse proxy on " + port)
	http.ListenAndServe(port, nil)
}
