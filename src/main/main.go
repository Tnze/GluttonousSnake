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
	t := time.Now()
	draw.OpenWindow(func() *gs.Snake {
		if time.Now().Sub(t)/time.Millisecond > 250 {
			score, isEnd := snake.Step(direction)
			fmt.Printf("\rscore:	%d", score)
			if isEnd {
				snake = gs.NewSnake()
				direction = 0
			}
			t = time.Now()
		}
		time.Sleep(20 * time.Millisecond)
		return snake
	}, func(dir int) {
		direction = dir
	})
}
