package main

import (
	"fmt"
)

type (
	guard struct {
		id        uint
		histogram [60]uint
	}
	nap struct {
		id    uint
		start uint
		end   uint
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
		naps   []nap
	)
	guards = make(guardslist)
	guards.new_guard(99)
	guards.new_guard(10)
	naps = append(naps, nap{10, 5, 25})
	naps = append(naps, nap{10, 30, 55})
	naps = append(naps, nap{99, 40, 50})
	naps = append(naps, nap{10, 24, 29})
	naps = append(naps, nap{99, 36, 46})
	naps = append(naps, nap{99, 45, 55})
	for i, _ := range naps {
		naps[i].add_to_histogram(&guards)
	}
	ans.sleepiest = guards.sleepiest_guard()
	fmt.Printf("The sleepiest guard is %v\n", ans.sleepiest.id)
	fmt.Printf("%v\n", ans.sleepiest.histogram)
	fmt.Printf("He slept the most during minute %v\n", ans.sleepiest.sleepiest_min())
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
	var max int
	for i, e := range g.histogram {
		if e > g.histogram[max] {
			max = i
		}
	}
	return uint(max)
}

func (n *nap) add_to_histogram(gl *guardslist) {
	var g *guard
	if (*gl)[n.id] == nil {
		panic("nonexistant guard")
	} else {
		g = (*gl)[n.id]
	}
	for i := n.start; i < n.end; i++ {
		g.histogram[i]++
	}
}
