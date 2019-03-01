package main

import (
	"fmt"
	"io/ioutil"
)

type (
	polymer []byte
)

func main() {
	var p polymer
	p, _ = ioutil.ReadFile("./input.txt")
	p = p[:len(p)-1]
	for ch := 'a' - 1; ch < 'z'; ch++ {
		l := p.react_ignoring(byte(ch))
		fmt.Printf("len %v\n", l)
	}
}

func (pc polymer) react_ignoring(ch byte) int {
	wrt := make(polymer, len(pc))
	copy(wrt, pc)

	fmt.Printf("ignoring %c/%c: ", ch, ch-('a'-'A'))

	w := 0
	for i, _ := range pc {
		wrt[w] = pc[i]
		if w > 1 && is_reactive(wrt[w], wrt[w-1]) {
			w -= 2
		} else if wrt[w] == ch || is_reactive(wrt[w], ch) {
			w -= 1
		}
		w++
	}
	return len(wrt[:w])
}

func is_reactive(m1, m2 byte) bool {
	d := int(m1) - int(m2)
	if d < 0 {
		d = -d
	}
	if d == 'a'-'A' {
		return true
	} else {
		return false
	}
}
