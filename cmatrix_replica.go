// программа, реализующая эффект матрицы в терминале
// необходима поддержка ANSI последовательностей

package main

import (
	"cmatrix_replica/functions"
	"fmt"
)
	

func main() {
	fmt.Print("\u001b[?25l") // спрятать курсор
	ended := make(chan struct{}) 
	functions.Terminal_clear()
	for x := 1; x < 101; x+=2 { // запуск горутин рендера "капель"
		go functions.Drop_render(x, 40, ended)
	}
	go functions.Q_catching(ended)
	<-ended //ожидание закрытия канала чтобы его прочесть и идти дальше
	functions.Terminal_clear()
	fmt.Print("\u001b[0;0H\u001b[0m") // вернуть курсор в 0;0, сбросить форматирование
}