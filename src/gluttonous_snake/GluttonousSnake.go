package gluttonous_snake

import (
	"fmt"
	"math/rand"
	"time"
)

const weight, hight = 16, 9

type snake [weight][hight]int

func (s *snake) SetBlock(b int, p [2]int) {
	s[p[0]][p[1]] = b
}

func (s *snake) GetBlock(p [2]int) int {
	return (*s)[p[0]][p[1]]
}

func RunGame(direction *[2]int) (sorce int) {
	var hand [2]int = [2]int{2, hight / 2}
	var s snake
	s[0][hight/2], s[1][hight/2], s[2][hight/2] = 1, 2, 3
	addFood(&s)
	for {
		var newHand [2]int = [2]int{(hand[0] + direction[0] + weight) % weight, (hand[1] + direction[1] + hight) % hight}
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

		printSnake(s)
		hand = newHand
		time.Sleep(time.Millisecond * 100)
	}
}

func through(s *snake) {
	for i := 0; i < weight; i++ {
		for j := 0; j < hight; j++ {
			if s.GetBlock([2]int{i, j}) != 0 {
				s.SetBlock(s.GetBlock([2]int{i, j})-1, [2]int{i, j})
			}
		}
	}
}

func score(s snake) (max int) {
	for i := 0; i < weight; i++ {
		for j := 0; j < hight; j++ {
			if max < s.GetBlock([2]int{i, j}) {
				max = s.GetBlock([2]int{i, j})
			}
		}
	}
	return
}

func printSnake(s snake) {
	fmt.Print("\n\n\n")
	for i := 0; i < hight; i++ {
		for j := 0; j < weight; j++ {
			if s.GetBlock([2]int{j, i}) > 0 {
				fmt.Print("●")
			} else if s.GetBlock([2]int{j, i}) < 0 {
				fmt.Print("☀")
			} else {
				fmt.Print("○")
			}
		}
		fmt.Println()
	}
}

func addFood(s *snake) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var x, y int = r.Intn(weight), r.Intn(hight)
	for s.GetBlock([2]int{x, y}) != 0 {
		x, y = r.Intn(weight+1), r.Intn(hight+1)
	}
	s.SetBlock(-1, [2]int{x, y})
}
