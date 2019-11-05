package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func g2g(w http.ResponseWriter, req *http.Request) {
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/generate_repository_dispatch", g2g)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}
