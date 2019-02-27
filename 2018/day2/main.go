package main

import (
  "strings"
)

func main() {
  twos := 0
  threes := 0
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
  for _, e := range inp {
    b := []byte(e)
    mp := make(map[byte]int, 32)
    for _, e := range b {
      mp[e]++
    }
    for _, e := range mp {
      if e == 2 {
        twos++
        break
      }
    }
    for _, e := range mp {
      if e == 3 {
        threes++
        break
      }
    }
  }
  println("twos:", twos, "threes:", threes)
  println("answer:", twos * threes)
}
