package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"strconv"
)

func main() {
	maxNum := 100
	secretNumber := rand.Intn(maxNum)

	fmt.Println("Please input your guess")

	for {
		reader := bufio.NewReader(os.Stdin)

		// 读取第一行输入
		input, err := reader.ReadString('\n')
	
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			continue
		}
	
		input = strings.TrimSuffix(input, "\n")
		input = strings.TrimSuffix(input, "\r")
	
		guess, err := strconv.Atoi(input)
	
		if err != nil {
			fmt.Println("Invalid Input, Please enter a Integer value.", err)
			continue
		}
	
		if guess < secretNumber {
			fmt.Println("Your guess is smaller than the correct number")
		} else if guess > secretNumber {
			fmt.Println("Your guess is bigger than the correct number")
		} else {
			fmt.Println("Correct! So smart!")
			return
		}
	}
}