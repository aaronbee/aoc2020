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

	var rules []rule

	scn := bufio.NewScanner(f)
	for scn.Scan() {
		line := scn.Text()
		if line == "" {
			break
		}

		splt := strings.Split(line, ":")
		index, err := strconv.Atoi(splt[0])
		if err != nil {
			panic(err)
		}
		if index >= len(rules) {
			rules = append(rules, make([]rule, index-len(rules)+1)...)
		}

		splt = strings.Split(splt[1], "|")
		var r or
		for _, opt := range splt {
			r = append(r, parseRule(strings.Fields(opt)))
		}
		rules[index] = r
	}

	rules[8] = &loop{a: seq{42}, b: seq{42, 8}}
	rules[11] = &loop{a: seq{42, 31}, b: seq{42, 11, 31}}

	var count int
outer:
	for scn.Scan() {
		for i := 0; i < len(scn.Bytes()); i++ {
			for j := 0; j < len(scn.Bytes()); j++ {
				rules[8].(*loop).count = i
				rules[11].(*loop).count = j
				iter := tokenIter{tokens: scn.Bytes()}
				if rules[0].match(rules, &iter) && iter.i == len(iter.tokens) {
					count++
					continue outer
				}
			}
		}
	}
	fmt.Println(count)
}

func parseRule(tokens []string) rule {
	var sq seq
	for i, t := range tokens {
		if t[0] == '"' {
			if len(tokens) != 1 {
				panic("expected one literal")
			}
			s, err := strconv.Unquote(t)
			if err != nil {
				panic(err)
			}
			if len(s) != 1 {
				panic("unexpected size of literal")
			}
			return byt(s[0])
		}
		if i != len(sq) {
			panic("mixed literal and subrule")
		}
		subrule, err := strconv.Atoi(t)
		if err != nil {
			panic(err)
		}
		sq = append(sq, subrule)
	}
	return sq
}

type tokenIter struct {
	tokens []byte
	i      int
}

type rule interface {
	match(rs []rule, iter *tokenIter) bool
}

type or []rule

func (r or) match(rs []rule, iter *tokenIter) bool {
	backup := iter.i
	for _, r := range r {
		iter.i = backup
		if r.match(rs, iter) {
			return true
		}
	}
	return false
}

type loop struct {
	a     seq
	b     seq
	count int
}

func (r *loop) match(rs []rule, iter *tokenIter) bool {
	if r.count == 0 {
		return r.a.match(rs, iter)
	}
	r.count--
	return r.b.match(rs, iter)
}

type seq []int

func (r seq) match(rs []rule, iter *tokenIter) bool {
	for _, i := range r {
		if !rs[i].match(rs, iter) {
			return false
		}
	}
	return true
}

type byt byte

func (r byt) match(rs []rule, iter *tokenIter) bool {
	if iter.i >= len(iter.tokens) {
		return false
	}
	if iter.tokens[iter.i] == byte(r) {
		iter.i++
		return true
	}
	return false
}
