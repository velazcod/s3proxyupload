package main

import (
  "github.com/zenazn/goji"
)

func main() {
  goji.Get("/status", status)
  goji.Post("/upload", uploadMedia)
  goji.Serve()
}
