// файл с подключаемыми функциями для рендера эффекта матрицы

package functions

import (
	"fmt"
 	"math/rand"
	"time"
	"github.com/eiannone/keyboard"
)


var symbols = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@#$%&")


func random_char() rune {
	// функция возвращающая случайную руну из среза выше
	return symbols[rand.Intn(len(symbols))]
}


func draw_char(col, line int, char rune) {
	// получает: col — номер столбца; line — номер строки; char — руну
	// печатает: ANSI последовательность перемещения курсора и зелёного цвета, а после саму руну
	fmt.Printf("\u001b[%d;%dH\u001b[2;32m%c", line, col, char)
}


func clear_char(col, line int) {
	// получает: col — столбец; line — строку
	// печатает пробел на этом месте
	fmt.Printf("\u001b[%d;%dH ", line, col)
}


func Terminal_clear() {
	// печатает ANSI последовательность очистки терминала
	fmt.Print("\u001b[2J")
}


func Q_catching(channel chan struct{}) {
	// получает channel — канал типа struct
	// закрывает канал при обнаружении напечатанной буквы "q"
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
	// получает: col — номер столбца; terminal_height — высоту рабочей области; ended — канал типа struct{}
	// в столбце с полученным номером последовательно вызывает функцию печати символа и пробела
	// в результате получается "капля" определённой длины, скорость которой определяет задержка печати
	time.Sleep(time.Duration(rand.Intn(500) + 100) * time.Millisecond) // задержка для неодновременного старта с другими горутинами
	for {
		select {
		case <-ended:
			return
		default:
			drop_length := rand.Intn(20)+5
			speed := time.Duration(rand.Intn(30) + 15) * time.Millisecond // задержка печати
			for line_head := 0; line_head < terminal_height + drop_length; line_head++ { // цикл отрисовки "капли"
				if line_head < terminal_height {
					draw_char(col, line_head, random_char())
				}
				clear_char(col, line_head - drop_length)
				time.Sleep(speed)
			}
		}
	}
}