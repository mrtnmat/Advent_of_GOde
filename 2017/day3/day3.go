package main

import (
	"fmt"
	"math"
)

const (
	nInput   = 277678
	gridSize = 1024
)

type grid [gridSize][gridSize]int

func main() {
	var gr grid
	gr.populate()
}

func (pGr *grid) populate() {
	x, y := 0, 0        //coordinates
	d, m, s := 1, 0, 1  //direction, movement size, steps
	off := gridSize / 2 //offset

	for s <= nInput {
		for 2*x*d <= m {
			pGr[y+off][x+off] = s
			if s == nInput {
				fmt.Printf("x:%v, y:%v, s: %v\n", x, y, s)
				fmt.Printf("dist:%v\n", manhDist(x, y))
			}
			x += d
			s++
		}
		for 2*y*d <= m {
			pGr[y+off][x+off] = s
			if s == nInput {
				fmt.Printf("x:%v, y:%v, s: %v\n", x, y, s)
				fmt.Printf("dist:%v\n", manhDist(x, y))
			}
			y += d
			s++
		}
		d *= -1
		m++
	}
}

func manhDist(x, y int) int {
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}
