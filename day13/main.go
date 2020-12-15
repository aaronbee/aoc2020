package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	s := bytes.Split(b, []byte{'\n'})
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
