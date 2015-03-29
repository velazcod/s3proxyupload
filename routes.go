package main

import (
  "github.com/zenazn/goji"
  "flag"
  "fmt"
  "os"
)

const AppVersion = "1.0"

func main() {

  version := flag.Bool("v", false, "prints current version")
  flag.Parse()
  if *version {
    fmt.Println(AppVersion)
    os.Exit(0)
  }

  if MEDIA_API_PORT == "" {
    fmt.Println("The MEDIA_API_PORT variable is required")
    os.Exit(0)
  }

  goji.Get("/status", status)
  goji.Post("/upload", uploadMedia)
  flag.Set("bind", ":" + MEDIA_API_PORT)
  goji.Serve()
}
