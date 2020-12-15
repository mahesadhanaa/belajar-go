package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/appconfig"
)

func hello(w http.ResponseWriter, req *http.Request) {
	config := appconfig.GetEnvironment(params)
	fmt.Fprintf(w, "hello\n")
}

func main() {

	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":6969", nil)
}
