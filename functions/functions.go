package functions

import (
	"fmt"
 	"math/rand"
	"time"
	"github.com/eiannone/keyboard"
)

var symbols = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@#$%^&*")

func Random_char(symbols []rune) rune {
	return symbols[rand.Intn(len(symbols))]
}

func Move_cursor(x, line int) {
	fmt.Printf("\u001b[%d;%dH", line, x)
}

func Terminal_clear() {
	fmt.Print("\u001b[2J")
}


func Tail_remover(x, line_start, line_finish, speed int) {
	for line := line_start; line < line_finish; line++ {
		Move_cursor(x, line)
		fmt.Print(" ")
		time.Sleep(time.Duration(speed) * time.Millisecond)
	}
}

func Q_catching(channel chan struct{}) {
	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
			return
		}
		if char == 'q' {
			close(channel)
			return
		}
	}
}

func Drop_render(x, terminal_height int, ended <-chan struct{}){
	for {
		select {
		case <-ended:
			return
		default:
			drop_length := rand.Intn(terminal_height/4)+3
			tail := 0
			speed := rand.Intn(100)+50
			for line := 0; line < terminal_height; line++ {
				Move_cursor(x, line)
				fmt.Printf("%c", Random_char(symbols))
				tail = line-drop_length
				if tail > 0 {
					Move_cursor(x, tail)
					fmt.Print(" ")
				}
				if terminal_height-line < 2 {
					go Tail_remover(x, terminal_height-drop_length, terminal_height, speed)
				}
				time.Sleep(time.Duration(speed) * time.Millisecond)
			}
			time.Sleep(time.Duration(rand.Intn(50)+10) * time.Millisecond)
		}
	}
}
