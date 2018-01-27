package main

import (
	"fmt"
	termbox "github.com/nsf/termbox-go"
	gs "gluttonous_snake"
	"os"
)

func main() {

	defer termbox.Close()
	var score int
	for {
		err := termbox.Init()
		if err != nil {
			panic(err)
		}
		d := [...]int{1, 0}
		go func() {
			for d[0] == 0 || d[1] == 0 {
				ev := termbox.PollEvent()
				switch {
				case ev.Key == termbox.KeyArrowUp:
					d[0], d[1] = 0, -1
				case ev.Key == termbox.KeyArrowDown:
					d[0], d[1] = 0, 1
				case ev.Key == termbox.KeyArrowLeft:
					d[0], d[1] = -1, 0
				case ev.Key == termbox.KeyArrowRight:
					d[0], d[1] = 1, 0
				case ev.Key == termbox.KeyEsc:
					os.Exit(0)
				}
				//fmt.Println("dir:", d)
			}
		}()
		score = gs.RunGame(&d)
		termbox.Close()
		fmt.Println("score:", score)

	}
}
