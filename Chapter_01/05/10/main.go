package main

import (
	"fmt"
	"math"
)

func main() {
	kayak := 275
	soccerBall := 19.50

	total := kayak + int(math.Round(soccerBall))

	fmt.Println(total)
}