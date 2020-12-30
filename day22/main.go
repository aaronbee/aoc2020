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

	for !d1.lose() && !d2.lose() {
		c1 := d1.play()
		c2 := d2.play()
		if c1 < c2 {
			d2.append(c2)
			d2.append(c1)
		} else {
			d1.append(c1)
			d1.append(c2)
		}
	}
	fmt.Println(d1.score(), d2.score())
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

func game(d1, d2 *deck) (int, int) {
	for !d1.lose() && !d2.lose() {
		c1 := d1.play()
		c2 := d2.play()
		if c1 < c2 {
			d2.append(c2)
			d2.append(c1)
		} else {
			d1.append(c1)
			d1.append(c2)
		}
	}
	return d1.score(), d2.score()
}
