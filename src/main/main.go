package main

import (
	"draw"
	"fmt"
	gs "gluttonous_snake"
	"time"
)

func main() {
	snake := gs.NewSnake()
	direction := 0
	t := 0
	draw.OpenWindow(func() *gs.Snake {
		if t > 10 {
			score, isEnd := snake.Step(direction)
			fmt.Printf("\rscore:	%d", score)
			t = 0
			if isEnd {
				snake = gs.NewSnake()
				direction = 0
			}
		}
		t++
		time.Sleep(20 * time.Millisecond)
		return snake
	}, func(dir int) {
		direction = dir
	})
}
