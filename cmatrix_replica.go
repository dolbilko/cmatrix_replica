package main

import (
	"fmt"
 	// "math/rand"
	"time"
)

func move_cursor(x, line int) {
	fmt.Printf("\u001b[%d;%dH", line, x)
}

func terminal_clear() {
	fmt.Print("\u001b[2J")
}

func tail_remover(x, line_start, line_finish int) {
	for line := line_start; line < line_finish; line++ {
		move_cursor(x, line)
		fmt.Print(" ")
		time.Sleep(time.Duration(50 * time.Millisecond))
	}
}

func drop_render(x, terminal_height int){
	for {
		drop_length := 5
		tail := 0
		for line := 0; line < terminal_height; line++ {
			move_cursor(x, line)
			fmt.Print("1")
			tail = line-drop_length
			if tail > 0 {
				move_cursor(x, tail)
				fmt.Print(" ")
			}
			time.Sleep(time.Duration(50 * time.Millisecond))
			if terminal_height-line < 2 {
				go tail_remover(x, terminal_height-drop_length, terminal_height)
				// time.Sleep(time.Duration(50 * time.Millisecond))
			}
		}
	}
}

// var symbols = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@#$%^&*")

func main() {
	terminal_clear()
	go drop_render(0, 20)
	// for x := 0; x < 20; x++ {
	// 	go drop_render(x, 10)
	// }
	// time.Sleep(time.Duration(5 * time.Second))
	time.Sleep(time.Duration(5 * time.Minute))
}