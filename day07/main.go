package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type graph map[string]*bag

type bag struct {
	color string
	c     []constraint
}

type constraint struct {
	count     int
	container *bag
	contains  *bag
}

var r = regexp.MustCompile(`(\w+ \w+) bags contain (\d+)`)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)

	g := make(graph)
	for s.Scan() {
		handleLine(g, s.Text())
	}
	part2(g)
}

func part1(g graph) {
	cache := make(map[*bag]int)
	goal := g["shiny gold"]
	for _, b := range g {
		traverse(b, goal, cache)
	}
	var count int
	for _, c := range cache {
		if c > 0 {
			count++
		}
	}
	fmt.Println(count)
}

func part2(g graph) {
	start := g["shiny gold"]
	cache := make(map[*bag]int)
	traverseToLeaf(start, cache)
	fmt.Println(cache[start])
}

func handleLine(g graph, line string) {
	split := strings.Split(line, " bags contain ")
	container := split[0]
	b, ok := g[container]
	if !ok {
		b = &bag{color: container}
		g[container] = b
	}

	if split[1] == "no other bags." {
		return
	}

	contains := strings.TrimSuffix(split[1], ".")
	for _, s := range strings.Split(contains, ",") {
		fields := strings.Fields(s)
		count, err := strconv.Atoi(fields[0])
		if err != nil {
			panic(err)
		}
		other := fields[1] + " " + fields[2]
		otherB, ok := g[other]
		if !ok {
			otherB = &bag{color: other}
			g[other] = otherB
		}
		b.c = append(b.c, constraint{count: count, container: b, contains: otherB})
	}
}

func traverse(b *bag, goal *bag, cache map[*bag]int) int {
	if c, ok := cache[b]; ok {
		return c
	}
	var val int
	for _, constraint := range b.c {
		if constraint.contains == goal {
			fmt.Println("found one in", b.color)
			val += constraint.count
			continue
		}
		val += traverse(constraint.contains, goal, cache)
	}
	cache[b] = val
	return val
}

func traverseToLeaf(b *bag, cache map[*bag]int) int {
	if c, ok := cache[b]; ok {
		return c
	}
	var val int
	for _, constraint := range b.c {
		fmt.Printf("%s contains %d %s\n", b.color, constraint.count, constraint.contains.color)
		val += constraint.count
		val += constraint.count * traverseToLeaf(constraint.contains, cache)
	}
	cache[b] = val
	return val
}
