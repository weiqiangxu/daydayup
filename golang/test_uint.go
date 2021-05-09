package main

import "fmt"

func main() {
	var a uint8
	var b uint8
	// []byte string
	a = 5
	b = 6
	fmt.Println("a-b=", a-b) // unit8相减得到一个unit8数字，当变成-1的时候超出限制，变为255
	fmt.Println("b-a=", b-a)

	// int8的取值范围为-128~127
	// 其实是一个二进制 - 8位，而8个二进制数量最大的是 11111111 ==》 255

	// 0 000 0101
	// 0 000 0110
	// 0 000 0001

	// 0 1111111 第一位是符号位 - 后面的最大是127（转10进制）如果第一位是1表示unsign无符号 - 最大就是256
	// 0111111 -- 255
}
