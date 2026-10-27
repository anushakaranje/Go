package main

import (
	"fmt"
	"math/rand"
)

func main() {
	number := rand.Intn(100) + 1
	var guess int
	attempts := 0

	fmt.Println("🎮 NUMBER GUESSING GAME")
	fmt.Println("Guess a number between 1 and 100!")

	for {
		fmt.Print("Enter your guess: ")
		_, err := fmt.Scan(&guess)

		if err != nil {
			fmt.Println("Please enter a valid number!")
			return
		}

		attempts++

		if guess < number {
			fmt.Println("Too low! Try again.")
		} else if guess > number {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Println("🎉 Correct! You won!")
			fmt.Println("Attempts:", attempts)
			break
		}
	}
	fmt.Println("Secret number:", number)
}
