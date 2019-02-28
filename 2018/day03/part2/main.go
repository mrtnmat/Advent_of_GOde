package main

import (
	"fmt"
	"strings"
)

const (
	fbsize int = 1024
)

type (
	fabric [fbsize][fbsize]int
	coor   struct{ x, y int }
	size   struct{ x, y int }
	patch  struct {
		id int
		coor
		size
		overlap bool
	}
	state struct {
		fb      fabric
		patches []patch
		overlap int
	}
)

func main() {
	var st state
	var lines []string

	lines = strings.Split(input, "\n")
	for i, _ := range lines {
		st.patches = append(st.patches, patch{})
		words := strings.Split(lines[i], " ")
		fmt.Sscanf(words[0], "#%d", &st.patches[i].id)
		fmt.Sscanf(words[2], "%d,%d:", &st.patches[i].coor.x, &st.patches[i].coor.y)
		fmt.Sscanf(words[3], "%dx%d", &st.patches[i].size.x, &st.patches[i].size.y)
		st.apply_patch(&st.patches[i])
	}

	st.fb.print_all()

	fmt.Printf("total overlap: %v\n", st.overlap)
	fmt.Printf("the ID of the only free patch is: %v\n", st.find_free_patch())
}

func (s *state) find_free_patch() int {
	for i, _ := range s.patches {
		if s.patches[i].overlap == false {
			fmt.Printf("%v\n", s.patches[i])
			return s.patches[i].id
		}
	}
	return -1
}

func (st *state) apply_patch(ptc *patch) {
	var ySlice []int
	var xSlice [][fbsize]int

	xSlice = st.fb[ptc.coor.x : ptc.coor.x+ptc.size.x]
	for i, _ := range xSlice {
		ySlice = xSlice[i][ptc.coor.y : ptc.coor.y+ptc.size.y]
		for j, _ := range ySlice {
			switch ySlice[j] {
			case 0:
				ySlice[j] = ptc.id
			default:
				st.overlap++
				st.patches[ySlice[j]-1].overlap = true
				fallthrough
			case -1:
				ySlice[j] = -1
				ptc.overlap = true
			}
		}
	}
}

func (fb *fabric) print_all() {
	fb.print(coor{0, 0}, size{fbsize, fbsize})
}

func (fb *fabric) print(c coor, s size) {
	var ySlice []int
	var xSlice [][fbsize]int

	xSlice = fb[c.x : c.x+s.x]
	for i, _ := range xSlice {
		ySlice = xSlice[i][c.y : c.y+s.y]
		line := make([]byte, 0, fbsize)
		for _, e := range ySlice {
			switch e {
			case -1:
				line = append(line, '#')
			case 0:
				line = append(line, '.')
			default:
				line = append(line, '\\')
			}
		}
		fmt.Printf("%s\n", line)
	}
}
