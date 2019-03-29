package main

import (
	"fmt"
	"os"
	"net/http"
	"github.com/gorilla/mux";
)

func getPort() string{
	var PORT string;
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT ="8080"
	}
	return PORT;
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world me"))
}

func main() {
	PORT := getPort();
	r := mux.NewRouter();
	r.HandleFunc("/", handleIndex)
	http.ListenAndServe("0.0.0.0:" + PORT, r)
	fmt.Printf("Listening to port %v", PORT)	
}