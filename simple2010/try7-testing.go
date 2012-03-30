package main

import (
  "fmt"
//  "flag"
  "testing"
  "time"
)

func TestFunc(t *testing.T) {
}

func TestFunc1(t *testing.T) {
  t.Log("hello")
  re, err := testing.CompileRegexp("mybench")
  if err != "" {
    t.Error("regexp compile failed")
    
  }

  if !re.MatchString("mybench") {
    t.Error("mybench regexp doesn't match mybench string")
  }
}

func TestFunc2(t *testing.T) {
  t.Error("Error occured")
}

func TestFunc3(t *testing.T) {
  t.Fatal("Fatal error occured")
}

func BenchmarkFunc (b *testing.B) {
  fmt.Println("Inside benchmark")
  time.Sleep(1e9)

}

// can't redeclare flag. Its already declared in testing package
//var matchBenchmarks = flag.String("benchmarks", "", "regular expression to select benchmarks to run")

//func main() {
//
////  if len(*matchBenchmarks) == 0 {
////    fmt.Println("Please specify -benchmarks='mybench' flag")
////  }
//
//  flag.Parse()
//
////  fmt.Println(flag.Lookup("v").Value)
//  if flag.Lookup("v").Value.String() == "false" {
//    fmt.Println("Please specify -v=true for verbose print")
//  }
//
//  if len(flag.Lookup("benchmarks").Value.String()) == 0 {
//    fmt.Println("Please specify -benchmarks='<regexp>' to run benchmarks. See benchmarks.go")
//  } else {
//    fmt.Println("You specified -" + flag.Lookup("benchmarks").Name + "=" + flag.Lookup("benchmarks").Value.String())
//    fmt.Println("Running Benchmarks")
//    bench := testing.Benchmark{"mybench", BenchmarkFunc}
//
//    // These 2 lines wont work because benchmark field is not public
//  //  b := &testing.B{benchmark: bench}
//  //  b.run()
//
//    testing.RunBenchmarks(&[...]testing.Benchmark{bench}) //Wont work because
//  //  of check for -benchmark flag on command line. So HAVE TO SUPPLY -benchmarks
//  //  on cmd line
//  }
//
//  fmt.Println("Running Testcases")
//  tst := testing.Test {"mytest", TestFunc}
//  tst1 := testing.Test {"mytest1", TestFunc1}
//  tst2 := testing.Test {"mytest2", TestFunc2}
//  tst3 := testing.Test {"mytest3", TestFunc3}
//  tstarr := [...]testing.Test{tst, tst1, tst2, tst3}  // FAIL
////  tstarr := [...]testing.Test{tst, tst1}  // PASS
//  testing.Main(&tstarr)
//
//}
