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
    id      int
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
	var fb fabric
	var lines []string

	for i, _ := range fb {
		for j, _ := range fb[i] {
			fb[i][j] = '.'
		}
	}

  st.fb = fb

	lines = strings.Split(input, "\n")
	for i, _ := range lines {
    st.patches = append(st.patches, patch{})
		words := strings.Split(lines[i], " ")
		fmt.Sscanf(words[0], "#%d", &st.patches[i].id)
		fmt.Sscanf(words[2], "%d,%d:", &st.patches[i].coor.x, &st.patches[i].coor.y)
		fmt.Sscanf(words[3], "%dx%d", &st.patches[i].size.x, &st.patches[i].size.y)
		st.claim_patch(st.patches[i])
	}

	//st.fb.print_all()

	fmt.Printf("total overlap: %v\n", st.overlap)
	fmt.Printf("the ID of the only free patch is: %v\n", st.find_free_patch())
}

func (s *state) find_free_patch() int {
  for i, _ := range s.patches {
    if s.patches[i].overlap == false {
      return s.patches[i].id
    }
  }
  return -1
}

func (st *state) claim_patch(ptc patch) {
	var ySlice []rune
	var xSlice [][fbsize]rune

	xSlice = st.fb[ptc.coor.x : ptc.coor.x+ptc.size.x]
	for i, _ := range xSlice {
		ySlice = xSlice[i][ptc.coor.y : ptc.coor.y+ptc.size.y]
		for j, _ := range ySlice {
			switch ySlice[j] {
			case '.':
				ySlice[j] = '/'
			case '/':
				ySlice[j] = '#'
				st.overlap++
        ptc.overlap = true
			}
		}
	}
}

func (fb *fabric) print_all() {
	fb.print(coor{0, 0}, size{fbsize, fbsize})
}

func (fb *fabric) print(c coor, s size) {
	var ySlice []rune
	var xSlice [][fbsize]rune

	xSlice = fb[c.x : c.x+s.x]
	for i, _ := range xSlice {
		ySlice = xSlice[i][c.y : c.y+s.y]
		fmt.Printf("%s\n", string(ySlice))
	}
}
