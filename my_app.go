package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  for _, kvPair := range os.Environ() {
    splitUp := strings.Split(kvPair, "=")
    key, value := splitUp[0], splitUp[1]
    fmt.Fprintf(os.Stdout, "key: %s, value: %s\n", key, value)
  }
}
