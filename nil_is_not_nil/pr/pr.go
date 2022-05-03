package main

import (
  "bytes"
  "log"
  "net/http"
)

func main() {
  //START_BODY OMIT
  var body *bytes.Buffer
  //END_BODY OMIT

  // setup stuff to keep the slides clean
  var err error
  var req *http.Request
  url := "http://test.url"

  //START_CHECK OMIT
  if body != nil {
    req, err = http.NewRequest(http.MethodGet, url, body)
  } else {
    req, err = http.NewRequest(http.MethodGet, url, nil)
  }
  //END_CHECK OMIT

  if err != nil {
    log.Fatalf("There was an error %s", err)
  }
  log.Printf("Req Body: %+v", req.Body)
}
