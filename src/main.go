/* the application is all about the user guessing a number by inputting it through the terminal and then the application generate a range of random number .....
now if the number guessed by the user is equal to the random number generated at the first time
*/

package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	MAX_TURN = 10
	GAME     = `Welcome to the First turn, lucky winner game, you are expected to pick a guess number and winner at the first round!`
)

func main() {
	// checking if the terminal has an arg
	if len(os.Args) == 1 {
		fmt.Println("A lucky number is required from the terminal")
		return
	}

	fmt.Println(GAME)
	fmt.Println()

	rand.NewSource(time.Now().UnixNano())
	arg, err := strconv.Atoi(os.Args[1])

	if err != nil {
		log.Fatalf("%v", err)
		return
	}

	if arg < 0 {
		fmt.Println("Select a non-negative lucky number")
		return
	}

	guess := 10
	var n int

	for turn := 0; turn < 1; turn++ {
		n = rand.Intn(guess + 1)

		if n == arg {
			fmt.Printf("WINNER, Guess at the first round")
			i := rand.Intn(2)

			switch i {
			case 0:
				fmt.Println("You the Winner at the lucky guess, thanks for participating ")
			case 1:
				fmt.Println("Lucky Winner")
			}
			return
		}

	}

	fmt.Printf("You lost the game, the generated number is %d\n", n)
}
