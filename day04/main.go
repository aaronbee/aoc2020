package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/aristanetworks/glog"
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
		if valid2(s.Text()) {
			count++
		}
	}
	fmt.Println(count)
}

func valid1(s string) bool {
	m := make(map[string]string)
	for _, f := range strings.Fields(s) {
		switch f[:3] {
		case "byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid":
			m[f[:3]] = f[4:]
		}
	}
	_, cidPresent := m["cid"]
	return len(m) == 8 || (len(m) == 7 && !cidPresent)
}

func valid2(s string) bool {
	var validFields int
	for _, f := range strings.Fields(s) {
		k, v := f[:3], f[4:]
	fieldswitch:
		switch k {
		case "byr":
			if len(v) != 4 {
				glog.Info("Bad byr:", v)
				break
			}
			if v < "1920" || v > "2002" {
				glog.Info("Bad byr:", v)
				break
			}
			validFields++
		case "iyr":
			if len(v) != 4 {
				glog.Info("Bad iyr:", v)
				break
			}
			if v < "2010" || v > "2020" {
				glog.Info("Bad iyr:", v)
				break
			}
			validFields++
		case "eyr":
			if len(v) != 4 {
				glog.Info("Bad eyr:", v)
				break
			}
			if v < "2020" || v > "2030" {
				glog.Info("Bad eyr:", v)
				break
			}
			validFields++
		case "hgt":
			if strings.HasSuffix(v, "cm") {
				if len(v) != 5 {
					glog.Info("Bad hgt:", v)
					break
				}
				v = v[:3]
				if v < "150" || v > "193" {
					glog.Info("Bad hgt:", v)
					break
				}
				validFields++
			} else if strings.HasSuffix(v, "in") {
				if len(v) != 4 {
					glog.Info("Bad hgt:", v)
					break
				}
				v = v[:2]
				if v < "59" || v > "76" {
					glog.Info("Bad hgt:", v)
					break
				}
				validFields++
			}
		case "hcl":
			if len(v) != 7 || v[0] != '#' {
				glog.Info("Bad hcl:", v)
				break
			}
			for _, c := range v[1:] {
				if (c >= 'a' && c <= 'f') || (c >= '0' && c <= '9') {
					continue
				}
				glog.Info("Bad hcl: ", v, " char: ", string(c))
				break fieldswitch
			}
			validFields++
		case "ecl":
			switch v {
			case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
				validFields++
			default:
				glog.Info("Bad ecl:", v)
			}
		case "pid":
			if len(v) != 9 {
				break
			}
			for _, c := range v {
				if c < '0' || c > '9' {
					glog.Info("Bad pid:", v)
					break fieldswitch
				}
			}
			validFields++
		}
	}
	return validFields == 7
}
