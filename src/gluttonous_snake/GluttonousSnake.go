package gluttonous_snake

import (
	"math/rand"
	"time"
)

const weight, hight = 16, 9

type Snake [weight][hight]int

func (s *Snake) SetBlock(b int, p [2]int) {
	s[p[0]][p[1]] = b
}

func (s *Snake) GetBlock(p [2]int) int {
	return (*s)[p[0]][p[1]]
}

func RunGame(direction *int, printer chan<- Snake) (sorce int) {
	var cDirection [2]int
	var hand [2]int = [2]int{2, hight / 2}
	var s Snake
	s[0][hight/2], s[1][hight/2], s[2][hight/2] = 1, 2, 3
	addFood(&s)
	for {
		switch *direction {
		case 1: //up
			cDirection = [2]int{0, -1}
		case 2: //down
			cDirection = [2]int{0, 1}
		case 3: //left
			cDirection = [2]int{-1, 0}
		case 4: //right
			cDirection = [2]int{1, 0}
		}
		var newHand [2]int = [2]int{(hand[0] + cDirection[0] + weight) % weight, (hand[1] + cDirection[1] + hight) % hight}
		b := s.GetBlock(newHand)
		switch {
		case b == 0:
			through(&s)
			s.SetBlock(s.GetBlock(hand)+1, newHand)
		case b < 0:
			s.SetBlock(s.GetBlock(hand)+1, newHand)
			addFood(&s)
		case b > 0:
			return score(s)
		}
		printer <- s
		hand = newHand
		time.Sleep(time.Millisecond * 400)
	}
}

func through(s *Snake) {
	for i := 0; i < weight; i++ {
		for j := 0; j < hight; j++ {
			if s.GetBlock([2]int{i, j}) != 0 {
				s.SetBlock(s.GetBlock([2]int{i, j})-1, [2]int{i, j})
			}
		}
	}
}

func score(s Snake) (max int) {
	for i := 0; i < weight; i++ {
		for j := 0; j < hight; j++ {
			if max < s.GetBlock([2]int{i, j}) {
				max = s.GetBlock([2]int{i, j})
			}
		}
	}
	return
}

func addFood(s *Snake) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var x, y int = r.Intn(weight), r.Intn(hight)
	for s.GetBlock([2]int{x, y}) != 0 {
		x, y = r.Intn(weight+1), r.Intn(hight+1)
	}
	s.SetBlock(-1, [2]int{x, y})
}
