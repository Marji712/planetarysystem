package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Solar System!")
	fmt.Println("There are 9 planets to explore.")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	name, _ := reader.ReadString('\n')
	fmt.Println("Nice to meet you,",name, ". My name is Eliza, I'm an old friend of Alexa")
	fmt.Println("Let's go on an adventure!")
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Shall I randomly choose a planet for you to vist? (Y or N)")
	answer, _ := reader.ReadString('\n')
	if strings.TrimRight(answer, "\n") == "Y" {
		fmt.Println("Name the planet you would like to visit")
	} else if strings.TrimRight(answer, "\n") == "N" {
		fmt.Println("Put random picker here")
	} else {
		fmt.Println("Throw error and let's go back to the question, idiot")
	}
	fmt.Println("Traveling to  ...")
	fmt.Println("Arrived a .  ")

}
