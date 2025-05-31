package route

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func Setup() {
	log.Println("setup routers")

	mux := http.NewServeMux()
	AccountRouter(mux)

	addr := getAddr()
	log.Fatal(http.ListenAndServe(addr, mux))
}

func getAddr() string {
	host := getHost()
	port := getPort()
	return host + ":" + port
}

func getHost() string {
	host := os.Getenv("HOST")
	if strings.TrimSpace(host) == "" {
		host = "localhost"
	}
	return host
}

func getPort() string {
	port := os.Getenv("PORT")
	if strings.TrimSpace(port) == "" {
		port = "8080"
	}
	return port
}
