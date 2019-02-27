package main

import (
	"fmt"
	"strings"
)

const (
	fbsize int = 1024
)

type (
	fabric [fbsize][fbsize]rune
	coor   struct{ x, y int }
	size   struct{ x, y int }
	patch  struct {
		coor
		size
	}
	state struct {
		fb      fabric
		overlap int
	}
)

func main() {
	var st state
	var fb fabric
	var lines []string
	var ptc patch

	for i, _ := range fb {
		for j, _ := range fb[i] {
			fb[i][j] = '.'
		}
	}

	st = state{fb, 0}

	lines = strings.Split(input, "\n")
	for i, _ := range lines {
		words := strings.Split(lines[i], " ")
		fmt.Sscanf(words[2], "%d,%d:", &ptc.coor.x, &ptc.coor.y)
		fmt.Sscanf(words[3], "%dx%d", &ptc.size.x, &ptc.size.y)
		st.claim_patch(ptc.coor, ptc.size)
	}

	st.fb.print_all()

	fmt.Print("total overlap: %v\n", st.overlap)
}

func (st *state) claim_patch(c coor, s size) {
	var ySlice []rune
	var xSlice [][1024]rune

	xSlice = st.fb[c.x : c.x+s.x]
	for i, _ := range xSlice {
		ySlice = xSlice[i][c.y : c.y+s.y]
		for j, _ := range ySlice {
			switch ySlice[j] {
			case '.':
				ySlice[j] = '/'
			case '/':
				ySlice[j] = '#'
				st.overlap++
			}
		}
	}
}

func (fb *fabric) print_all() {
	fb.print(coor{0, 0}, size{fbsize, fbsize})
}

func (fb *fabric) print(c coor, s size) {
	var ySlice []rune
	var xSlice [][1024]rune

	xSlice = fb[c.x : c.x+s.x]
	for i, _ := range xSlice {
		ySlice = xSlice[i][c.y : c.y+s.y]
		fmt.Printf("%s\n", string(ySlice))
	}
}
