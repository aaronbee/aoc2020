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
	spc := space{a: make([][][]uint8, 1)}
	s := bufio.NewScanner(f)
	for s.Scan() {
		cp := make([]byte, len(s.Bytes()))
		for i, b := range s.Bytes() {
			switch b {
			case '.':
			case '#':
				cp[i] = 1
			default:
				panic("unexpected char")
			}

		}
		spc.a[0] = append(spc.a[0], cp)
	}
	for range make([]int, 6) {
		next := space{}
		next.init(&spc)
		next.cycle(&spc)
		spc = next
	}
	fmt.Println(spc.count())
}

type space struct {
	a [][][]uint8
}

func (spc *space) count() int {
	var count int
	for _, plane := range spc.a {
		for _, row := range plane {
			for _, cube := range row {
				count += int(cube)
			}
		}
	}
	return count
}

func (spc *space) index(x, y, z int) uint8 {
	if z < 0 || z >= len(spc.a) {
		return 0
	}
	plane := spc.a[z]
	if y < 0 || y >= len(plane) {
		return 0
	}
	row := plane[y]
	if x < 0 || x >= len(row) {
		return 0
	}
	return row[x]
}

func (spc *space) neighbors(x, y, z int) int {
	var count uint8
	for _, dz := range []int{-1, 0, 1} {
		for _, dy := range []int{-1, 0, 1} {
			for _, dx := range []int{-1, 0, 1} {
				if dz == 0 && dy == 0 && dx == 0 {
					continue
				}
				count += spc.index(x+dx, y+dy, z+dz)
			}
		}
	}
	return int(count)
}

func (spc *space) init(prev *space) {
	spc.a = make([][][]uint8, len(prev.a)+2)
	for z := range spc.a {
		spc.a[z] = make([][]uint8, len(prev.a[0])+2)
		for y := range spc.a[z] {
			spc.a[z][y] = make([]uint8, len(prev.a[0][0])+2)
		}
	}
}

func (spc *space) cycle(prev *space) {
	for z, plane := range spc.a {
		prevZ := z - 1
		for y, row := range plane {
			prevY := y - 1
			for x := range row {
				prevX := x - 1
				count := prev.neighbors(prevX, prevY, prevZ)
				if count == 3 {
					row[x] = 1
				} else if count == 2 && prev.index(prevX, prevY, prevZ) == 1 {
					row[x] = 1
				}
			}
		}
	}
}
