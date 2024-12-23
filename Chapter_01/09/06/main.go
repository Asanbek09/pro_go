package main

import "fmt"

type calcFunc func(float64) float64

func printPrice(product string, price float64, calculator calcFunc) {
    fmt.Println("Product:", product, "Price:", calculator(price))
}

func selectCalculator(price float64) calcFunc {
    if (price > 100) {
        var withTax calcFunc = func(price float64) float64 {
            return price + (price * 0.2)
        }
        return withTax
    }
    withoutTax := func(price float64) float64 {
        return price
    }
    return withoutTax
}

func main() {
    products := map[string]float64 {
        "kayak:": 275,
        "jacket": 48.95,
    }

    for product, price := range products {
        printPrice(product, price, selectCalculator(price))
    }
}