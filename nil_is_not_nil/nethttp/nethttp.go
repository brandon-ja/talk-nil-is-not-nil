package main

import (
	"bytes"
	"io"
	"net/http"
)

func main() {
	var bodyBuf *bytes.Buffer
	body := io.Reader(bodyBuf)
	var req *http.Request

	//START OMIT
	if body != nil {
		switch v := body.(type) {
		case *bytes.Buffer:
			req.ContentLength = int64(v.Len()) // HL
			buf := v.Bytes()
			req.GetBody = func() (io.ReadCloser, error) {
				r := bytes.NewReader(buf)
				return io.NopCloser(r), nil
			}
			//END OMIT
		}
	}
}

//START_SIG OMIT
func NewRequest(method, url string, body io.Reader) (*Request, error) {
	//END_SIG OMIT
	return nil, nil
}
