package main

import "strings"
import "strconv"


func main(){
  answer := 0
  str := strings.Split(input, "\n")
  num := make([]int, 0, 1024)
  m := make(map[int]int, 1024)
  for _, e := range str {
    n, _ := strconv.Atoi(e)
    num = append(num, n)
  }
OuterLoop:
  for {
    for _, e := range num {
      m[answer]++
      if m[answer] > 1 {
        break OuterLoop
      }
      answer += e
    }
  }
  println("The answer is:", answer)
}
