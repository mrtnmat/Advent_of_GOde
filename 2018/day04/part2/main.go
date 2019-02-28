package main

import (
	"fmt"
	"sort"
	"strings"
)

type (
	guardslist map[uint]*guard
	guard      struct {
		id        uint
		histogram [60]uint
	}

	nap struct {
		id    uint
		start uint
		end   uint
	}
	answer struct {
		sleepiest *guard
	}
)

func main() {
	var (
		ans    answer
		guards guardslist
		naps   []nap
		rec    []string
	)
	rec = strings.Split(input, "\n")
	sort.Strings(rec)
	naps = parse_rec(rec)
	guards = make(guardslist)
	for i, _ := range naps {
		naps[i].add_to_histogram(guards)
	}
	ans.sleepiest = guards.sleepiest_guard()
	fmt.Printf("The sleepiest guard is %v\n", ans.sleepiest.id)
	fmt.Printf("%v\n", ans.sleepiest.histogram)
	m, _ := ans.sleepiest.sleepiest_min()
	fmt.Printf("He slept the most during minute %v\n", m)
	fmt.Printf("The answer is %v\n", ans.sleepiest.id*m)

	g, m := guards.most_consistent_guard()
	fmt.Printf("The most consistent sleeper is guard #%v\n", g.id)
	fmt.Printf("You can often find him asleep at 00:%v\n", m)
	fmt.Printf("The answer to the second part is %v\n", g.id*m)
}

func (gl guardslist) new_guard(id uint) {
	g := new(guard)
	g.id = id
	gl[id] = g
}

func (gl guardslist) most_consistent_guard() (*guard, uint) {
	var g *guard
	var m uint
	for _, e := range gl {
		em, ev := e.sleepiest_min()
		if g != nil {
			_, gv := g.sleepiest_min()
			if ev > gv {
				g = e
				m = em
			}
		} else {
			g = e
			m = em
		}
	}
	return g, m
}

func (gl guardslist) sleepiest_guard() *guard {
	var g *guard
	for _, e := range gl {
		if g == nil || e.sleep_time() > g.sleep_time() {
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

func (g *guard) sleepiest_min() (min, val uint) {
	for i, e := range g.histogram {
		if e > g.histogram[min] {
			min = uint(i)
			val = uint(e)
		}
	}
	return
}

func (n *nap) add_to_histogram(gl guardslist) {
	var g *guard
	if gl[n.id] == nil {
		gl.new_guard(n.id)
	}
	g = gl[n.id]

	for i := n.start; i < n.end; i++ {
		g.histogram[i]++
	}
}

func parse_rec(rec []string) []nap {
	var naps []nap
	var buf []string
	var pn *nap
	var cur uint

	naps = make([]nap, 0, len(rec)/2)

	for _, e := range rec {
		buf = strings.Split(e, " ")
		switch buf[2] {
		case "Guard":
			fmt.Sscanf(buf[3], "#%d", &cur)
		case "falls":
			pn = new(nap)
			pn.id = cur
			fmt.Sscanf(buf[1], "00:%d", &pn.start)
		case "wakes":
			fmt.Sscanf(buf[1], "00:%d", &pn.end)
			naps = append(naps, *pn)
		}
	}
	return naps
}
