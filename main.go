package main

import (
	"cellbesmanos/exercise"
	"fmt"
	"math"
)

// returns the total sum of the numbers from 1 up to the "to" parameter
// ex.: if to is 2 then 1 + 2, it will return 3
func getTotalNumbersSumTo(to int) int {
	sum := 0

	// the initialization and post statements are optional
	for i := 0; i < to; i++ {
		sum += i
	}

	return sum
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

// crazy stuff
func pow(x, n, lim float64) float64 {
	// an else statement chained here will also have access to the v variable
	if v := math.Pow(x, n); v < lim {
		return v
	}

	return lim
}

func main() {
	fmt.Println("Hello World!")

	fmt.Println(getTotalNumbersSumTo(10))

	// go's while loop
	num := 1
	for num < 1000 {
		num += num
	}
	fmt.Println(num)

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))

	fmt.Println(exercise.NumWithClosestSquareTo(20))
}