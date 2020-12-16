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
			c.applyMask(strings.Split(line, " = ")[1])
		case "mem[":
			var addr, value uint64
			if _, err := fmt.Sscanf(line, "mem[%d] = %d", &addr, &value); err != nil {
				panic(err)
			}
			c.set(addr, value)
		}
	}

	fmt.Println(c.sum())
}

type computer struct {
	zeroMask uint64
	oneMask  uint64
	memory   map[uint64]uint64
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
