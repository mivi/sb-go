package main

import (
  "os"
  "fmt"
)

func trystr(s []string) {
  os.Stdout.WriteString("hello\n")
  //os.Stdout.WriteString(*s)
  fmt.Println(s)
  fmt.Println(*s)
}
func main() {
  ss := string("vishal")
//  trystr(&ss);
  trystr(ss[1:3]);
//  os.Stdout.WriteString(ss[1:3])
//  trystr(ss[1:3])
}


