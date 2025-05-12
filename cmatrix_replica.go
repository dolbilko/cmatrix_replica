package main

import (
	"fmt"
 	"math/rand"
	"time"
)


var symbols = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@#$%^&*")

func main() {
	fmt.Print("\u001b[2J")
	fmt.Printf("\u001b[%d;%dH", 5, 7)
	fmt.Print("а")
	time.Sleep(time.Duration(5 * time.Second))
}