package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"strings"

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
	fmt.Println("Which port would you like to set the proxy on?")
	reader := bufio.NewReader(os.Stdin)
	port, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	port = strings.Replace(port, "\n", "", 1)
	if port[0] != ':' {
		port = ":" + port
	}
	http.ListenAndServe(port, nil)
}
