package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	var sum int
	scn := bufio.NewScanner(f)
	for scn.Scan() {
		s := strings.ReplaceAll(scn.Text(), "(", "( ")
		s = strings.ReplaceAll(s, ")", " )")
		tokens := strings.Fields(s)
		v, i := eval(tokens)
		if i != len(tokens) {
			panic(fmt.Errorf("tokens not all consumed: Exp %d Got: %d", i, len(tokens)))
		}
		sum += v
	}
	fmt.Println("final sum", sum)
}

type state struct {
	left int
	oper string
}

func (s *state) apply(val int) {
	if s.oper == "" {
		s.left = val
		return
	}
	switch s.oper {
	case "+":
		s.left += val
	case "*":
		s.left *= val
	}
	s.oper = ""
}

func (s *state) setOper(oper string) {
	if s.oper != "" {
		panic("oper already set")
	}
	s.oper = oper
}

func eval(tokens []string) (int, int) {
	var s state
	for i := 0; i < len(tokens); i++ {
		switch t := tokens[i]; t {
		case ")":
			return s.left, i + 1
		case "(":
			val, consumed := eval(tokens[i+1:])
			s.apply(val)
			i += consumed
			if tokens[i] != ")" {
				panic(fmt.Errorf("bad return from eval: Exp \")\", Got: %q", tokens[i]))
			}
		case "+", "*":
			s.setOper(t)
		default:
			val, err := strconv.Atoi(t)
			if err != nil {
				panic(err)
			}
			s.apply(val)
		}
	}
	return s.left, len(tokens)
}
