package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	flag.Parse()
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	s.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF {
			advance := len(data)
			data = bytes.TrimSpace(data)
			if len(data) > 0 {
				return advance, data, nil
			}
			return advance, nil, nil
		}

		i := bytes.Index(data, []byte("\n\n"))
		if i == -1 {
			return 0, nil, nil
		}
		return i + 2, data[:i], nil
	})

	var count int
	for s.Scan() {
		g := convertGroup2(s.Bytes())
		count += bits.OnesCount32(g)
	}

	fmt.Println(count)
}

func convertGroup2(group []byte) uint32 {
	lines := bytes.Split(group, []byte("\n"))
	union := uint32(0xff_ff_ff_ff)
	for _, line := range lines {
		var val uint32
		for _, c := range line {
			if c == '\n' {
				continue
			}
			offset := c - 'a'
			val |= 1 << offset
		}
		union &= val
	}
	return union
}

func convertGroup1(group []byte) uint32 {
	var val uint32
	for _, c := range group {
		if c == '\n' {
			continue
		}
		offset := c - 'a'
		val |= 1 << offset
	}
	return val
}
