package server

import (
	"io"
	"log"
	"net/http"

	"github.com/SAHIL-Sharma21/caching-proxy/cache"
)

func StartServer(port, origin string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleRequest(w, r, origin)
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// function to handle requests
func handleRequest(w http.ResponseWriter, r *http.Request, origin string) {
	//checking if the requst is cached or not

	if cachedResponse, found := cache.Cache[r.URL.Path]; found {
		log.Printf("serving from the cache: %s\n", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(cachedResponse))
		return
	}

	proxyURL := origin + r.URL.Path
	// fmt.Println(r.URL.Path)
	log.Printf("fetching from origin %s", proxyURL)

	resp, err := http.Get(proxyURL)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	//cache the response
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cache.Cache[r.URL.Path] = string(body)

	//writing response to client
	for k, v := range resp.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}

	w.WriteHeader(resp.StatusCode)
	_, _ = w.Write(body)
}
