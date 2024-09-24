package main

import "fmt"

func main() {
	fmt.Println("Преобразование")

	kayak := 275
	soccerBall := 19.50

	total := kayak + int(soccerBall)
	//total := kayak + soccerBall - ошибка invalid operation: kayak + soccerBall (mismatched types int and float64)

	//total := float64(kayak) + soccerBall

	fmt.Println(total)
	fmt.Println(int8(total))
}