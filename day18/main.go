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
		iter := tokenIter{tokens: tokens}
		v := expression(&iter)
		fmt.Println(scn.Text(), "=", v)
		sum += v
	}
	fmt.Println("final sum", sum)
}

// part 2

// recursive descent parser
//
// expression : factor { "*" factor }
// factor : term { "+" term }
// term : NUM
//      | expression

type tokenIter struct {
	tokens []string
	i      int
}

func (iter *tokenIter) accept(s string) bool {
	if iter.i >= len(iter.tokens) {
		return false
	}
	if s == iter.tokens[iter.i] {
		iter.i++
		return true
	}
	return false
}

func expect(b bool) {
	if !b {
		panic("unexpected")
	}
}

func (iter *tokenIter) num() int {
	i, err := strconv.Atoi(iter.tokens[iter.i])
	if err != nil {
		panic(err)
	}
	iter.i++
	return i
}

func expression(iter *tokenIter) int {
	val := factor(iter)
	for iter.accept("*") {
		val *= factor(iter)
	}
	return val
}

func factor(iter *tokenIter) int {
	val := term(iter)
	for iter.accept("+") {
		val += term(iter)
	}
	return val
}

func term(iter *tokenIter) int {
	if iter.accept("(") {
		i := expression(iter)
		expect(iter.accept(")"))
		return i
	}
	return iter.num()
}

// part 1

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
