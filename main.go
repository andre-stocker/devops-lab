package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from AKS! Env=%s\n", os.Getenv("ENV"))
  })

  log.Println("Listening on :8080")
  log.Fatal(http.ListenAndServe(":8080", nil))
}
