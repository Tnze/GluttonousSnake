package gluttonous_snake

import (
	"math/rand"
	"time"
)

//地图大小
const weight, hight = 16, 9

//表示蛇
type Snake [weight][hight]int

func (s *Snake) SetBlock(b int, p [2]int) {
	s[p[0]][p[1]] = b
}

func (s *Snake) GetBlock(p [2]int) int {
	return (*s)[p[0]][p[1]]
}

//运行游戏（支持并行运行）
func RunGame(direction *int, printer chan<- Snake) (sorce int) {
	var cDirection [2]int
	var hand [2]int = [2]int{2, hight / 2}
	var s Snake
	s[0][hight/2], s[1][hight/2], s[2][hight/2] = 1, 2, 3
	addFood(&s)
	for {
		switch *direction { //解析方向
		case 1: //up
			cDirection = [2]int{0, -1}
		case 2: //down
			cDirection = [2]int{0, 1}
		case 3: //left
			cDirection = [2]int{-1, 0}
		case 4: //right
			cDirection = [2]int{1, 0}
		}
		//计算蛇头下一步的位置
		var newHand [2]int = [2]int{(hand[0] + cDirection[0] + weight) % weight, (hand[1] + cDirection[1] + hight) % hight}
		b := s.GetBlock(newHand) //取蛇头将要碰到的物体
		switch {
		case b == 0: //如果是空的
			through(&s)
			s.SetBlock(s.GetBlock(hand)+1, newHand)
		case b < 0: //如果是食物
			s.SetBlock(s.GetBlock(hand)+1, newHand)
			addFood(&s)
		case b > 0: //如果是蛇身
			return score(s)
		}
		printer <- s
		hand = newHand
		time.Sleep(time.Millisecond * 400)
	}
}

//使蛇变短
func through(s *Snake) {
	for i := 0; i < weight; i++ {
		for j := 0; j < hight; j++ {
			if s.GetBlock([2]int{i, j}) != 0 {
				s.SetBlock(s.GetBlock([2]int{i, j})-1, [2]int{i, j})
			}
		}
	}
}

//计算分数
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

//随机添加食物
func addFood(s *Snake) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var x, y int = r.Intn(weight), r.Intn(hight)
	for s.GetBlock([2]int{x, y}) != 0 {
		x, y = r.Intn(weight+1), r.Intn(hight+1)
	}
	s.SetBlock(-1, [2]int{x, y})
}
