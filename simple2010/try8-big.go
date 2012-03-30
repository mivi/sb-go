package main

import (
  "fmt"
  "testing"
  "big"
)

func TestNewInt(t *testing.T) {
  x := big.NewInt(int64(int64(1<<62)+646213))
  x.Mul(x, big.NewInt(int64(1<<37)))
  fmt.Print("Bitlen ")
  fmt.Println(x.BitLen())

  fmt.Print("Binomial ")
  fmt.Print(x.Binomial(int64(5), int64(2)))
  fmt.Print(" ")
  fmt.Println(x)
}

func TestQuo(t *testing.T) {
  fmt.Println("Testing Quo")
  x := big.NewInt(int64(0))
  z := big.NewInt(int64(1<<62))
  z.Mul(z, big.NewInt(int64((1<<30)+342342143)))
  fmt.Print("z ")
  fmt.Println(z)

  r := big.NewInt(int64(0))
  x.QuoRem(z, big.NewInt(int64((1<<23)+2342344257)), r)
  fmt.Print("x ")
  fmt.Println(x)
  fmt.Print("r ")
  fmt.Println(r)
}

func TestQuoSimple(t *testing.T) {
  fmt.Println("Testing QuoSimple")
  x := big.NewInt(int64(0))
  z := big.NewInt(int64(513))
  fmt.Print("z ")
  fmt.Println(z)

  r := big.NewInt(int64(0))
  x.QuoRem(z, big.NewInt(int64(25)), r)
  fmt.Print("x ")
  fmt.Println(x)
  fmt.Print("r ")
  fmt.Println(r)
}

func TestDiv(t *testing.T) {
  fmt.Println("Testing Div")
  x := big.NewInt(int64(0))
  z := big.NewInt(int64(1<<62))
  z.Mul(z, big.NewInt(int64((1<<30)+342342143)))
  fmt.Print("z ")
  fmt.Println(z)

  r := big.NewInt(int64(0))
  x.DivMod(z, big.NewInt(int64((1<<23)+2342344257)), r)
  fmt.Print("x ")
  fmt.Println(x)
  fmt.Print("r ")
  fmt.Println(r)
}

func TestDivSimple(t *testing.T) {
  fmt.Println("Testing DivSimple")
  x := big.NewInt(int64(0))
  z := big.NewInt(int64(513))
  fmt.Print("z ")
  fmt.Println(z)

  r := big.NewInt(int64(0))
  x.DivMod(z, big.NewInt(int64(25)), r)
  fmt.Print("x ")
  fmt.Println(x)
  fmt.Print("r ")
  fmt.Println(r)
}

func TestRat(t *testing.T) {
  fmt.Println("Testing Rat")
  x := big.NewRat(34234212327, 5464523)
  fmt.Print("x ")
  fmt.Print(x)
  fmt.Print(" ")
  fmt.Println(x.FloatString(5))
  if x.FloatString(5) != "6264.81256" {
    t.Error("Float String rep not matching")
  }
}

func TestRatString(t *testing.T) {
  fmt.Println("Testing RatString")
  x := big.NewRat(0,1)
  _, ok := x.SetString("5.23")
  fmt.Print("x ")
  fmt.Print(x)
  fmt.Print(" ")
  fmt.Println(x.FloatString(4))
  fmt.Print("ok ")
  fmt.Println(ok)
}

func TestRatSub(t *testing.T) {
  fmt.Println("Testing RatSub")
  x := big.NewRat(0,1).Sub(big.NewRat(1,3), big.NewRat(1,2))
  fmt.Print("x ")
  fmt.Println(x)
  if (x.Num().Cmp(big.NewInt(-1)) != 0) || (x.Denom().Cmp(big.NewInt(6)) != 0) {
    t.Error("Num and Denom don't match")
  }
}
