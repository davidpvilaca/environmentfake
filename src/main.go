package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
* Generate a random float64 botween 18-35
* and return
**/
func getRandTemp() float64 {
	rand.Seed(time.Now().UnixNano())
	min := 18
	max := 35
	return float64(rand.Intn(max-min)+min) + rand.Float64()
}

// Main function
// infinity loop generating random sleeping 1s
func main() {
	for {
		temp := getRandTemp()
		for i := 1; i < 10001; i++ {
			temp = (temp + getRandTemp()) / 2
		}
		fmt.Println(temp)
		time.Sleep(time.Second)
	}
}
