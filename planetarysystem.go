package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"encoding/json"
	"io/ioutil"
)

func main() {
	jsonread()

	fmt.Println("Welcome to the Solar System!")
	fmt.Println("There are 9 planets to explore.")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	name, _ := reader.ReadString('\n')
	fmt.Println("Nice to meet you,",name, ". My name is Eliza, I'm an old friend of Alexa")
	fmt.Println("Let's go on an adventure!")

	pickchoice()			//function to take input or random planet

						//close json file

	fmt.Println("Traveling to  ...")
	fmt.Println("Arrived a .  ")

}

func pickchoice() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Shall I randomly choose a planet for you to vist? (Y or N)")
	answer, _ := reader.ReadString('\n')
	if strings.TrimRight(answer, "\n") == "Y" {
		fmt.Println("Name the planet you would like to visit")
	} else if strings.TrimRight(answer, "\n") == "N" {
		fmt.Println("Put random picker here")
	} else {
		pickchoice()
	}
}

func jsonread() {
	jsonFile, err := ioutil.ReadFile(strings.ToLower(os.Args[0]))	
	if err != nil {								//check if there was an error
		fmt.Println(err)
		return
	}
	var planets map[string]interface{}
	json.Unmarshal([]byte(jsonFile), &planets)

	fmt.Println(planets["Mars"])
}
