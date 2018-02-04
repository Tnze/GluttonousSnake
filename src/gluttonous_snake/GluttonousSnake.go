package gluttonous_snake

import (
	"math/rand"
	"time"
)

//地图大小
const (
	Weight = 16
	Hight  = 9
)

//Snake 结构表示蛇
type Snake struct {
	s             [Weight][Hight]int
	hand          [2]int
	isEnd         bool
	lastDirection [2]int
}

//SetBlock 方法设置地图上的元素
func (s *Snake) SetBlock(b int, p [2]int) {
	s.s[(p[0]+Weight)%Weight][(p[1]+Hight)%Hight] = b
}

//GetBlock 方法获取地图上的元素
func (s *Snake) GetBlock(p [2]int) int {
	return s.s[(p[0]+Weight)%Weight][(p[1]+Hight)%Hight]
}

//NewSnake 函数初始化一条新蛇
func NewSnake() *Snake {
	var s Snake
	s.s[0][Hight/2], s.s[1][Hight/2], s.s[2][Hight/2] = 1, 2, 3
	s.lastDirection = [2]int{1, 0}
	s.hand = [2]int{2, Hight / 2}
	addFood(&s)
	return &s
}

//Step 方法使游戏演化
func (s *Snake) Step(direction int) (sorce int, isEnd bool) {
	if s.isEnd { //判断游戏是否已经结束
		return score(s), true
	}
	var cDirection [2]int

	switch direction { //解析方向
	case 1: //up
		cDirection = [2]int{0, -1}
	case 2: //down
		cDirection = [2]int{0, 1}
	case 3: //left
		cDirection = [2]int{-1, 0}
	case 4: //right
		cDirection = [2]int{1, 0}
	default:
		cDirection = s.lastDirection
	}
	s.lastDirection = cDirection
	//计算蛇头下一步的位置
	var newHand = [2]int{(s.hand[0] + cDirection[0] + Weight) % Weight, (s.hand[1] + cDirection[1] + Hight) % Hight}
	b := changeBlock( s.GetBlock(newHand)) //取蛇头将要碰到的物体
	switch {
	case b == 0: //如果是空的
		through(s)
		s.SetBlock(s.GetBlock(s.hand)+1, newHand)
	case b < 0: //如果是食物
		s.SetBlock(s.GetBlock(s.hand)+1, newHand)
		addFood(s)
	case b > 0: //如果是蛇身
		s.isEnd = true
		return score(s), true
	}
	s.hand = newHand

	return score(s), false
}

//使蛇变短
func through(s *Snake) {
	for i := 0; i < Weight; i++ {
		for j := 0; j < Hight; j++ {
			location := [2]int{i,j}
			s.SetBlock(changeBlock(s.GetBlock(location)), location)
		}
	}
}

func changeBlock(b int) int {
	if b !=0 {
		return b-1
	}
	return 0
}

//计算分数
func score(s *Snake) (max int) {
	for i := 0; i < Weight; i++ {
		for j := 0; j < Hight; j++ {
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
	var x, y int = r.Intn(Weight), r.Intn(Hight)
	for s.GetBlock([2]int{x, y}) != 0 {
		x, y = r.Intn(Weight+1), r.Intn(Hight+1)
	}
	s.SetBlock(-1, [2]int{x, y})
}
