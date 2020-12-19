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

	candidates := make([]map[*field]struct{}, len(yourTicket))
	for i, v := range yourTicket {
		fields := root.fields(v)
		m := make(map[*field]struct{})
		for _, f := range fields {
			m[f] = struct{}{}
		}
		candidates[i] = m
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
			for can := range cans {
				var found bool
				for _, f := range fieldsByIndex[i] {
					if f == can {
						found = true
						break
					}
				}
				if !found {
					delete(cans, can)
				}
			}
		}
	}
	finalFields := make(map[int]string)
	for len(finalFields) != len(candidates) {
		for i, cans := range candidates {
			if len(cans) != 1 {
				continue
			}
			var final *field
			for final = range cans {
			}
			finalFields[i] = final.label
			for _, cans := range candidates {
				delete(cans, final)
			}
		}
	}
	fmt.Println(finalFields)

	result := 1
	for i, label := range finalFields {
		if strings.HasPrefix(label, "departure ") {
			result *= yourTicket[i]
		}
	}
	fmt.Println(result)
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

func (n *node) insert(nn *node) {
	if nn.beg < n.beg {
		if n.left == nil {
			n.left = nn
		} else {
			n.left.insert(nn)
		}
		if n.left.max > n.max {
			n.max = n.left.max
		}
		return
	}
	if n.right == nil {
		n.right = nn
	} else {
		n.right.insert(nn)
	}
	if n.right.max > n.max {
		n.max = n.right.max
	}
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
	return n.left.present(i) || n.right.present(i)
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
	fields := append(n.left.fields(i), n.right.fields(i)...)
	if i <= n.end {
		return append(fields, n.f)
	}
	return fields
}
