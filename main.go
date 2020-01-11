package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	newDeckOptionNumber = "1"
	printOptionNumber = "2"
	shuffleOptionNumber = "3"
	dealOptionNumber = "4"
	saveADeckOptionNumber = "5"
	getADeckOptionNumber = "6"
	closeOptionNumber = "7"
)

func main() {

	// Options for the user
	fmt.Println("Enter a number 1 to 6 to pick an option!")
	fmt.Println(newDeckOptionNumber, ": New deck")
	fmt.Println(printOptionNumber, ": Print")
	fmt.Println(shuffleOptionNumber, ": Shuffle")
	fmt.Println(dealOptionNumber, ": Deal")
	fmt.Println(saveADeckOptionNumber, ": Save a deck to a file")
	fmt.Println(getADeckOptionNumber, ": get a deck from a file")
	fmt.Println(closeOptionNumber, ": close")

	// gets user input
	fmt.Print("Enter the number identifying your action :")
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)

	// processes the user input
	if userInput == newDeckOptionNumber {
		fmt.Println("New deck")
		NewDeck()

	} else if userInput == printOptionNumber {
		fmt.Println("Print")

	} else {
		fmt.Println("not a correct input :", userInput)
	}
}
