package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var data [][]byte

	s := bufio.NewScanner(f)
	for s.Scan() {
		if len(s.Bytes()) == 0 {
			continue
		}
		row := make([]byte, len(s.Bytes()))
		copy(row, s.Bytes())
		data = append(data, row)
	}
	f.Close()

	result := traverse(data, 1, 1)
	result *= traverse(data, 3, 1)
	result *= traverse(data, 5, 1)
	result *= traverse(data, 7, 1)
	result *= traverse(data, 1, 2)
	fmt.Println(result)
}

func tree(data [][]byte, x, y int) bool {
	row := data[y]
	x = x % len(row)
	return row[x] == '#'
}

func traverse(data [][]byte, dx, dy int) int {
	var count int
	var x, y int
	for y < len(data) {
		if tree(data, x, y) {
			count++
		}
		x += dx
		y += dy
	}

	return count
}
