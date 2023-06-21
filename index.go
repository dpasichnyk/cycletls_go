package main

import (
	"flag"
	"log"
	nhttp "net/http"
	"os"

	"github.com/dpasichnyk/cycletls_go/app"
)

func setupRoutes() {
	nhttp.HandleFunc("/", app.WSEndpoint)
}

func main() {
	port, exists := os.LookupEnv("WS_PORT")
	var addr *string
	if exists {
		addr = flag.String("addr", ":"+port, "http service address")
	} else {
		addr = flag.String("addr", ":9112", "http service address")
	}

	setupRoutes()
	log.Fatal(nhttp.ListenAndServe(*addr, nil))
}
