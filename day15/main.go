package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := os.Args[1]
	seen := make(map[int]int)
	turn := 1
	var next int
	for _, s := range strings.Split(input, ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		next = seen[i]
		seen[i] = turn
		fmt.Println(turn, ":", i)
		turn++
	}
	for ; turn <= 2020; turn++ {
		var last int
		if next > 0 {
			last = turn - 1 - next
		}
		fmt.Println(turn, ":", last)
		next = seen[last]
		seen[last] = turn
	}
}
