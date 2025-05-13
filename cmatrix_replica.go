package main

import (
	"fmt"
 	"math/rand"
	"time"
)

var symbols = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@#$%^&*")

func random_char(symbols []rune) rune {
	return symbols[rand.Intn(len(symbols))]
}

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
		time.Sleep(time.Duration(50) * time.Millisecond)
	}
}

func drop_render(x, terminal_height int){
	for {
		drop_length := rand.Intn(terminal_height/4)+3
		tail := 0
		speed := rand.Intn(100)+50
		for line := 0; line < terminal_height; line++ {
			move_cursor(x, line)
			fmt.Printf("%c", random_char(symbols))
			tail = line-drop_length
			if tail > 0 {
				move_cursor(x, tail)
				fmt.Print(" ")
			}
			if terminal_height-line < 2 {
				go tail_remover(x, terminal_height-drop_length, terminal_height)
			}
			time.Sleep(time.Duration(speed) * time.Millisecond)
		}
		// time.Sleep(time.Duration(rand.Intn(100) * time.Millisecond))
	}
}


func main() {
	terminal_clear()
	for x := 1; x < 50; x++ {
		go drop_render(x, 10)
	}
	time.Sleep(time.Duration(5 * time.Second))
	terminal_clear()
	move_cursor(0, 0)
}