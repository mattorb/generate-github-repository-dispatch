package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os/user"
	"path/filepath"
	"strconv"
)

func postGithubRepositoryDispatch(org string, repo string, token string) {
	client := &http.Client{}
	jsonData := map[string]string{"event_type": "checklinks"}
	jsonValue, _ := json.Marshal(jsonData)
	url := "https://api.github.com/repos/" + org + "/" + repo + "/dispatches"
	ghreq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	ghreq.Header.Add("Accept", "application/vnd.github.everest-preview+json")
	ghreq.Header.Add("Authorization", "token "+token)
	ghresp, err := client.Do(ghreq)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(ghresp.Body)
		fmt.Println("POST " + url)
		fmt.Println("HTTP " + strconv.Itoa(ghresp.StatusCode) + " " + string(data)) // so far, HTTP 204: Success, but no content.
	}
}

func generateRepositoryDispatch(w http.ResponseWriter, req *http.Request) {
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))

	usr, _ := user.Current()
	dir := usr.HomeDir
	tokenPath := filepath.Join(dir, ".ghtoken")

	postGithubRepositoryDispatch("mattorb", "blog", stringFromFile(tokenPath))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/generate_repository_dispatch", generateRepositoryDispatch)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}
