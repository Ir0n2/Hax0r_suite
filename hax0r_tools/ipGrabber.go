package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	fmt.Println("User IP address:", ip)
	//fmt.Fprint(w, "Hello, World!")
	http.ServeFile(w, r, "/home/zerocool/pagetest.html")
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
