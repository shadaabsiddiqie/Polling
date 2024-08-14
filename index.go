package main

import (
	"log"
	"net/http"
)

func main() {
	// go LongPolling()
	go ShotPolling()
	htmlHandler := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "view/index.html")
	}
	setDataHandler := func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write(DataToWrite)
	}
	http.HandleFunc("/setData", setDataHandler)
	http.HandleFunc("/index", htmlHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
