package ai

import gs "gluttonous_snake"
import "fmt"

//NextStep 函数完成贪吃蛇ai的工作
func NextStep(s *gs.Snake) (direction int) {
	var canFindPath, canFindTail = false, false
	path, canFindPath := findFood(s)

	if canFindPath {
		fictitiousSnake := *s
		trySnake, err := moveSnake(&fictitiousSnake, path)
		if err != nil {
			panic(err)
		}
		_, canFindTail = findTail(trySnake)
		if canFindTail {
			return path[0]
		}
	}
	fmt.Println("安全步骤")
	return getSafestStep(s)
}

func moveSnake(s *gs.Snake, path []int) (*gs.Snake, error) {
	var err error
	for _, v := range path {
		_, isEnd := s.Step(v)
		if isEnd {
			err = fmt.Errorf("模拟执行路径过程中游戏意外结束")
		}
	}
	return s, err
}

func getSafestStep(s *gs.Snake) int {
	O, F, Food := getOAF(s)
	cO := NewCell(O, 0, 0, nil)
	n := cO.GetNeighborhoodsPos()
	var d, max int
	first := true
	for _, v := range n {
		v = gs.FormattingCoordinates(v)
		_, canFindTail := aStar(s, v, F)
		if canFindTail && s.GetBlock(v) <= 0 {
			if ManhattanDistance(O, Food) >= max {
				d = getDirection(O, v)
				max = ManhattanDistance(O, Food)
			} else if first {
				first = false
				d = getDirection(O, v)
				max = ManhattanDistance(O, Food)
			} else {
				fmt.Println("找不到尾巴")
			}
		}
	}
	if max == 0 {
		fmt.Println("安全计算失败")
	}
	return d
}

func getDirection(p1, p2 [2]int) int {
	switch {
	case (p1[0]-p2[0]) == 1 || (p1[0]-p2[0]) == 1-gs.Weight:
		return 3
	case (p1[0]-p2[0]) == -1 || (p1[0]-p2[0]) == gs.Weight-1:
		return 4
	case (p1[1]-p2[1]) == 1 || (p1[1]-p2[1]) == 1-gs.Hight:
		return 1
	case (p1[1]-p2[1]) == -1 || (p1[1]-p2[1]) == gs.Hight-1:
		return 2

	}
	fmt.Println("p1 : ", p1, "p2 : ", p2)
	return 0
}
