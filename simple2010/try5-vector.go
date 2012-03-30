package main

import (
  "container/vector"
  "fmt"
  )

func infcheck(y interface{}) {
  fmt.Printf("%T\n", y);
  fmt.Print(y);
  x1 := y;
 // y1 := interface{}(20);
  fmt.Print(x1.(int) < 2);
//  fmt.Print(x1 < y1);
  x:=interface{}(y);
  z:=x
  fmt.Printf("%T\n", x);
  fmt.Print(x.(int) < 10);
  fmt.Printf("%T\n", z);
//  fmt.Print(x.Less(2))
}

func checkpoint(x *int) {
  fmt.Print(x)
  fmt.Print("\n")
}

func checkTypeComp () {
  fmt.Print("checkTypeComp \n")
  type MyInt int
  x := MyInt(10)
  fmt.Print(x < 20)
  fmt.Print("\n")

  //type S struct { X, Y, Z int}
  type S1 struct {X, Y, Z int}

//  y:= S{10, 20, 30}
  y1 := struct{X, Y, Z int} {1000, 2000, 3000}
  z := S1{100, 200, 300}
  z = y1
  fmt.Print(z)
  fmt.Print("\n")
  fmt.Print("End checkTypeComp \n")
}

func main() {
  re1:=new(vector.Vector)
  re1.Push(10)
  fmt.Print(re1.Pop())

  infcheck(5)

  re1.Push(10)
  re1.Push(20)
//  fmt.Print(re1.Less(0,1))
  fmt.Print("\n")
  i := 10
  checkpoint(&i);
  checkTypeComp()
//  fmt.Print(int32(10).Less(5));
}
