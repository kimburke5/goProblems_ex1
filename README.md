# goProblemSheet1
How to clone your repository in Github
1.	On GitHub, navigate to the main page of the repository.
2. Under the repository name, click Clone or download.
3. In the Clone with HTTPs section, click  to copy the clone URL for the repository.
4. Open Git Bash.
5. Change the current working directory to the location where you want the cloned directory to be made.
6. Type git clone, and then paste the URL you copied in Step 2.
7. Press Enter. Your local clone will be created.
//Resource: https://help.github.com/articles/cloning-a-repository/

Exercises
_________________________________________
Exercise 1. Kon’nichiwa, Sekai!

Write a program that prints Hello, world! in Japanese (using Japanese characters) to the screen.

To run this program you enter go build "ex1.go" in the terminal, then
go run "ex1.go" - this will see if the program will run, it will throw errors 
any if they are present.

This program when ran will print "こんにちは, 世界" to terminal which is Japanese for "Hello World"

To add to your repository do the following commands in the terminal

git add "ex1.go"

git commit -m "ex1"

git push

//Resources: https://golang.org/

_________________________________________
Exercise 2. Current Time

Write a program that prints the current time and date to the console.

To run this program you enter go build "ex2.go" in the terminal, then
go run "ex2.go" - this will see if the program will run, it will throw errors 
any if they are present.

Using the time.Now() function 
This program when ran will print the current date and time to terminal

To add to your repository do the following commands in the terminal

git add "ex2.go"

git commit -m "ex2"

git push

//Resources: https://tour.golang.org/welcome/4

_________________________________________
Exercise 3. FizzBuzz

Write a program that prints the numbers from 1 to 100, each on a new line, to the console, except for the following conditions. 
For multiples of three print Fizz instead of the number, and for multiples of five print Buzz. 
For numbers that are multiples of both three and five print FizzBuzz.

To run this program you enter go build "ex3.go" in the terminal, then
go run "ex3.go" - this will see if the program will run, it will throw errors 
any if they are present.

Using a for loop to run 100 times, this program when ran will print the correct fizz, buzz and fizzbuzz
in the correct required locations to terminal

To add to your repository do the following commands in the terminal

git add "ex3.go"

git commit -m "ex3"

git push

//Resources: https://golangcode.com/fizz-buzz-test-in-go/

_________________________________________
Exercise 4. Factorial digit sum

Write a function that calculates the sum of the digits of the factorial of a number.n! means n x (n − 1) ... x 3 x 2 x 1. 
For example, 10! = 10 x 9 x ... x 3 x 2 x 1 = 3628800, and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27. 
Then find the sum of the digits in the number 100!.

To run this program you enter go build "ex4.go" in the terminal, then
go run "ex4.go" - this will see if the program will run, it will throw errors 
any if they are present.

This program when ran will get the sum of the 158 digits from the factorial value of 100 and
print to the terminal. The factorial function calculates te factorial of a number,
then the add function gets the sum of the value. The main function calls both funtions 
to work together and the result is the sum of the digits from the factorial value.

To add to your repository do the following commands in the terminal

git add "ex4.go"

git commit -m "ex4"

git push

//Resources: https://stackoverflow.com/questions/46395819/get-sum-of-bigint-number-golang

_________________________________________
Exercise 5. Guessing Game

Write a guessing game where the user must guess a randomly generated number. 
After every guess the program tells the user whether their number was too high or too low. At the end, the number of 
tries should be printed. It counts only as one try if they input the same number multiple times consecutively.

To run this program you enter go build "ex5.go" in the terminal, then
go run "ex5.go" - this will see if the program will run, it will throw errors 
any if they are present.

This program is a guessing game, it first generates a random number using math/rand in the 
function xrand(). In the main function I assigned the random number to be between 1 - 50. 
When ran the program tells user to enter a number between 1-50, giving them 10 tries. The
program then enters a for loop that will execute a max of 10 times. The program then prompts
user to makes a guess. The guess is then tested with if statements and tells user after each 
guess if they are too high or too low. The program will stop once the 10 guesses are up
or they guess correctly.

To add to your repository do the following commands in the terminal

git add "ex5.go"

