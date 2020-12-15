package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	width := bytes.IndexByte(b, '\n')
	initial := bytes.ReplaceAll(b, []byte{'\n'}, nil)
	s := &seats{s: initial,
		width:  width,
		height: len(initial) / width,
	}
	t := &seats{s: make([]byte, len(s.s)), width: s.width, height: s.height}
	var modified bool
	modified = s.iterate(t)
	for modified {
		s, t = t, s
		modified = s.iterate(t)
	}
	fmt.Println(bytes.Count(t.s, []byte{'#'}))
}

type seats struct {
	s      []byte
	width  int
	height int
}

func (s *seats) String() string {
	var buf strings.Builder
	buf.Grow(len(s.s) + s.height)
	for y := 0; y < s.height; y++ {
		buf.Write(s.s[y*s.width : (y+1)*s.width])
		if y+1 != s.height {
			buf.WriteByte('\n')
		}
	}
	return buf.String()
}

func (s *seats) index1(x, y int) byte {
	if x < -1 || x > s.width || y < -1 || y > s.height {
		panic(fmt.Errorf("out of bounds %d x %d", x, y))
	}
	if x == -1 || x == s.width || y == -1 || y == s.height {
		return '.'
	}
	return s.s[y*s.width+x]
}

func (s *seats) see(x, y, dx, dy int) int {
	if x < -1 || x > s.width || y < -1 || y > s.height {
		panic(fmt.Errorf("out of bounds %d x %d", x, y))
	}
	for {
		x += dx
		y += dy
		if x == -1 || x == s.width || y == -1 || y == s.height {
			return 0
		}
		switch s.s[y*s.width+x] {
		case '#':
			return 1
		case 'L':
			return 0
		}
	}
	return 0
}

func equal(a, b *seats) bool {
	return bytes.Equal(a.s, b.s)
}

func (s *seats) iterate(n *seats) bool {
	copy(n.s, s.s)
	var modified bool
	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			c := s.index1(x, y)
			if c == '.' {
				continue
			}
			var count int
			count += s.see(x, y, -1, -1)
			count += s.see(x, y, 0, -1)
			count += s.see(x, y, 1, -1)
			count += s.see(x, y, -1, 0)
			count += s.see(x, y, 1, 0)
			count += s.see(x, y, -1, 1)
			count += s.see(x, y, 0, 1)
			count += s.see(x, y, 1, 1)
			if c == 'L' && count == 0 {
				modified = true
				n.s[y*s.width+x] = '#'
			} else if c == '#' && count >= 5 {
				modified = true
				n.s[y*s.width+x] = 'L'
			}
		}
	}
	return modified
}
