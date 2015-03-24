package main

import (
  "crypto/sha1"
  "fmt"
)

func keyOf(data []byte) string {
  sha := sha1.New()
  sha.Write(data)
  return fmt.Sprintf("%x", string(sha.Sum(nil)))
}
