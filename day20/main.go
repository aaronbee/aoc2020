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
	defer f.Close()

	var tiles []*tile
	m := make(tilesBySide)
	scn := bufio.NewScanner(f)
	for scn.Scan() {
		var t tile
		fmt.Sscanf(scn.Text(), "Tile %d:", &t.id)
		for scn.Scan() {
			if len(scn.Bytes()) == 0 {
				break
			}
			t.contents = append(t.contents,
				append([]byte(nil), scn.Bytes()...))
		}
		tiles = append(tiles, &t)
		m.insert(&t)
	}

	result := 1
	var count int
	for _, t := range tiles {
		neighs := m.sidesWithNeighbors(t)
		if neighs == 2 {
			count++
			result *= t.id
		}
	}
	fmt.Printf("Found %d candidate tiles. Multiplied: %d\n", count, result)
}

type tilesBySide map[string][]*tile

func (m tilesBySide) insert(t *tile) {
	for _, sd := range t.sides() {
		s := string(sd)
		m[s] = append(m[s], t)
	}
}

func (m tilesBySide) sidesWithNeighbors(t *tile) int {
	var count int
outer:
	for _, sd := range t.sides() {
		for _, neighbor := range m[string(sd)] {
			if neighbor != t {
				count++
				continue outer
			}
		}
		flip(sd)
		for _, neighbor := range m[string(sd)] {
			if neighbor != t {
				count++
				break
			}
		}
	}
	return count
}

type tile struct {
	id       int
	contents [][]byte
}

func (t *tile) sides() [][]byte {
	result := make([][]byte, 4)
	result[0] = append([]byte(nil), t.contents[0]...)
	result[2] = append([]byte(nil), t.contents[len(t.contents)-1]...)

	result[1] = make([]byte, len(t.contents[0]))
	result[3] = make([]byte, len(t.contents[0]))

	for i, row := range t.contents {
		result[3][i] = row[0]
		result[1][i] = row[len(row)-1]
	}

	return result
}

func flip(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
