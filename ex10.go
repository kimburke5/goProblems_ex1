// Kimberly Burke
// G00269948
//https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go

package main 

import "fmt"

func main() { 
		
		//Prompts and takes user input
		input := ""
		fmt.Println("Please input a word") 
		fmt.Scanf("%s", &input)
		
        // Get Unicode code points. 
        n := 0
        rune := make([]rune, len(input))
        for _, r := range input { 
                rune[n] = r
                n++
        } 
        rune = rune[0:n]
        // Reverses user input
        for i := 0; i < n/2; i++ { 
                rune[i], rune[n-1-i] = rune[n-1-i], rune[i] 
        } 
        // Convert back to UTF-8. 
        output := string(rune)
        fmt.Println(output)
}
