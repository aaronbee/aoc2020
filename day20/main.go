package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"strings"
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

	var corner *tile
	for _, t := range tiles {
		neighs := m.sidesWithNeighbors(t)
		if len(neighs) != 2 {
			continue
		}
		corner = t
		if neighs[0] == 0 && neighs[1] == 1 {
			t.rotation = 1
		} else if neighs[0] == 1 && neighs[1] == 2 {
		} else if neighs[0] == 2 && neighs[1] == 3 {
			t.rotation = 3
		} else if neighs[0] == 0 && neighs[1] == 3 {
			t.rotation = 2
		} else {
			panic("huh")
		}
		break
	}
	sideLength := int(math.Sqrt(float64(len(tiles))))
	arranged := make([][]*tile, sideLength)
	for i := range arranged {
		arranged[i] = make([]*tile, sideLength)
		row := arranged[i]
		if i == 0 {
			row[0] = corner
		} else {
			above := arranged[i-1][0]
			row[0] = m.getWithSide(above.sides()[2], 0)
		}
		m.remove(row[0])
		row[0].applytransformations()
		expectedNeighbors := 2
		if i == sideLength-1 {
			expectedNeighbors--
		}
		for j := range row[1:] {
			row[j+1] = m.getWithSide(row[j].sides()[1], 3)
			m.remove(row[j+1])
			row[j+1].applytransformations()
		}
	}

	var img [][]byte
	for _, row := range arranged {
		n := len(row[0].contents)
		for i := 1; i < n-1; i++ {
			img = append(img, nil)
			for _, t := range row {
				part := t.contents[i][1 : n-1]
				img[len(img)-1] = append(img[len(img)-1], part...)
			}
		}
	}

	//                   #
	// #    ##    ##    ###
	//  #  #  #  #  #  #
	pattern := [][]int{
		{18},
		{0, 5, 6, 11, 12, 17, 18, 19},
		{1, 4, 7, 10, 13, 16},
	}
	var found bool

	for r := 0; r < 8; r++ {
		for i := 0; i < len(img)-2; i++ {
			for j := 0; j < len(img[i])-19; j++ {
				if match(img[i][j:], pattern[0]) &&
					match(img[i+1][j:], pattern[1]) &&
					match(img[i+2][j:], pattern[2]) {
					found = true
					mark(img[i][j:], pattern[0])
					mark(img[i+1][j:], pattern[1])
					mark(img[i+2][j:], pattern[2])
				}
			}
		}
		if found {
			break
		}
		img = rotate(img, 1)
		if r == 3 {
			for i := range img {
				flip(img[i])
			}
		}
	}

	for _, row := range img {
		fmt.Println(string(row))
	}
	var count int
	for _, row := range img {
		count += bytes.Count(row, []byte{'#'})
	}
	fmt.Println(count)
}

func match(row []byte, pattern []int) bool {
	for _, i := range pattern {
		if row[i] != '#' {
			return false
		}
	}
	return true
}

func mark(row []byte, pattern []int) {
	for _, i := range pattern {
		row[i] = 'O'
	}
}

type tilesBySide map[string][]*tile

func (m tilesBySide) insert(t *tile) {
	for _, sd := range t.sides() {
		s := string(sd)
		m[s] = append(m[s], t)
	}
}

func (m tilesBySide) remove(t *tile) {
	cp := *t
	cp.flip = false
	cp.rotation = 0
	for _, sd := range cp.sides() {
		slc := m[string(sd)]
		orig := len(slc)
		for i, tt := range slc {
			if t == tt {
				slc[i] = slc[len(slc)-1]
				slc = slc[:len(slc)-1]
				break
			}
		}
		if len(slc) == orig {
			panic("remove didn't remove anything")
		}
		if len(slc) == 0 {
			delete(m, string(sd))
		} else {
			m[string(sd)] = slc
		}
	}
}

