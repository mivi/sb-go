package main

import (
	"fmt"
	//  "io"
	"io/ioutil"
	"strconv"
	"os"
	"testing"
	"flag"
  "time"
)

var dirtouse = flag.String("dirtoread", "/home/vjain/shorthair/trunk/R/csv", "dir to read file from.. used for benchmark")

func myreadfile(s string, c chan int) {
	//  file, err := os.Open(s, os.O_RDONLY, 0)
	//  if err != nil {
	//    fmt.Println(err)
	//    return -1
	//  }
	//  defer file.Close()
	//  var buf [1000000]byte
	//  len := int64(0)
	//  for {
	//    n, e := file.Read(buf[0:])
	//    fmt.Println(e)
	//    fmt.Println(n)
	//    len += int64(n)
	//    if e == os.EOF {
	//      fmt.Println("Total: " + s + " " + strconv.Itoa64(len))
	//      return len
	//    }
	//  }
	//  return 0

	bt, err := ioutil.ReadFile(s)
	if err != nil {
		c <- -1
		return
	}
	//fmt.Println(s + " " + strconv.Itoa(len(bt)))
	c <- len(bt)
}

var lst = []string(nil)
//var lst = (*([]*os.FileInfo))(nil)

//var lst,_ = ioutil.ReadDir(".")
func BenchmarkAsyncRead(b *testing.B) {
  if len(lst) < 1 {
    lstFileInfo, _ := ioutil.ReadDir(*dirtouse)
    lst = make([]string, len(lstFileInfo))
    for cnt, i := range lstFileInfo {
      lst[cnt] = *dirtouse + "/" + i.Name
    }
  }

	fmt.Println("Count of files in " + *dirtouse + " " + strconv.Itoa(len(lst)))
	if len(lst) < 1 {
		fmt.Println("Provide another directory in -dirtoread. Can't read from " + *dirtouse)
		os.Exit(2)
	}

  start := time.Nanoseconds()
	c := make(chan int)
	for _, i := range lst {
		go myreadfile(i, c)
	}

	for i := 0; i < len(lst); i++ {
		<-c
	}
  fmt.Print((time.Nanoseconds()-start)/1e6)
  fmt.Println("ms")
}

func BenchmarkInseqread(b *testing.B) {
  if len(lst) < 1 {
    lstFileInfo, _ := ioutil.ReadDir(*dirtouse)
    lst = make([]string, len(lstFileInfo))
    for cnt, i := range lstFileInfo {
      lst[cnt] = *dirtouse + "/" + i.Name
    }
  }

	fmt.Println("Count of files in " + *dirtouse + " " + strconv.Itoa(len(lst)))
	if len(lst) < 1 {
		fmt.Println("Provide another directory in -dirtoread. Can't read from " + *dirtouse)
		os.Exit(2)
	}

  start := time.Nanoseconds()
	c := make(chan int)
	for _, i := range lst {
		go myreadfile(i, c)
		<-c
	}
  fmt.Print((time.Nanoseconds()-start)/1e6)
  fmt.Println("ms")
}

