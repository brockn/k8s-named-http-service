package main

import (
	"fmt"
	"net/http"
  "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
  if os.Getenv("SERVICE_NAME") == "" {
	  fmt.Fprintf(w, "ServiceName: Null")
  } else {
	  fmt.Fprintf(w, "ServiceName: %s", os.Getenv("SERVICE_NAME"))
  }
}

func main() {
	http.HandleFunc("/", handler)
  http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	http.ListenAndServe(":8080", nil)
}
