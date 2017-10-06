// Kimberly Burke
// G00269948
//https://www.socketloop.com/tutorials/golang-find-biggest-largest-number-in-array

package main

 import "fmt"

 func main() {
	 //initializing an array with 6 integers
         arr := []uint{
                 28, 33, 16,
                 7, 5, 88,
         }
		 // assume first value is the smallest in the array
         max := arr[0] 
		 min := arr[0]

         for _, value := range arr {
                 if value > max {
                         max = value // found another smaller value, replace previous value in max
                 }
				 if value < min {
                         min = value // found another smaller value, replace that value with min
                 }
         }
		 //prints the smallest and the largest integers in the array
         fmt.Println("The biggest/largest value is : ", max)
		 fmt.Println("The smallest/lowest value is : ", min)
 }