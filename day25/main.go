package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	pkA, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	pkB, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

	lsA := determineLoopSize(pkA)
	lsB := determineLoopSize(pkB)

	fmt.Println("a", lsA, "b", lsB)
	fmt.Println("key", loop(pkB, lsA))
}

func determineLoopSize(key int) int {
	count := 0
	for val := 1; val != key; count++ {
		val *= 7
		val %= 20201227
	}
	return count
}

func loop(subj, count int) int {
	key := 1
	for i := 0; i < count; i++ {
		key *= subj
		key %= 20201227
	}
	return key
}
