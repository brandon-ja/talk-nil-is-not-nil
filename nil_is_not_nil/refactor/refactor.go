package main

import (
"bytes"
"log"
"net/http"
)

func main() {
	// setup stuff to keep the slides clean
	var body *bytes.Buffer
	var err error
	var req *http.Request
	url := "http://test.url"

	//START OMIT
	req, err = http.NewRequest(http.MethodGet, url, body)
	//END OMIT

	if err != nil {
		log.Fatalf("There was an error %s", err)
	}
	log.Printf("Req Body: %s", req.Body)
}
