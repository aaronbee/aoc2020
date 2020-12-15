package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	var seatIDs []uint16
	s := bufio.NewScanner(f)
	for s.Scan() {
		id := toSeatID(s.Bytes())
		seatIDs = append(seatIDs, id)
	}

	sort.Slice(seatIDs, func(i, j int) bool {
		return seatIDs[i] < seatIDs[j]
	})

	curr := seatIDs[0]
	for _, id := range seatIDs[1:] {
		if id > curr+1 {
			fmt.Println(id)
		}
		curr = id
	}
}

func toSeatID(b []byte) uint16 {
	var val uint16
	for _, c := range b {
		switch c {
		case 'F', 'L':
		case 'B', 'R':
			val++
		}
		val <<= 1
	}
	return val >> 1
}
