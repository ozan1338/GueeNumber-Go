package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var userInput int
	var secretNumber int
	var chance int = 1

	rand.Seed(time.Now().Unix())
	secretNumber = rand.Intn(10)


	
	for chance <= 10{
		fmt.Printf("Please Input a Number: ")
		fmt.Scan(&userInput)
		if userInput == secretNumber {
			fmt.Println("You Success Guessing")
			break;
		} else if userInput < secretNumber {
			fmt.Println("Too Low")
		} else if userInput > secretNumber {
			fmt.Println("Too High")
		}
		fmt.Println("you still have chance: ",10 - chance)
		chance++
	}
}