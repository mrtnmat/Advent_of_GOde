package main

import (
	"fmt"
)

type (
	guard struct {
    id uint
    histogram [60]uint
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
  guards.new_guard(12)
  guards.new_guard(10)
	ans.sleepiest = guards.sleepiest_guard()
	fmt.Printf("The answer is %v\n", ans.sleepiest.id*ans.sleepiest.sleepiest_min())
}

func (gl *guardslist) new_guard(id uint) {
  g := new(guard)
  g.id = id
  (*gl)[id] = g
}

func (gl *guardslist) sleepiest_guard() *guard {
  var g *guard
	for _, e := range *gl {
    if g == nil {
      g = e
    } else if e.sleep_time() > g.sleep_time() {
      g = e
    }
	}
	return g
}

func (g *guard) sleep_time() uint {
  var t uint
  for _, e := range g.histogram {
    t += e
  }
  return t
}

func (g *guard) sleepiest_min() uint {
  return 2
}
