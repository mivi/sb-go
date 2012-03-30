package main

import (
  "testing"
	"regexp"
)

func testall(t *testing.T, tst map[string][]string, reg *regexp.Regexp) {
  for k,v:= range tst {
    cnt := int(0)
    cnt = 0
    for s:= range reg.AllMatchesStringIter(k, 0) {
//    fmt.Println(s)
      if(cnt == len(v)) {
        t.Error(k + " parse returned more elems than in v")
        break
      }
      if(s != v[cnt]) {
        t.Error(s+ " doesn't matches " + v[cnt])
      }
      cnt++
    }
  }
}

func testall2d(t *testing.T, tst map[string][][]string, reg *regexp.Regexp, indx int) {
  m := make(map[string][]string)
  for k,va:= range tst {
    m[k] = va[indx]
  }
  testall(t, m, reg)
}

func TestMyCSVDocParse(t *testing.T) {
  tst := map[string][]string{
            "2009-01-01 ,,\" dfdf sds \", \"Vishal Jain\"":
                []string{"2009-01-01 ,,\" dfdf sds \", \"Vishal Jain\""},

            "2009-01-01 ,,\" dfdf sds \", \"Vishal Jain\" 10, \n \t\" dfsdf&*$&^\" , 20":
                []string{ "2009-01-01 ,,\" dfdf sds \", \"Vishal Jain\" 10, \n",
                              " \t\" dfsdf&*$&^\" , 20"},

            "2009-01-01 ,,\" dfdf sds \", \"Vishal Jain\" 10, \r\n \t\" dfsdf&*$&^\" , 20":
                []string{ "2009-01-01 ,,\" dfdf sds \", \"Vishal Jain\" 10, \r\n",
                              " \t\" dfsdf&*$&^\" , 20"},
          }
  // (([^"\r\n])*("([^"]|"")*")*)*(\r\n|\n|$)   // by NFA
  x,_ := regexp.Compile("(([^\"\r\n])*(\"([^\"]|\"\")*\")*)*(\r\n|\n|$)")
  testall(t, tst, x)
}

var lineparsetst = map[string][][]string{
                "\t,  ,2009-01-01 ,\"  \t\n\r \",  ":
                [][]string{
                  []string{                                // unfiltered
                          "\t",
                          ",",
                          " ",
                          " ",
                          ",",
                          "2009-01-01 ",
                          ",",
                          "\"  \t\n\r \",",
                          " ",
                          " "},
                  []string{                                // filtered
                          "",
                          "",
                          "2009-01-01",
                          "  \t\n\r ",
                          ""},
                          },

                "2009-01-01 ,,\" dfdf sds \", \"Vishal Jain\"":
                [][]string{
                  []string{                                // unfiltered
                          "2009-01-01 ",
                          ",",
                          ",",
                          "\" dfdf sds \",",
                          " ",
                          "\"Vishal Jain\""},
                  []string{                                // filtered
                          "2009-01-01",
                          "",
                          "dfdf sds",
                          "Vishal Jain"},
                          },

              "2009-01-01,vishaljain":
                [][]string{
                  []string{
                          "2009-01-01,",
                          "vishaljain"},
                  []string{
                          "2009-01-01",
                          "vishaljain"},
                          },

              "2009-01-01,\"Vishal Jain\"":
                [][]string{
                  []string{
                          "2009-01-01,",
                          "\"Vishal Jain\""},
                  []string{
                          "2009-01-01",
                          "Vishal Jain"},
                          },

              "\"2009-01-01\",\"Vishal Jain\",10":
                [][]string{
                  []string{
                          "\"2009-01-01\",",
                          "\"Vishal Jain\",",
                          "10"},
                  []string{
                          "2009-01-01",
                          "Vishal Jain",
                          "10"},
                          },

          }

func TestMyCSVLineParse(t *testing.T) {
  // ([^\s",]*|("([^"]|"")*"))(?=[,\s]|$)    // by NFA
  // ([^\s",]*|("([^"]|"")*"))([,\s]|$)    // used below (replace \s with all
  // white space characters) (look ahead not supported)
  x,_ := regexp.Compile("([^\a\b\f\n\r\t\v \",]*|(\"([^\"]|\"\")*\"))([,\a\b\f\n\r\t\v ]|$)")
  //x,e := regexp.Compile(`([^\s",]*|("([^"]|"")*"))(?=[,\s])`)
  //x,e := regexp.Compile(`([^\t\n",]*|("([^"]|"")*"))(?=[,\t\n])`)

  testall2d(t, lineparsetst, x, 0)
}

func TestbennadelCSVLineParse(t *testing.T) {
  // http://bennadel.com/blog/976-Regular-Expressions-Make-CSV-Parsing-In-ColdFusion-So-Much-Easier-And-Faster-.htm

  // ("(?:[^"]|"")*"|[^",\r\n]*)(,|\r\n?|\n)?
  x,_ := regexp.Compile("(\"(?:[^\"]|\"\")*\"|[^\",\r\n]*)(,|\r\n?|\n)?");
//  x,e := regexp.Compile(`("(?:[^"]|"")*"|[^",\r\n]*)(,|\r\n?|\n)?`);
  testall2d(t, lineparsetst, x, 0)
}

func TestStringSlice(t *testing.T) {
  s := "Vishal Jain"
  if(s[0:2] != "Vi") {
    t.Errorf("String " + s + " slice failed")
  }

  f := func(str string) (string){
    return str[2:6]
  }

  str := f(s[4:])
  if(str != " Jai") {
    t.Errorf("String '" + str + "' slice failed")
  }

//  f1 := func(str *string) (string){
//    return (*str)[2:6]
//  }
//
//  str = f1(s[4:])
//  if(str != " Jai") {
//    t.Errorf("String '" + str + "' slice failed")
//  }

  fchgslice := func(str string) {
    str[2]='I'
  }

  fchgslice(s[4:])
  if(s != "VishalIJain") {
    t.Errorf("String slice not passed by reference")
  }

}
