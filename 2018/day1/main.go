package main

import "strings"
import "strconv"

func main(){
  answer := 0
  str := strings.Split(input, "\n")
  for _, e := range str {
    n, _ := strconv.Atoi(e)
    answer += n
  }
  println("The answer is:", answer)
}
