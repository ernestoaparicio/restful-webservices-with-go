package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

type UUID struct {
}

func giveRandomUUID(w http.ResponseWriter, r *http.Request) {
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, fmt.Sprintf("%x", b))
}

func main() {
	newMux := http.NewServeMux()
	newMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		giveRandomUUID(w, r)
	})

	newMux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Float64())
	})

	newMux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Intn(100))
	})
	http.ListenAndServe(":8000", newMux)
}
