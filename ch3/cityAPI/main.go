package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restful-webservices-with-go/ch3/multipleMiddleware"

	"github.com/justinas/alice"
)

type city struct {
	Name string
	Area uint64
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var tempCity city
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&tempCity)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		fmt.Printf("Got %s city area of %d sq miles!\n", tempCity.Name, tempCity.Area)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("201 - Created"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method Not Allowed"))
	}
}

//func main() {
//	http.HandleFunc("/city", postHandler)
//	http.ListenAndServe(":8000", nil)
//}

//func main() {
//	originalHandler := http.HandlerFunc(postHandler)
//	http.Handle("/city", multipleMiddleware.FilterContentType(multipleMiddleware.SetServerTimeCookie(originalHandler)))
//	http.ListenAndServe(":8000", nil)
//}

// with alice
func main() {
	originalHandler := http.HandlerFunc(postHandler)
	chain := alice.New(multipleMiddleware.FilterContentType, multipleMiddleware.SetServerTimeCookie).Then(originalHandler)
	http.Handle("/city", chain)
	http.ListenAndServe(":8000", nil)
}
