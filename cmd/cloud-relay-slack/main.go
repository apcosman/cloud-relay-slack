package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"io/ioutil"
	//"encoding/json"
)

var hc http.Client
var api_call string = "https://api-http.littlebitscloud.cc/devices/243c200ccecb/output?access_token=97612f7c1ce1b4bd1e2c31707658ec1145a765abad9f1c4af67e6fb6267931c0"

func main() {
    http.HandleFunc("/", hello)

    hc = http.Client{}

    fmt.Println("listening...")
    err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    if err != nil {
      panic(err)
    }
}

func hello(res http.ResponseWriter, req *http.Request) {
    req.ParseForm()
    fmt.Printf("%s\n", req.Form)

    req_req, err := http.NewRequest("POST", api_call, nil)
	if err != nil {
		log.Fatal(err)
	}

	client_req, err := hc.Do(req_req)
	defer client_req.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	req_res_body, err := ioutil.ReadAll(client_req.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", req_res_body)
}
