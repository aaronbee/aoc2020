package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	c := computer{memory: make(map[uint64]uint64)}

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		switch line[:4] {
		case "mask":
			c.applyMask2(strings.Split(line, " = ")[1])
		case "mem[":
			var addr, value uint64
			if _, err := fmt.Sscanf(line, "mem[%d] = %d", &addr, &value); err != nil {
				panic(err)
			}
			c.set2(addr, value)
		}
	}

	fmt.Println(c.sum())
}

type computer struct {
	zeroMask uint64
	oneMask  uint64
	floats   []uint8
	memory   map[uint64]uint64
}

func (c *computer) applyMask2(s string) {
	if len(s) != 36 {
		panic(fmt.Errorf("wrong mask length: %d %q", len(s), s))
	}
	c.oneMask = 0
	c.floats = nil
	for i, r := range s {
		switch r {
		case 'X':
			c.floats = append(c.floats, uint8(35-i))
		case '1':
			c.oneMask ^= 1 << (35 - i)
		}
	}
}

func (c *computer) applyMask(s string) {
	if len(s) != 36 {
		panic(fmt.Errorf("wrong mask length: %d %q", len(s), s))
	}
	c.zeroMask = 0xF_FFFF_FFFF // 36 1's
	c.oneMask = 0
	for i, r := range s {
		switch r {
		case '0':
			c.zeroMask ^= 1 << (35 - i)
		case '1':
			c.oneMask ^= 1 << (35 - i)
		}
	}
}

func flipBits(in <-chan uint64, out chan<- uint64, index uint8) {
	for addr := range in {
		out <- addr
		out <- addr ^ (1 << index)
	}
	close(out)
}

func (c *computer) genAddrs(addr uint64) <-chan uint64 {
	ch := make(chan uint64, 1)
	ch <- addr
	close(ch)
	for _, index := range c.floats {
		out := make(chan uint64)
		go flipBits(ch, out, index)
		ch = out
	}
	return ch
}

func (c *computer) set2(addr, value uint64) {
	addr |= c.oneMask
	for addr := range c.genAddrs(addr) {
		c.memory[addr] = value
	}
}

func (c *computer) set(addr, value uint64) {
	value &= c.zeroMask
	value |= c.oneMask
	c.memory[addr] = value
}

func (c *computer) sum() (total uint64) {
	for _, v := range c.memory {
		total += v
	}
	return
}
