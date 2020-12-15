package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const size = 25

type circleSumBuffer struct {
	m   map[int]int
	buf []int
}

var errInvalid = errors.New("invalid")

func (b *circleSumBuffer) append(i int) error {
	if len(b.m) != size {
		b.m[i] = len(b.buf)
		b.buf = append(b.buf, i)
		return nil
	}

	var valid bool
	for j, k := range b.m {
		if l, ok := b.m[i-j]; ok && k != l {
			valid = true
			break
		}
	}
	if !valid {
		return errInvalid
	}
	old := b.buf[len(b.buf)-25]
	delete(b.m, old)
	b.m[i] = len(b.buf)
	b.buf = append(b.buf, i)
	if len(b.m) != 25 {
		panic(fmt.Errorf("unexpected size of map: %d", len(b.m)))
	}
	return nil
}

func (b *circleSumBuffer) findSum(i int) int {
	lo, hi := 0, 2
	sum := b.buf[0] + b.buf[1]
	for sum != i {
		if sum < i {
			sum += b.buf[hi]
			hi++
		} else {
			sum -= b.buf[lo]
			lo++
		}
	}
	min, max := b.buf[lo], b.buf[lo]
	for _, j := range b.buf[lo:hi] {
		if j < min {
			min = j
		}
		if j > max {
			max = j
		}
	}
	return min + max
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b := &circleSumBuffer{m: map[int]int{}}
	s := bufio.NewScanner(f)
	var bad int
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			panic(err)
		}
		if b.append(i) != nil {
			fmt.Println("Part 1:", i)
			bad = i
			break
		}
	}
	fmt.Println("Part 2:", b.findSum(bad))
}
