package main

import (
  "strings"
)

func main() {
/*
  input =
`abcdef
bababc
abbcde
abcccd
aabcdd
abcdee
ababab`
*/
  inp := strings.Split(input, "\n")
  for i, e := range inp {
    b := []byte(e)
    for j := i + 1; j < len(inp); j++ {
      d := count_diff(b, []byte(inp[j]))
      //println(d, i, j)
      if d == 1{
        println("The answer is:", find_common(b, []byte(inp[j])))
      }
    }
  }
}

func count_diff(s1, s2 []byte) int {
  diff := 0
  for i, e := range s1 {
    if e != s2[i] {
      diff++
    }
  }
  return diff
}

func find_common(s1, s2 []byte) string {
  common := make([]byte, 0, 32)
  for i, e := range s1 {
    if e == s2[i] {
      common = append(common, e)
    }
  }
  return string(common)
}
