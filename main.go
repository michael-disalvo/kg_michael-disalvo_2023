package main

import (
	"fmt"
	"os"
	"unicode"
)

//holds the name for each digit
var number_names = [10]string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}

/* This progam will take command line arguments that are numbers, like 23 or 7, and convert them to the form TwoThree and Seven respectively.
*  Any argument that is not a numebr will not be printed back out.
 */
func main() {
	//get command line arguments and store in args
	args := os.Args[1:]
	var output string = ""

	for i, number := range args {
		word_form, success := process_number(number)
		//only if the number was sucessfully processed should we add it to output
		if success {
			output += word_form
		}
		//if we are not at the last arg, add a comma afterwards
		if i != len(args)-1 && success {
			output += ","
		}
	}

	fmt.Println(output)
}

//takes in a number in form of string and returns the string representation of that number ("5" -> "Five")
func process_number(digit_form string) (string, bool) {
	var word_form string = ""
	//iterate over the characters of the input string
	for _, c := range digit_form {
		if unicode.IsDigit(c) {
			//to get int value from rune c subtract the value of '0'
			word_form += number_names[c-'0']
		} else {
			// raise an error
			return "ERROR", false
		}
	}
	return word_form, true
}
