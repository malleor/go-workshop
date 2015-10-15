package main

import (
  "fmt"
)

func main() {
  fmt.Println("Hello! This is a vending machine!")

  items := []string{"chocolate", "lolipop"}
  price := []int{12, 7}

  coin_ch := make(chan int)
  item_ch := make(chan string)
  
  choose := func(choice int) {
    go func(i string, p int) {
      fmt.Println("OK. Waiting for", p, "in cash...")

      total := 0
      for {
        coin := <- coin_ch
        total += coin
        fmt.Println("Got", coin)
        if total >= p {
          fmt.Println(total, "is enough. Returning a", i)
          item_ch <- i
          return
        }
      }
    }(items[choice], price[choice])
  }

  insert := func(coin int) {
    coin_ch <- coin
  }

  wait := func() {
    item := <- item_ch
    fmt.Println("Nice! Got", item)
  }

  choose(1)
  insert(5)  
  insert(2)
  wait()

}

