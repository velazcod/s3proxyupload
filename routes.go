package main

import (
  "github.com/zenazn/goji"
  "flag"
)

func main() {
  goji.Get("/status", status)
  goji.Post("/upload", uploadMedia)
  flag.Set("bind", ":" + MEDIA_API_PORT)
  goji.Serve()
}
