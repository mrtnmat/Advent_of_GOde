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
  var numAnswer int
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
    fmt.Printf("%v\n", err)
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
	  //sort answer
    sort.IntSlice(numSlice).Sort()
    fmt.Printf("%v\n", numSlice)
    numAnswer += numSlice[len(numSlice) - 1] - numSlice[0]
    fmt.Printf("%v - %v\n", numSlice[len(numSlice) - 1], numSlice[0])
	}

	//print answer
  fmt.Printf("%v\n", numAnswer)
}
