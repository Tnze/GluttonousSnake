package ai

import (
	gs "github.com/Tnze/GluttonousSnake/gluttonous_snake"
)

//import "fmt"

func findFood(s *gs.Snake) (path []int, canFindPath bool) {
	O, _, Food := getOAF(s)
	return aStar(s, O, Food)
}

//getOAF 函数搜索蛇头、蛇尾和食物, O是蛇头，F是蛇尾，Food是食物
func getOAF(s *gs.Snake) (O, F, Food [2]int) {
	max := 0
	for i := 0; i < gs.Weight; i++ {
		for j := 0; j < gs.Hight; j++ {
			pos := [2]int{i, j}
			block := s.GetBlock(pos)
			if block < 0 {
				Food = pos
			} else if block == 1 {
				F = pos
			} else if block > max {
				max = block
				O = pos
			}
		}
	}
	return
}

type Cell struct {
	Pos  [2]int //坐标
	H, G int
	ln   *Cell //前驱结点
}

//得到可以联通的结点
func (c *Cell) GetNeighborhoodsPos() [][2]int {
	list := [][2]int{
		[2]int{c.Pos[0] + 1, c.Pos[1]},
		[2]int{c.Pos[0] - 1, c.Pos[1]},
		[2]int{c.Pos[0], c.Pos[1] + 1},
		[2]int{c.Pos[0], c.Pos[1] - 1}}

	return list
}

//GetF 函数获取F值
func (c *Cell) GetF() int {
	return c.H + c.G
}

func NewCell(pos [2]int, G, H int, ln *Cell) *Cell {
	var c Cell
	c.Pos = pos
	c.G = G
	c.H = H
	c.ln = ln
	return &c
}

//ManhattanDistance 返回两个点之间的曼哈顿距离
func ManhattanDistance(p1, p2 [2]int) int {
	x := p1[0] - p2[0]
	y := p1[1] - p2[1]
	for x > 0 {
		x -= gs.Weight
	}
	for y > 0 {
		y -= gs.Hight
	}
	for x < 0 {
		x += gs.Weight
	}
	for y < 0 {
		y += gs.Hight
	}
	x = min(x, gs.Weight-x)
	y = min(y, gs.Hight-y)
	return x + y
}
func min(n1, n2 int) int {
	if n1 > n2 {
		return n2
	} else {
		return n1
	}
}

//abs 是绝对值函数
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

//从终点生成路径
func (c *Cell) getPath() []int {
	stack := NewStack()
	now := c
	for now.ln != nil {
		stack.Push(getDirection(now.ln.Pos, now.Pos))
		now = now.ln
	}
	path := make([]int, 0)
	for {
		v, b := stack.Pop()
		if !b {
			//fmt.Println("path:  ", path)
			return path
		}
		path = append(path, v)
	}
}

func findTail(s *gs.Snake) (path []int, canFindPath bool) {
	O, F, _ := getOAF(s) //获取蛇头蛇尾
	return aStar(s, O, F)
}

func aStar(s *gs.Snake, p1, p2 [2]int) ([]int, bool) {
	open := NewQueue()
	close := NewQueue()
	open.push(NewCell(p1, 0 /* ManhattanDistance(p1, p2)*/, 0, nil)) //将起点加入到Open
	for {
		//fmt.Println("Running.......")
		U, bErr := open.popMinF()
		if !bErr {
			return nil, false
		}
		close.push(U)
		if U.Pos == p2 {
			return U.getPath(), true
		}
		VPoses := U.GetNeighborhoodsPos() //取U的邻居
		for _, v := range VPoses {
			if s.GetBlock(v) > 1 {
				continue
			}
			v = gs.FormattingCoordinates(v)
			if close.exist(v) {
				//若V已在close表中
				//什么都不做^_^
			} else if open.exist(v) {
				//若V在open表中
				if open.get(v).G > U.G+1 {
					open.get(v).G = U.G + 1
					open.get(v).ln = U
				}
			} else {
				//若V不在open表中
				newCell := NewCell(v, U.G+1, 0 /*ManhattanDistance(v, p2)*/, U)
				open.push(newCell)
			}
		}
	}
}
