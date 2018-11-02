package main

import (
	"ai"
	"draw"
	"fmt"
	gs "gluttonous_snake"
	"time"
)

func main() {
	snake := gs.NewSnake()
	direction := 0
	t := time.Now()
	err := draw.OpenWindow(func() *gs.Snake {
		if time.Now().Sub(t)/time.Millisecond > 100 {
			score, isEnd := snake.Step(ai.NextStep(snake))
			fmt.Printf("\rscore:	%d", score)
			if isEnd {
				fmt.Printf("\rFinal score:	%d\n", score)
				time.Sleep(4 * time.Second)
				snake = gs.NewSnake()
				direction = 0
				fmt.Println(direction)
			}
			t = time.Now()
		}
		time.Sleep(20 * time.Millisecond)
		return snake
	}, func(dir int) {})
	if err != nil {
		panic(err)
	}
}
