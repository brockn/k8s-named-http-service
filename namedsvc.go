package main

import (
	"fmt"
	"net/http"
  "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "ServiceName: %s, Path: %s\n", os.Getenv("SERVICE_NAME"), r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
  http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	http.ListenAndServe(":8080", nil)
}
