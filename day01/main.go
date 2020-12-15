package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part2()
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	vals := map[int]struct{}{}

	s := bufio.NewScanner(f)
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if l == "" {
			continue
		}
		v, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		if _, ok := vals[2020-v]; ok {
			fmt.Println(v * (2020 - v))
			os.Exit(0)
		}
		vals[v] = struct{}{}
	}

	panic("should not get here")
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	vals := map[int]struct{}{}

	s := bufio.NewScanner(f)
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if l == "" {
			continue
		}
		v, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		vals[v] = struct{}{}
	}
	f.Close()

	for i := range vals {
		for j := range vals {
			if _, ok := vals[2020-i-j]; ok {
				fmt.Println(i * j * (2020 - i - j))
				os.Exit(0)
			}
		}
	}
}
