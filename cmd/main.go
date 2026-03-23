package main

import (
	"log"
	"net/http"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func LoggingMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rec := &statusRecorder{
			ResponseWriter: w,
			status: http.StatusOK,
		}
		
		h.ServeHTTP(rec, r)
		
		log.Printf("%s %s %d", r.Method, r.URL.Path, rec.status)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.NotFound(w, r)
		return
	}
	
	w.Write([]byte("Hello, World!"))
}

func main() {
	http.HandleFunc("/", LoggingMiddleware(HelloHandler))

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}