// Kimberly Burke
// G00269948
//http://golangcookbook.blogspot.ie/2012/11/guess-number-game-in-golang.html

package main
import(
    "fmt"
    "math/rand"
    "time"
)

//this generates random number between given range
func xrand(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func main() {
	//initializing variables
    myrand := xrand(1, 50)
    guessTaken := 0
    var guess int

    fmt.Printf(" I am thinking of a number between 1 and 50. You Have 10 Tries! \n")
    
    //for loop - allowing user to only hav 10 tries
    for guessTaken < 10 {
        fmt.Println("Take a guess...")
        fmt.Scanf("%d", &guess)
		fmt.Scanf("%d")
        guessTaken++

        if guess < myrand {//if guess is lower then this is printed
            fmt.Println("Your guess is too low.")
        }
        if guess > myrand {//if guess is higher then this is printed
            fmt.Println("Your guess is too high.")
        }
        if guess == myrand {//if guess is correcet then for loop stops
            break
        }
    }
    if guess == myrand {//if guess is correct then this is printed
        fmt.Printf("Good job! You guessed my number in %d tries\n", guessTaken)
    } else {//If user has not guessed in the allowed 10 tries
        fmt.Printf("Sorry no guesses left! The number I had in mind was %d\n", myrand)
    }
}