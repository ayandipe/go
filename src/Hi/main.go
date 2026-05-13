package main

import (
	"bufio"
	keyboard "dummy/Keyboard"
	"dummy/Trimspace"
	"dummy/calculatepaintneeded"
	greeting "dummy/greetings"
	"dummy/greetings/deutsch"
	"dummy/pass_fail"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strconv"
)

func main() {
	fmt.Println("Type your score: ")
	var status string
	grade, err := keyboard.Getfloat()
	if grade >= 60.0 {
		status = "passing"
	} else {
		status = "failed"
	}
	fmt.Println("the grade of", grade, "is", status)

	target := rand.Intn(100)

	fmt.Println("I picked a number from 1-100")
	fmt.Println("Can you guess it?")
	readers := bufio.NewReader(os.Stdin)
	success := false

	for guess := 0; guess < 10; guess++ {

		fmt.Println("You have", 10-guess, "guesses left")
		fmt.Print("make a guess: ")
		inputs, err := readers.ReadString('\n')
		if err != nil {
			log.Fatal(err)
			return
		}
		inputs = Trimspace.Trimspacemanual(inputs)
		guess, err := strconv.Atoi(inputs)
		if err != nil {
			log.Fatal(err)
			return
		}
		if guess < target {
			fmt.Println("Your input is too low")
		} else if guess > target {
			fmt.Println("Your input is too high")
		} else {
			success = true
			fmt.Println("Congrats! your input is accurate")
			break
		}
		// fmt.Println(target)

	}
	if !success {
		fmt.Println("Oops! your guess was not accurate", target)
	}
	result, err := calculatepaintneeded.Calculatepaintneeded(9.0, 7.0)
	fmt.Println(err)
	fmt.Printf("the paint needed in litres is: %0.3f\n", result)

	fmt.Printf("%12s | %s\n", "Product", "Amount in Cent")
	fmt.Printf("-------------------------------\n")
	fmt.Printf("%12s | %d\n", "Rice", 60)
	fmt.Printf("%12s | %d\n", "Gun", 90)
	fmt.Printf("%12s | %d\n", "Goat", 890)

	check := pass_fail.Checkgrade()
	fmt.Println(check)
	amount := 6
	fmt.Println(amount)
	fmt.Println(&amount)
	mynum := &amount
	fmt.Println(mynum)
	fmt.Println(*mynum)
	var myint float64
	mypos := &myint
	fmt.Println(reflect.TypeOf(&myint))
	fmt.Println(mypos)
	fmt.Println(reflect.TypeOf(&myint))
	fmt.Print("Enter temperature in Fahrenheit: ")
	fahrenheit, err := keyboard.Getfloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrenheit - 32) * 5 / 9
	// fmt.Println(celsius)
	fmt.Printf("%.5f celsius\n", celsius)

	greeting.Hello()
	greeting.Hi()
	deutsch.Deutsch()
	deutsch.GutenTag()
}
