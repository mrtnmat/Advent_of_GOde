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

func (pGr *grid) computeTile(x, y int) {
	value := pGr[x-1][y+1] + pGr[x][y+1] + pGr[x+1][y+1] + pGr[x-1][y] + pGr[x][y] + pGr[x+1][y] + pGr[x-1][y-1] + pGr[x][y-1] + pGr[x+1][y-1]
	pGr[x][y] = value
}

func (pGr *grid) populate() {
	x, y := 0, 0        //coordinates
	d, m, s := 1, 0, 1  //direction, movement size, square
	off := gridSize / 2 //offset
	pGr[x+off][y+off] = 1
	lastValue := 0

	for lastValue < nInput {
		for 2*x*d <= m {
			pGr.computeTile(x+off, y+off)
			lastValue = pGr[x+off][y+off]
			if lastValue > nInput {
				fmt.Printf("x:%v, y:%v, value: %v\n", x, y, lastValue)
				fmt.Printf("dist:%v\n", manhDist(x, y))
			}
			x += d
			s++
		}
		for 2*y*d <= m {
			pGr.computeTile(x+off, y+off)
			lastValue = pGr[x+off][y+off]
			if lastValue > nInput {
				fmt.Printf("x:%v, y:%v, value: %v\n", x, y, lastValue)
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
