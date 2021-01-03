package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	m := make(map[coordinate]bool)
	s := bufio.NewScanner(f)
	for s.Scan() {
		c := coordinate{}
		for i := 0; i < len(s.Bytes()); i++ {
			switch s.Bytes()[i] {
			case 'e':
				c = c.e()
			case 'w':
				c = c.w()
			case 'n':
				i++
				switch s.Bytes()[i] {
				case 'e':
					c = c.ne()
				case 'w':
					c = c.nw()
				}
			case 's':
				i++
				switch s.Bytes()[i] {
				case 'e':
					c = c.se()
				case 'w':
					c = c.sw()
				}
			}
		}
		if m[c] {
			delete(m, c)
		} else {
			m[c] = true
		}
	}
	fmt.Println(len(m))

	for i := 0; i < 100; i++ {
		m = flip(m)
	}
	fmt.Println(len(m))
}

type coordinate struct {
	x, y int
}

func (c coordinate) e() coordinate  { return coordinate{c.x + 1, c.y} }
func (c coordinate) w() coordinate  { return coordinate{c.x - 1, c.y} }
func (c coordinate) ne() coordinate { return coordinate{c.x + (c.y & 1), c.y + 1} }
func (c coordinate) nw() coordinate { return coordinate{c.x + (c.y & 1) - 1, c.y + 1} }
func (c coordinate) se() coordinate { return coordinate{c.x + (c.y & 1), c.y - 1} }
func (c coordinate) sw() coordinate { return coordinate{c.x + (c.y & 1) - 1, c.y - 1} }

func bToI(b bool) int {
	if b {
		return 1
	}
	return 0
}

func flip(m map[coordinate]bool) map[coordinate]bool {
	n := make(map[coordinate]bool)
	for c := range m {
		neighbors := bToI(m[c.e()]) + bToI(m[c.w()]) + bToI(m[c.ne()]) + bToI(m[c.nw()]) + bToI(m[c.se()]) + bToI(m[c.sw()])
		if neighbors == 1 || neighbors == 2 {
			n[c] = true
		}
		for _, c := range []coordinate{c.e(), c.w(), c.ne(), c.nw(), c.se(), c.sw()} {
			if m[c] {
				continue
			}
			neighbors := bToI(m[c.e()]) + bToI(m[c.w()]) + bToI(m[c.ne()]) + bToI(m[c.nw()]) + bToI(m[c.se()]) + bToI(m[c.sw()])
			if neighbors == 2 {
				n[c] = true
			}
		}
	}
	return n
}
