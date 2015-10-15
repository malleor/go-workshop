package main

import (
  "fmt"
)

func main() {
  fmt.Println("Hello there!")

  ch := make(chan string)

  go func() {  
    fmt.Println("goroutine!")
    ch <- "done"
  }()

  <- ch
}

