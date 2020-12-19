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

	var yourTicket []int
	s.Scan()
	if s.Text() != "your ticket:" {
		panic("expected your ticket")
	}
	s.Scan()
	for _, v := range strings.Split(s.Text(), ",") {
		v, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		yourTicket = append(yourTicket, v)
	}

	s.Scan() // empty line
	if s.Text() != "" {
		panic("expected empty")
	}
	s.Scan() // "nearby tickets:"
	if s.Text() != "nearby tickets:" {
		panic("expected nearby tickets")
	}

	candidates := make([][]*field, len(yourTicket))
	for i, v := range yourTicket {
		candidates[i] = root.fields(v)
	}
	for _, cans := range candidates {
		for i, can := range cans {
			if i != 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%q", can.label)
		}
		fmt.Println()
	}
	for s.Scan() {
		fieldsByIndex := make([][]*field, 0, len(candidates))
		for _, v := range strings.Split(s.Text(), ",") {
			v, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			fields := root.fields(v)
			if len(fields) == 0 {
				fieldsByIndex = nil
				break
			}
			fieldsByIndex = append(fieldsByIndex, fields)
		}
		if fieldsByIndex == nil {
			continue
		}
		for i, cans := range candidates {
			for j, can := range cans {
				if can == nil {
					continue
				}
				var found bool
				for _, f := range fieldsByIndex[i] {
					if f == can {
						found = true
						break
					}
				}
				if !found {
					cans[j] = nil
				}
			}
		}
	}
	finalFields := make([]string, len(candidates))
	for i, cans := range candidates {
		for _, can := range cans {
			if can != nil {
				if finalFields[i] != "" {
					fmt.Printf("Multiple fields found for %d: %s %s\n", i, finalFields[i], can.label)
				}
				finalFields[i] = can.label
			}
		}
	}
	fmt.Println(finalFields)
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

func (n *node) fields(i int) []*field {
	if n == nil {
		return nil
	}
	if i > n.max {
		return nil
	}
	if i < n.beg {
		return n.left.fields(i)
	}
	fields := n.right.fields(i)
	if i <= n.end {
		return append(fields, n.f)
	}
	return fields
}
