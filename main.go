package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//creates an http server listening on localhost:8080
	http.HandleFunc("/", myFunc)
	http.ListenAndServe(":8080", nil)
}

//Sends a greeting with the server's IP included
func myFunc(w http.ResponseWriter, r *http.Request) {
	myIP := getMyIP()
	_, err := fmt.Fprintln(w, "Hello from", myIP)
	if err != nil {
		log.Fatalln(err)
	}
}

//Gets the public IP of the server
func getMyIP() string {
	resp, err := http.Get("http://api.ipify.org")
	if err != nil {
		log.Fatalln(err)
	}
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(ip)
}
