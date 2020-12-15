package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var joltages []int
	s := bufio.NewScanner(f)
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			panic(err)
		}
		joltages = append(joltages, i)
	}

	sort.Ints(joltages)

	part2(joltages)
}

func part1(joltages []int) {
	var jumpBy1 int
	var jumpBy3 int
	prev := 0
	for _, j := range joltages {
		switch j - prev {
		case 1:
			jumpBy1++
		case 2:
		case 3:
			jumpBy3++
		default:
			panic(fmt.Errorf("invalid jump %d to %d", prev, j))
		}
		prev = j
	}
	jumpBy3++

	fmt.Println(jumpBy1 * jumpBy3)
}

func part2(joltages []int) {
	fmt.Println(recurse(joltages, 0, make(map[int]int)))
}

func recurse(joltages []int, cur int, cache map[int]int) int {
	fmt.Println("cur", cur)
	if v, ok := cache[cur]; ok {
		return v
	}
	if len(joltages) == 0 {
		return 1
	}

	var count int
	for i, v := range joltages {
		if v > cur+3 {
			break
		}
		count += recurse(joltages[i+1:], joltages[i], cache)
	}
	cache[cur] = count
	return count
}
