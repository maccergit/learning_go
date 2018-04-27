package main

import (
  "net/http"
  "log"
  "encoding/json"
)

type guessHandler struct {
}

type Response struct {
  Message string `json:"message"`
}

func (g *guessHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  resp := &Response{Message: "Hi there"}
  if err := json.NewEncoder(w).Encode(resp); err != nil {
    log.Panicf("whoops! %s", err)
  }
}

func main() {
  log.Println("server started")

  http.Handle("/guesses", &guessHandler{})
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal(err)
  }
}
