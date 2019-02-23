package main

import "fmt"

const inp = 277678

func main() {
	x, y := 0, 0
	m, s, d := 1, 1, 1
OuterLoop:
	for s <= inp {
		pC := &x
		for y*2*d < m {
			if x*2*d >= m {
				pC = &y
			}
			*pC += d
			s++
			if s == inp {
				fmt.Printf("%v %v %v\n", s, x, y)
				break OuterLoop
			}
		}
		d *= -1
		m++
	}
}
