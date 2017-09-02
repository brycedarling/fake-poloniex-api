package main

import (
	"flag"
	"fmt"
	"gopkg.in/jcelliott/turnpike.v2"
	"log"
	"net/http"
)

func main() {
	var realm string
	flag.StringVar(&realm, "realm", "realm1", "realm")

	var port int
	flag.IntVar(&port, "port", 8000, "port")

	flag.Parse()

	addr := fmt.Sprintf(":%d", port)

	turnpike.Debug()

	s := turnpike.NewBasicWebsocketServer(realm)

	allowAllOrigin := func(r *http.Request) bool { return true }

	s.Upgrader.CheckOrigin = allowAllOrigin

	server := &http.Server{
		Handler: s,
		Addr:    addr,
	}

	log.Println("turnpike server starting on port", port)

	log.Fatal(server.ListenAndServe())
}
