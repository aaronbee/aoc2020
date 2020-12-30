package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var d1 deck
	s.Scan()
	for s.Scan() {
		line := s.Text()
		if line == "" {
			break
		}
		c, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		d1.append(c)
	}
	var d2 deck
	s.Scan()
	for s.Scan() {
		line := s.Text()
		c, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		d2.append(c)
	}

	fmt.Println(game(&d1, &d2))
}

type deck struct {
	cards []int
}

func (d *deck) play() int {
	c := d.cards[0]
	d.cards = d.cards[1:]
	return c
}

func (d *deck) append(x int) {
	d.cards = append(d.cards, x)
}

func (d *deck) lose() bool {
	return len(d.cards) == 0
}

func (d *deck) score() int {
	var s int
	for i, c := range d.cards {
		s += c * (len(d.cards) - i)
	}
	return s
}

func (d *deck) count() int {
	return len(d.cards)
}

func (d *deck) copy(i int) *deck {
	return &deck{cards: append([]int(nil), d.cards[:i]...)}

}

func (d *deck) signature() string {
	byt := make([]byte, len(d.cards))
	for i, c := range d.cards {
		byt[i] = byte(c)
	}
	return string(byt)
}

func game(d1, d2 *deck) (int, int) {
	cache := make(map[string]struct{})
	for !d1.lose() && !d2.lose() {
		sig := d1.signature() + " " + d2.signature()
		if _, ok := cache[sig]; ok {
			return 1, 0
		}
		cache[sig] = struct{}{}
		c1 := d1.play()
		c2 := d2.play()
		var s1, s2 int
		if c1 <= d1.count() && c2 <= d2.count() {
			s1, s2 = game(d1.copy(c1), d2.copy(c2))
		} else {
			s1, s2 = c1, c2
		}
		if s1 < s2 {
			d2.append(c2)
			d2.append(c1)
		} else {
			d1.append(c1)
			d1.append(c2)
		}
	}
	return d1.score(), d2.score()
}
