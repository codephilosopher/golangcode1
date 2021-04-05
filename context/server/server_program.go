package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:3000", nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	select {
	case <-time.After(3 * time.Second):
		fmt.Fprintln(w, "hello")
	case <-ctx.Done():
		log.Fatalln(ctx.Err().Error())
		http.Error(w, ctx.Err().Error(), http.StatusInternalServerError)
	}
}