git commit -m "ex5"

git push

//Resources: http://golangcookbook.blogspot.ie/2012/11/guess-number-game-in-golang.html

_________________________________________
Exercise 6. Largest and Smallest in List

Write a function that returns the largest and smallest elements in a list.

To run this program you enter go build "ex6.go" in the terminal, then
go run "ex6.go" - this will see if the program will run, it will throw errors 
any if they are present.

This program organises numbers in an array to find  the largest and smallest elements.
First initializing an array with 6 integers, initialize min[] and max[] to 0.
Enters a for loop until all integers in the initial array have passed through.
Program prints the min and max values in the array to the terminal.

To add to your repository do the following commands in the terminal

git add "ex6.go"

git commit -m "ex6"

git push

//Resources: https://www.socketloop.com/tutorials/golang-find-biggest-largest-number-in-array

_________________________________________
Exercise 7. Palindrome Test

Write a function that tests whether a string is a palindrome. A palindrome is a string that reads the same in reverse, e.g. radar.

To run this program you enter go build "ex7.go" in the terminal, then
go run "ex7.go" - this will see if the program will run, it will throw errors 
any if they are present.

This program tests whether a string is a palindrome. This program prompts and takes user input.
Then converts user entry to lowercase. In func isP() it checks if the string reads as a palindrome funning it 
through a for loop. If it is not a palindrome hen this is printed to the terminal.
If it is, then it breaks from the for loop and prints that it is a palindrome to the terminal.

To add to your repository do the following commands in the terminal

git add "ex7.go"

git commit -m "ex7"

git push

//Resources: http://www.golangpro.com/2016/01/check-string-palindrome-go.html?m=1

_________________________________________
Exercise 8. Merge list and Sort

Write a function that merges two sorted lists into a new sorted list, e.g. merge([1,4,6], [2,3,5]) = [1,2,3,4,5,6].

To run this program you enter go build "ex8.go" in the terminal, then
go run "ex8.go" - this will see if the program will run, it will throw errors 
any if they are present.

I started with arry of integers []int{2, 1}, []int{4, 3}. The main function prints the two lists before they are manipulated.
Then the two lists are merged using append. The mergesort funtion is then called,
this orders the integers within the appended list. The result is then printed to the terminal

To add to your repository do the following commands in the terminal

git add "ex8.go"

git commit -m "ex8"

git push

//Resources: https://rosettacode.org/wiki/Sorting_algorithms/Merge_sort#Go
             https://stackoverflow.com/questions/16248241/concatenate-two-slices-in-go

_________________________________________
Exercise 9. Newton’s method for square roots

Write a function to calculate the square root of a number using Newton’s method. 
Newton’s method is to approximate the square root of a number x by picking a starting point z and then 
repeating the following operation.

z_next = z - ((z*z - x) / (2 * z))
This is repeated while the values of z_next and z are different. After each iteration the value of z is replaced with that of z_next.
Run a few tests to determine how close you are to the math.Sqrt value?
Hint: to declare and initialize a floating point value, give it floating-point syntax or use a conversion:
z := float64(1)
z := 1.0

To run this program you enter go build "ex9.go" in the terminal, then
go run "ex9.go" - this will see if the program will run, it will throw errors 
any if they are present.

There is a funtion called func Newt() this uses Newtons method formula to calculate square roots.
Then the program uses the math.Sqrt in the func Sqtr() to calculate square roots. 
The main function outputs the comparison between newtons method and math.sqrt.
 
To add to your repository do the following commands in the terminal

git add "ex9.go"

git commit -m "ex9"

git push

//Resources: https://gist.github.com/sighmin/9173219

_________________________________________
Exercise 10. Reverse String

Write a function to reverse a string in Go.

To run this program you enter go build "ex10.go" in the terminal, then
go run "ex10.go" - this will see if the program will run, it will throw errors 
any if they are present.

Program prompts and takes user input.
Gets Unicode code points of user input using a for loop. Then reverses user input in a for loop.
Then converts it back to UTF-8 and outputs reverse string to the terminal.
 
To add to your repository do the following commands in the terminal

git add "ex10.go"

git commit -m "ex10"

git push

//Resources: https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go

