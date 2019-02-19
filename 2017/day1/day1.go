package main

import "os"
import "fmt"
import "unicode"
import "strconv"

func main() {
	fd, _ := os.Open("./input")
	var nAnswer uint
	var sBuf = make([]byte, 1, 1)
	var sValues = make([]uint8, 0, 512)

	//read as long as it is a number
	for fd.Read(sBuf); unicode.IsNumber(rune(sBuf[0])); fd.Read(sBuf) {
		sValues = append(sValues, sBuf[0])
	}

	//convert all char numbers to uints numbers
	for i, e := range sValues {
		v, _ := strconv.Atoi(string(e))
		sValues[i] = uint8(v)
	}

	//compute answer
	for i, e := range sValues {
		if e == sValues[(i+1)%len(sValues)] {
			nAnswer += uint(e)
		}
	}

	//print answer
	fmt.Printf("There are %v values\n", len(sValues))
	fmt.Printf("The answer is: %v\n", nAnswer)
}
