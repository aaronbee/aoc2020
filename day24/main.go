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
		m[c] = !m[c]
	}
	var count int
	for _, v := range m {
		if v {
			count++
		}
	}
	fmt.Println(count)
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
