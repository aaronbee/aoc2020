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
	for _, c := range []byte(input) {
		cur.n = &cup{v: int(c - '0')}
		cur = cur.n
	}
	head = head.n
	cur.n = head

	game(head)
	for head.v != 1 {
		head = head.n
	}
	for c := head.n; c != head; c = c.n {
		fmt.Print(c.v)
	}
	fmt.Println()
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

func game(cur *cup) {
	for i := 0; i < 100; i++ {
		dest := 1 + ((cur.v - 2 + 9) % 9)
		three := pickup3(cur)
		for {
			var found bool
			for c := three; c != nil; c = c.n {
				if c.v == dest {
					found = true
					dest = 1 + ((dest - 2 + 9) % 9)
				}
			}
			if !found {
				break
			}
		}
		for c := cur; ; c = c.n {
			if c.v == dest {
				place3(c, three)
				break
			}
		}
		cur = cur.n
	}
}
