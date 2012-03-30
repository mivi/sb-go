package main

import (
  "fmt"
  "time"
  )

func pump(ch chan int) {
  for i:=0;;i++ {
    ch <- i
  }
}

func suck(ch chan int) {
  for {
    fmt.Println(<-ch)
  }
}

func main() {
  ch := make(chan int)
  go pump(ch)
  fmt.Println(<-ch)

  go suck(ch)
  fmt.Println("Me sucking ", <-ch)
  time.Sleep(2*1e9)
}
