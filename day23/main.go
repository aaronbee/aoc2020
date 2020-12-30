package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := os.Args[1]
	head := &cup{}
	cur := head
	m := make(map[int]*cup, 1_000_000)
	for _, c := range []byte(input) {
		cur.n = &cup{v: int(c - '0')}
		cur = cur.n
		m[cur.v] = cur
	}
	for i := 10; i <= 1_000_000; i++ {
		cur.n = &cup{v: i}
		cur = cur.n
		m[cur.v] = cur
	}
	head = head.n
	cur.n = head

	game(head, m)
	one := m[1]
	fmt.Println(one.n.v * one.n.n.v)
}

type cup struct {
	v int
	n *cup
}

func pickup3(c *cup) *cup {
	next := c.n
	third := c.n.n.n
	c.n = third.n
	third.n = nil
	return next
}

func place3(c *cup, three *cup) {
	three.n.n.n = c.n
	c.n = three
}

func str(c *cup) string {
	var buf strings.Builder
	head := c
	fmt.Fprint(&buf, head.v)
	for c := head.n; c != head; c = c.n {
		fmt.Fprint(&buf, c.v)
	}
	return buf.String()
}

func decr(i int) int {
	return 1 + ((i - 2 + 1000000) % 1000000)
}

func game(cur *cup, m map[int]*cup) {
	for i := 0; i < 10_000_000; i++ {
		dest := decr(cur.v)
		three := pickup3(cur)
		for {
			var found bool
			for c := three; c != nil; c = c.n {
				if c.v == dest {
					found = true
					dest = decr(dest)
				}
			}
			if !found {
				break
			}
		}

		place3(m[dest], three)
		cur = cur.n
	}
}
