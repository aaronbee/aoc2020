package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)

	m := &machine{}
	for s.Scan() {
		m.inst = append(m.inst, parse(s.Text()))
	}

	fmt.Println(part2(m))
}

func part2(m *machine) int {
	for i, inst := range m.inst {
		switch inst.typ {
		case "nop":
			if inst.value > 0 {
				inst.typ = "jmp"
				if m.execute() == nil {
					fmt.Println("success on swapping", i)
					return m.accumulator
				}
				inst.typ = "nop"
			}
		case "jmp":
			if inst.value < 0 {
				inst.typ = "nop"
				if m.execute() == nil {
					fmt.Println("success on swapping", i)
					return m.accumulator
				}
				inst.typ = "jmp"
			}
		}
	}
	panic("what")
}

var errLoop = errors.New("loop")

func parse(s string) *instruction {
	fs := strings.Fields(s)
	v, err := strconv.Atoi(fs[1])
	if err != nil {
		panic(err)
	}
	return &instruction{typ: fs[0], value: v}
}

type instruction struct {
	typ   string
	value int
}

type machine struct {
	accumulator int
	inst        []*instruction
}

func (m *machine) execute() error {
	var pc int
	m.accumulator = 0
	seen := make(map[int]struct{})
	for {
		if pc >= len(m.inst) {
			return nil
		}
		if _, ok := seen[pc]; ok {
			return errLoop
		}
		seen[pc] = struct{}{}
		i := m.inst[pc]
		switch i.typ {
		case "acc":
			m.accumulator += i.value
			pc++
		case "jmp":
			pc += i.value
		case "nop":
			pc++
		}
	}
}
