package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

func main() {
	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	part2(bytes.Split(b, []byte{'\n'}))
}

func part2(s [][]byte) {
	type pair struct {
		id     int
		offset int
	}
	var ids []pair
	for i, idB := range bytes.Split(s[1], []byte{','}) {
		if idB[0] == 'x' {
			continue
		}
		id, err := strconv.Atoi(string(idB))
		if err != nil {
			panic(err)
		}
		offset := (id - i) % id
		if offset < 0 {
			offset += id
			if offset < 0 {
				panic("huh")
			}
		}
		ids = append(ids, pair{id, offset})
	}

	sort.Slice(ids, func(i, j int) bool {
		return ids[j].id < ids[i].id
	})
	candidate := ids[0].offset
	step := ids[0].id
	for _, id := range ids[1:] {
		fmt.Println("candidate:", candidate, "step:", step)
		fmt.Printf("Looking for %d mod %d\n", id.offset, id.id)
		for {
			if candidate%id.id == id.offset {
				break
			}
			candidate += step
		}
		step *= id.id
	}
	fmt.Println(candidate)
}

func part1(s [][]byte) {
	t, err := strconv.Atoi(string(s[0]))
	if err != nil {
		panic(err)
	}
	minWait := 0xFFFFFFFF
	minID := 0
	for _, idB := range bytes.Split(s[1], []byte{','}) {
		if idB[0] == 'x' {
			continue
		}
		id, err := strconv.Atoi(string(idB))
		if err != nil {
			panic(err)
		}
		prev := (t / id) * id
		if prev == id {
			minWait = 0
			minID = id
			break
		}
		next := prev + id
		wait := next - t
		fmt.Println("id", id, "prev", prev, "next", next, "wait", wait)
		if wait < minWait {
			minWait = wait
			minID = id
		}
	}
	fmt.Println("minWait", minWait, "minID", minID)
	fmt.Println(minWait * minID)
}
