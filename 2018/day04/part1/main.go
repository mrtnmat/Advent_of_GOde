package main

import (
	"fmt"
	"sort"
	"strings"
)

const ()

type (
	guard_id uint
	day      uint8
	min      uint8

	sleepstogram map[min]min

	guardlist map[guard_id]*guard
	guard     struct {
		id          guard_id
		sleep_total min
		hs          sleepstogram
		naps
	}

	naps []nap
	nap  struct {
		day
		start min
		end   min
	}

	input_text [][]string

	status struct {
		guardlist
		current   *guard
		sleepiest *guard
	}
)

func main() {
	fmt.Print("hello\n")

	var (
		state       status
		lines       []string
		split_lines input_text
	)

	state.guardlist = make(map[guard_id]*guard)
	lines = strings.Split(input, "\n")
	for i, _ := range lines {
		split_lines = append(split_lines, strings.Split(lines[i], " "))
	}
	//sort lines
	split_lines.sort()

	for _, words := range split_lines {
		fmt.Printf("%v\n", words)
		switch {
		case words[2] == "Guard":
			var id guard_id
			//parse id
			fmt.Sscanf(words[3], "#%d", &id)
			//if first time you see that guard, add a guard
			if _, ok := state.guardlist[id]; !ok {
				state.guardlist.add_guard(id)
			}
			state.current = state.guardlist[id]
		case words[2] == "falls":
			//parse nap start time
			var s min
			fmt.Sscanf(words[1], "00:%d", &s)
			//put guard to sleep
			state.current.sleep(s)
		case words[2] == "wakes":
			//parse nap end time
			var s min
			fmt.Sscanf(words[1], "00:%d", &s)
			//wake current guard
			state.current.wake(s)
		}
		/*
		 */
	}
	for _, e := range state.guardlist {
		e.compute_histogram()
		e.compute_sleeptotal()
		//fmt.Printf("%v\n", e)
	}
	state.sleepiest = state.guardlist.sleepiest()
	fmt.Printf("%v is the sleepiest guard\n", state.sleepiest.id)
	fmt.Printf("he slept the most during minute %v\n", state.sleepiest.hs.highest())
	fmt.Printf("the answer is: %v\n", uint(state.sleepiest.id)*uint(state.sleepiest.hs.highest()))
	/*
	 */
}

func (hs *sleepstogram) highest() min {
	var h min
	for i, e := range *hs {
		if e > (*hs)[h] {
			h = i
		}
	}
	return h
}

func (g *guard) compute_sleeptotal() {
	for _, e := range g.hs {
		g.sleep_total += e
	}
}

func (g *guard) compute_histogram() {
	g.hs = sleepstogram(make(map[min]min))
	for _, n := range g.naps {
		for i := n.start; i < n.end; i++ {
			g.hs[i]++
		}
	}
}

func (g *guard) wake(end min) {
	g.naps[len(g.naps)-1].end = end
}

func (g *guard) sleep(start min) {
	g.naps = append(g.naps, nap{start: start})
}

func (gl *guardlist) sleepiest() *guard {
	var pg *guard
	for _, e := range *gl {
		if pg == nil {
			pg = e
		} else if e.sleep_total > pg.sleep_total {
			pg = e
		}
	}
	return pg
}

func (gl *guardlist) add_guard(id guard_id) {
	(*gl)[id] = new(guard)
	(*gl)[id].id = id
}

func (inp *input_text) sort() {
	sort.Slice(*inp, func(i, j int) bool {
		return (*inp)[i][0] < (*inp)[j][0] || (*inp)[i][1] < (*inp)[j][1] && (*inp)[i][0] == (*inp)[j][0]
	})
}
