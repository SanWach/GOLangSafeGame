/*
 * Author: Mr_Akihiro
 * Description: This file contains the implementation of a simple number guessing game logic.
 *              The user is prompted to guess the digits of a secret passcode one by one.
 * License: MIT License | Feel free to use it and build on it!
 */

package main

import (
	"fmt"
)

func main() {
	secretPasscode := 111 // You can enter any number you want in here

	firstDigit := secretPasscode / 100        // Extracts the first digit
	secondDigit := (secretPasscode / 10) % 10 // Extracts the second digit
	thirdDigit := secretPasscode % 10         // Extracts the third digit

	var guess int

	// First digit guess
	fmt.Print("Please enter your first guess: ")
	fmt.Scan(&guess)
	if guess == firstDigit {
		fmt.Println("Result OK! Please memorize this number")
	} else {
		fmt.Println("Wrong digit! Please restart your attempt!")
		return // Exits if the first guess is wrong
	}

	// Second digit guess
	fmt.Print("Please guess your second number: ")
	fmt.Scan(&guess)
	if guess == secondDigit {
		fmt.Println("Result OK! Please memorize this number")
	} else {
		fmt.Println("Wrong digit! Please restart your attempt!")
		return // Exits if the second guess is wrong
	}

	// Third digit guess
	fmt.Print("Please guess your third number: ")
	fmt.Scan(&guess)
	if guess == thirdDigit {
		fmt.Println("Result OK! You've guessed all digits correctly!")
	} else {
		fmt.Println("Wrong digit! Please restart your attempt!")
		return // Exits if the third guess is wrong
	}

	fmt.Println("Congratulations! You've unlocked the secret passcode.")
}
