// программа, реализующая эффект матрицы в терминале
// необходима поддержка ANSI последовательностей

package main

import (
	"cmatrix_replica/functions"
	"fmt"
)
	

func main() {
	fmt.Print("\u001b[?25l")
	ended := make(chan struct{})
	functions.Terminal_clear()
	for x := 1; x < 50; x+=2 {
		go functions.Drop_render(x, 30, ended)
	}
	go functions.Q_catching(ended)
	<-ended
	functions.Terminal_clear()
	fmt.Print("\u001b[0;0H")
}