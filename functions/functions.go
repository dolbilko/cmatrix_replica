package functions

import (
	"fmt"
 	"math/rand"
	"time"
	"github.com/eiannone/keyboard"
)


var symbols = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@#$%&")


func random_char() rune {
	return symbols[rand.Intn(len(symbols))]
}


func draw_char(col, line int, char rune) {
	fmt.Printf("\u001b[%d;%dH%c", line, col, char)
}


func clear_char(col, line int) {
	fmt.Printf("\u001b[%d;%dH ", line, col)
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


func Drop_render(col, terminal_height int, ended <-chan struct{}){
	time.Sleep(time.Duration(rand.Intn(1000) + 500) * time.Millisecond)
	for {
		select {
		case <-ended:
			return
		default:
			drop_length := rand.Intn(20)+5
			speed := time.Duration(rand.Intn(100) + 50) * time.Millisecond
			for line_head := -drop_length; line_head < terminal_height + drop_length; line_head++ {
				if line_head - drop_length > 0 {
					clear_char(col, line_head - drop_length)
				}
				draw_char(col, line_head, random_char())
				time.Sleep(speed)
			}
			time.Sleep(time.Duration(rand.Intn(500)+10) * time.Millisecond)
		}
	}
}