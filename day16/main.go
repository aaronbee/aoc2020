package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	var root *node
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		if line == "" {
			break
		}
		var f field
		split := strings.Split(line, ":")
		f.label = split[0]
		_, err := fmt.Sscanf(split[1], "%d-%d or %d-%d", &f.a1, &f.a2, &f.b1, &f.b2)
		if err != nil {
			panic(err.Error() + ": " + line)
		}
		if root == nil {
			root = &node{f: &f, beg: f.a1, end: f.a2, max: f.a2}
		} else {
			root.insert(&node{f: &f, beg: f.a1, end: f.a2, max: f.a2})
		}
		root.insert(&node{f: &f, beg: f.b1, end: f.b2, max: f.b2})
	}

	// ignore your ticket
	for s.Scan() {
		if s.Text() == "nearby tickets:" {
			break
		}
	}

	var errCount int
	for s.Scan() {
		for _, v := range strings.Split(s.Text(), ",") {
			i, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			if !root.present(i) {
				errCount += i
			}
		}
	}
	fmt.Println(errCount)
}

type field struct {
	label  string
	a1, a2 int
	b1, b2 int
}

// https://en.wikipedia.org/wiki/Interval_tree#Augmented_tree
type node struct {
	f           *field
	beg, end    int
	max         int
	left, right *node
}

func (n *node) insert(nn *node) int {
	if nn.beg < n.beg {
		if n.left == nil {
			n.left = nn
			return n.max
		}
		n.left.insert(nn)
		return n.max
	}
	var max int
	if n.right == nil {
		n.right = nn
		max = nn.max
	} else {
		max = n.right.insert(nn)
	}
	if max > n.max {
		n.max = max
	}
	return n.max
}

func (n *node) present(i int) bool {
	if n == nil {
		return false
	}
	if i > n.max {
		return false
	}
	if i < n.beg {
		return n.left.present(i)
	}
	if i <= n.end {
		return true
	}
	return n.right.present(i)
}