func (m tilesBySide) getWithSide(sd []byte, side int) *tile {
	cand1 := m[string(sd)]
	flip(sd)
	cand2 := m[string(sd)]
	if len(cand1)+len(cand2) == 0 {
		panic("tile not found")
	}
	var t *tile
	if len(cand1) == 1 {
		t = cand1[0]
		t.flip = true
	} else if len(cand2) == 1 {
		t = cand2[0]
	} else {
		fmt.Printf("candidates for %s\n", string(sd))
		fmt.Println("orig")
		for _, t := range cand1 {
			fmt.Println(t)
		}
		fmt.Println("flipped")
		for _, t := range cand1 {
			fmt.Println(t)
		}
		panic("too many tiles")
	}

	for i, sid := range t.sides() {
		if bytes.Equal(sd, sid) {
			t.rotation = side - i + 4
			if !bytes.Equal(t.sides()[side], sd) {
				panic(fmt.Errorf("rotation didn't work"))
			}
			return t
		}
	}
	fmt.Println("looking for", string(sd))
	for _, sd := range t.sides() {
		fmt.Println(string(sd))
	}
	panic("didn't find matching side")
}

func (m tilesBySide) sidesWithNeighbors(t *tile) []int {
	var result []int
outer:
	for i, sd := range t.sides() {
		for _, neighbor := range m[string(sd)] {
			if neighbor != t {
				result = append(result, i)
				continue outer
			}
		}
		flip(sd)
		for _, neighbor := range m[string(sd)] {
			if neighbor != t {
				result = append(result, i)
				break
			}
		}
	}
	return result
}

type tile struct {
	id       int
	contents [][]byte
	rotation int
	flip     bool
}

func (t *tile) String() string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "Tile %d:\n", t.id)
	for _, row := range t.contents {
		buf.Write(row)
		buf.WriteByte('\n')
	}
	return buf.String()
}

func (t *tile) sides() [][]byte {
	result := make([][]byte, 4)
	result[0] = append([]byte(nil), t.contents[0]...)
	result[2] = append([]byte(nil), t.contents[len(t.contents)-1]...)
	flip(result[2])

	result[1] = make([]byte, len(t.contents[0]))
	result[3] = make([]byte, len(t.contents[0]))

	for i, row := range t.contents {
		result[3][len(t.contents)-1-i] = row[0]
		result[1][i] = row[len(row)-1]
	}
	if t.flip {
		flip(result[0])
		flip(result[1])
		flip(result[2])
		flip(result[3])
		result[1], result[3] = result[3], result[1]
	}
	switch t.rotation % 4 {
	case 0:
	case 1:
		result[0], result[1], result[2], result[3] = result[3], result[0], result[1], result[2]
	case 2:
		result[0], result[1], result[2], result[3] = result[2], result[3], result[0], result[1]
	case 3:
		result[0], result[1], result[2], result[3] = result[1], result[2], result[3], result[0]
	default:
		panic(fmt.Errorf("invalid rotation %d", t.rotation))
	}

	return result
}

func (t *tile) applytransformations() {
	if t.flip {
		for i := range t.contents {
			flip(t.contents[i])
		}
		t.flip = false
	}
	t.contents = rotate(t.contents, t.rotation%4)

	t.rotation = 0
}

func flip(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func rotate(a [][]byte, c int) [][]byte {
	switch c {
	case 0:
		return a
	case 1:
		n := len(a)
		m := make([][]byte, n)
		for i := range a {
			m[i] = make([]byte, n)
			for j := range a {
				if i == 84 || j == 84 || n-j-1 == 84 {

				}
				m[i][j] = a[n-j-1][i]
			}
		}
		return m
	case 2:
		for i := range a {
			flip(a[i])
		}
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
		return a
	case 3:
		n := len(a)
		m := make([][]byte, n)
		for i := range a {
			m[i] = make([]byte, n)
			for j := range a {
				m[i][j] = a[j][n-i-1]
			}
		}
		return m
	}
	panic("unreachable")
}
