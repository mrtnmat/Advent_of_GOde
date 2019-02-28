package main

import (
	"fmt"
)

type (
	guard struct {
    id uint
	}
	guardslist map[uint]*guard
	answer     struct {
		sleepiest *guard
	}
)

func main() {
	var (
		ans    answer
		guards guardslist
	)
  guards = make(guardslist)
  guards[0] = &guard{10}
	ans.sleepiest = guards.sleepiest_guard()
	fmt.Printf("The answer is %v\n", ans.sleepiest.id*ans.sleepiest.sleepiest_min())
}

func (gl *guardslist) sleepiest_guard() *guard {
  var g *guard
	for _, e := range *gl {
    if g == nil {
      g = e
    } else if e.sleep_time() < g.sleep_time() {
      g = e
    }
	}
	return g
}

func (g *guard) sleep_time() uint {
	return 4
}

func (g *guard) sleepiest_min() uint {
  return 2
}
