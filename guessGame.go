// package
package main

// import stuff
import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {

	// create a random target number
	randomNumber := rand.Intn(100)
	// fmt.Println("The random number is: ", randomNumber)

	// Get the number from the user
	fmt.Println("In this game you have to Guess the random number between 1 and 100")

	fmt.Println("Enter the number: ")

	tryCount := 10

	for {
		tryCount--
		reader := bufio.NewReader(os.Stdin)   // this sets up the reader
		input, err := reader.ReadString('\n') // this reads the input when the user clicks enters or goes to the new line. also the error is set up here

		// Here if there is an error its gonna throw one
		if err != nil {
			log.Fatal(err)
			continue
		}

		input = strings.TrimSpace(input)       // we take input from above and trim it
		userNumber, err := strconv.Atoi(input) // we take the trimmed input and parses it to an integer
		// Here if there is an error its gonna throw one
		if err != nil {
			log.Fatal(err)
			continue
		}

		if userNumber < 0 || userNumber > 100 {
			fmt.Println("Please put a number within the limit!!!")
			continue
		}

		if userNumber > randomNumber {
			fmt.Println("Oops. Your guess was HIGH")
			fmt.Println("You have", tryCount, "tries left")
		} else if userNumber < randomNumber {
			fmt.Println("Oops. Your guess was LOW")
			fmt.Println("You have", tryCount, "tries left")
		} else {
			fmt.Println("Good Job! You guessed it!")
			break
		}

		if tryCount <= 0 {
			fmt.Println("Sorry. You didnt guess my number")
			fmt.Println("My number was: ", randomNumber)
			break
		}

	}

}
