package functions

import (
	"fmt"
 	"math/rand"
	"time"
	// "sync"
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
			time.Sleep(time.Duration(rand.Intn(1000) + 50) * time.Millisecond)
			drop_length := rand.Intn(terminal_height/4)+3
			speed := rand.Intn(100)+50
			tail_start := -drop_length
			for line := 0; line < terminal_height; line++ {
				Move_cursor(x, line)
				fmt.Printf("%c", Random_char(symbols))
				if tail_start >= 0 && tail_start < terminal_height {
					Move_cursor(x, tail_start)
					fmt.Print(" ")
				}
				tail_start++
				time.Sleep(time.Duration(speed) * time.Millisecond)
			}
			for line := terminal_height - drop_length; line != terminal_height; line++ {
				Move_cursor(x, line)
				fmt.Print(" ")
				time.Sleep(time.Duration(speed) * time.Millisecond)
			}
			time.Sleep(time.Duration(rand.Intn(50)+10) * time.Millisecond)
		}
	}
}