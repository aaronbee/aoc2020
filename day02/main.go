package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var valid int
	s := bufio.NewScanner(f)
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if l == "" {
			continue
		}
		if evalPassword2(l) {
			valid++
			fmt.Println("valid:", l)
		} else {
			fmt.Println("invalid:", l)
		}
	}
	f.Close()
	fmt.Println(valid)
}

func evalPassword1(l string) bool {
	fs := strings.Fields(l)
	rang, chr, password := fs[0], fs[1], fs[2]
	chr = strings.TrimSuffix(chr, ":")
	minMax := strings.Split(rang, "-")
	min, err := strconv.Atoi(minMax[0])
	if err != nil {
		panic(err)
	}
	max, err := strconv.Atoi(minMax[1])
	if err != nil {
		panic(err)
	}

	c := strings.Count(password, chr)
	return min <= c && c <= max
}

func evalPassword2(l string) bool {
	fs := strings.Fields(l)
	rang, chrS, password := fs[0], fs[1], fs[2]
	chr := chrS[0]
	minMax := strings.Split(rang, "-")
	min, err := strconv.Atoi(minMax[0])
	if err != nil {
		panic(err)
	}
	max, err := strconv.Atoi(minMax[1])
	if err != nil {
		panic(err)
	}

	fmt.Println("min", min, "max", max, "chr", string(chr), "pass", password)
	var first, second bool
	if password[min-1] == chr {
		first = true
	}
	if password[max-1] == chr {
		second = true
	}
	return (first || second) && !(first && second)
}
