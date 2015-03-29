package main

import (
  "github.com/zenazn/goji"
  "flag"
)

func main() {
  if MEDIA_API_PORT == "" {
    fmt.Println("The MEDIA_API_PORT variable is required")
    os.Exit(0)
  }

  goji.Get("/status", status)
  goji.Post("/upload", uploadMedia)
  flag.Set("bind", ":" + MEDIA_API_PORT)
  goji.Serve()
}
