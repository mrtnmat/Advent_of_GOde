package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

const (
	slp = time.Duration(1 << 10)
)

type (
	polymer []byte
)

func main() {
	var p polymer
	p, _ = ioutil.ReadFile("./input.txt")
	p = p[:len(p)-1]
	p = p.react()
	fmt.Printf("%v monomers left after reacting completely.\n", len(p))
}

func (pc polymer) react() polymer {
	w := 0
	for i, _ := range pc {
		pc[w] = pc[i]
		fmt.Printf("%c", pc[w])
		time.Sleep(slp)
		if w > 1 && is_reactive(pc[w], pc[w-1]) {
			fmt.Print("\010\040\010")
			time.Sleep(slp)
			w--
			fmt.Print("\010\040\010")
			time.Sleep(slp)
			w--
		}
		w++
	}
	fmt.Print("\n")
	return pc[:w]
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
