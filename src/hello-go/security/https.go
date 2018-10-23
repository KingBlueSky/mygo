package main

import (
	"fmt"
	"net/http"
)

const (
	ServerPort       = 8088
	ServerDomain     = "localhost"
	ResponseTemplate = "hello"
)

func rootHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Length", fmt.Sprint(len(ResponseTemplate)))
	w.Write([]byte(ResponseTemplate))
}

func main() {
	http.HandleFunc("/", rootHandler)
	//http.ListenAndServe(":8888", nil)
	http.ListenAndServeTLS(fmt.Sprintf(":%d", ServerPort), "rui.crt", "rui.key", nil)
}
