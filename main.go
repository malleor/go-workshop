package main

import (
  "fmt"
  "time"
  "sync"
  "math/rand"
)

func main() {
  fmt.Println("Hello there!")

  N := 10000

  wg := sync.WaitGroup{}
  wg.Add(N)

  rand.Seed(42)
  for ę:=0; ę<N; ę++ {
    go func(n int) {
      t := time.Millisecond*time.Duration(rand.Int()%500)
      time.Sleep(t)
      fmt.Println("goroutine", n, "slept", t)
      wg.Done()
    }(ę)
  }

  wg.Wait()
}

