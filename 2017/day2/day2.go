package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var ansArr [2]int
	//open the input file for reading
	fd, err := os.Open("./input")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	//make new bufio.reader for the input file
	var lineSlice = make([]byte, 1024)
	rdr := bufio.NewReader(fd)

	for {
		//read line
		lineSlice, err = rdr.ReadSlice('\n')
		if err == io.EOF {
			break
		}

		//read each number
		numReader := bufio.NewReader(strings.NewReader(string(lineSlice)))
		var numSlice = make([]int, 0)
		for {
			bufSlice, err := numReader.ReadSlice('\t')
			numBuf, _ := strconv.Atoi(string(bufSlice[:len(bufSlice)-1]))
			numSlice = append(numSlice, numBuf)
			if err == io.EOF {
				break
			}
		}

		//sort answer and get the diff between largest and smallest
		sort.IntSlice(numSlice).Sort()
		ansArr[0] += numSlice[len(numSlice)-1] - numSlice[0]

		//compute second answer
	OuterLoop:
		for i, e := range numSlice {
			for _, f := range numSlice[i+1:] {
				if (f % e) == 0 {
					//fmt.Printf("i: %v, j: %v\n", i, j)
					//fmt.Printf("e: %v, f: %v f/e: %v\n", e, f, f/e)
					ansArr[1] += f / e
					break OuterLoop
				}
			}
		}
	}

	//print answer
	fmt.Printf("The answer to the first part is: %v\n", ansArr[0])
	fmt.Printf("The answer to the second part is: %v\n", ansArr[1])
}
