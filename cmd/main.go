package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"

	"github.com/SAHIL-Sharma21/caching-proxy/cache"
	"github.com/SAHIL-Sharma21/caching-proxy/server"
)

func main() {
	fmt.Println("Building caching proxy server in go lang")

	//initializing the cache
	cache.Init()

	//CLI strtuctre to accept flags like server url and option like cache expiry
	//definign cmd line flags
	port := flag.String("port", "8080", "PORT to start the caching server")
	origin := flag.String("origin", "", "Origin server url to proxy requests")

	//parsing command line flags
	flag.Parse()

	if *origin == "" {
		slog.Info("No url provided to proxy requests")
		log.Fatal("You must provide and origin URl using --origin flag")
	}

	//starting server
	log.Printf("Starting caching server on port %s forwarding requests to %s", *port, *origin)
	server.StartServer(*port, *origin)
}
