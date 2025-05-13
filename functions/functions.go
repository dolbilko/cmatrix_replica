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


func Move_cursor(x, line_head int) {
	fmt.Printf("\u001b[%d;%dH", line_head, x)
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
			for line_head := 0; line_head < terminal_height; line_head++ {
				Move_cursor(x, line_head)
				fmt.Printf("%c", Random_char(symbols))
				for line1 := tail_start; line1 > 0; line1-- {
					Move_cursor(x, line1)
					fmt.Print(" ")
				} 
				for line2 := line_head+1; line2 < terminal_height; line2++ {
					Move_cursor(x, line2)
					fmt.Print(" ")
				}
				tail_start++
				time.Sleep(time.Duration(speed) * time.Millisecond)
			}
			for line_head := terminal_height - drop_length; line_head < terminal_height; line_head++ {
				Move_cursor(x, line_head)
				fmt.Print(" ")
				time.Sleep(time.Duration(speed) * time.Millisecond)
			}
			for line := 0; line > terminal_height; line++ {
				Move_cursor(x, line)
				fmt.Print(" ")
			}
			time.Sleep(time.Duration(rand.Intn(50)+10) * time.Millisecond)
		}
	}
}