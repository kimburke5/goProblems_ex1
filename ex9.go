// Kimberly Burke
// G00269948
//https://gist.github.com/sighmin/9173219
package main

//import packages
import (
    "fmt"
    "math"
)

//function for newton method
func Newt(x float64) float64 {
    if x == 0 { return 0 }
    z := 1.0
    for i := 0; i < int(x); i++ {
        z = z - ((math.Pow(z, 2) - x) / (2 * z))
    }
    return z
}

//funtion for math.sqrt
func Sqrt(x float64) float64 {
    return math.Sqrt(x)
}

func main() {
	//outputs the comparison between newtons method and math.sqrt
    times := 15
    for i := 0; i < times; i++ {
        sqrt := Sqrt(float64(i))
        newt := Newt(float64(i))
        fmt.Println(i, "squared:")
        fmt.Println("  Sqrt:", sqrt)
        fmt.Println("  Newt:", newt)
        fmt.Println("  Difference:", math.Abs(sqrt-newt))
    }
}