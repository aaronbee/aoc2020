package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)

	shp := &ship{wpX: 10, wpY: 1}
	for s.Scan() {
		line := s.Text()
		action := line[0]
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		shp.move2(action, value)
		fmt.Printf("%s x:%d y:%d wpx:%d wpy:%d\n", line, shp.x, shp.y, shp.wpX, shp.wpY)
	}
	fmt.Println(shp.dist())
}

type ship struct {
	dir int
	x   int
	y   int

	wpX int
	wpY int
}

func (s *ship) dist() int {
	val := s.x
	if s.x < 0 {
		val = -s.x
	}
	if s.y < 0 {
		val -= s.y
	} else {
		val += s.y
	}
	return val
}

func (s *ship) move2(action byte, value int) {
	switch action {
	case 'N':
		s.wpY += value
	case 'S':
		s.wpY -= value
	case 'E':
		s.wpX += value
	case 'W':
		s.wpX -= value
	case 'L':
		switch value {
		case 90:
			s.wpX, s.wpY = -s.wpY, s.wpX
		case 180:
			s.wpX, s.wpY = -s.wpX, -s.wpY
		case 270:
			s.move2('R', 90)
		}
	case 'R':
		switch value {
		case 90:
			s.wpX, s.wpY = s.wpY, -s.wpX
		case 180:
			s.wpX, s.wpY = -s.wpX, -s.wpY
		case 270:
			s.move2('L', 90)
		}
	case 'F':
		s.x += s.wpX * value
		s.y += s.wpY * value
	default:
		panic(fmt.Errorf("unexpected action: %s", string(action)))
	}
}

func (s *ship) move(action byte, value int) {
	switch action {
	case 'N':
		s.y += value
	case 'S':
		s.y -= value
	case 'E':
		s.x += value
	case 'W':
		s.x -= value
	case 'L':
		s.dir += value
		s.dir %= 360
	case 'R':
		s.dir -= value
		s.dir = (s.dir + 360) % 360
	case 'F':
		switch s.dir {
		case 0:
			s.x += value
		case 90:
			s.y += value
		case 180:
			s.x -= value
		case 270:
			s.y -= value
		default:
			panic(fmt.Errorf("unexpected dir: %d", s.dir))
		}
	default:
		panic(fmt.Errorf("unexpected action: %s", string(action)))
	}
}
