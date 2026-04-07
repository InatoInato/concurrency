package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello?"))
	})

	fmt.Println("Running in port 8080")
	
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}