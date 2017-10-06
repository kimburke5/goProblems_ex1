// Kimberly Burke
// G00269948
// http://www.golangpro.com/2016/01/check-string-palindrome-go.html?m=1

package main

//importing packages needed
import (
	"fmt"
	"strings"
)
//program allows integers to be entered - no input validation
//however returns result palindrome if the order is the same 
//when read both directions

func main() {
	//initializing string variable
	var ip string
	//prompts and takes user input
	fmt.Println("Enter string:")
	fmt.Scanf("%s\n", &ip)
	//converts user entry to lowercase
	ip = strings.ToLower(ip)
	fmt.Println(isP(ip))
}
//Function to test if the string entered is a Palindrome
func isP(s string) string {
	mid := len(s) / 2
	last := len(s) - 1

	for i := 0; i < mid; i++ {
	
	//printed if string is not a palindrome
	if s[i] != s[last-i] {
		return "NO. It's not a Palimdrome."
  	}
}
 //printed if string is a paindrome
 return "YES! You've entered a Palindrome"
}