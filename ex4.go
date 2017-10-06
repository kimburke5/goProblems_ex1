//Kimberly Burke
//G00269948
//https://stackoverflow.com/questions/46395819/get-sum-of-bigint-number-golang

package main

import (
    "fmt"
    "math/big"
)

//factorial function to get factorial value of a num
func factoral(n uint64) (r *big.Int) {

    one, bn := big.NewInt(1), new(big.Int).SetUint64(n)

    r = big.NewInt(1)
    if bn.Cmp(one) <= 0 {
        return
    }
    for i := big.NewInt(2); i.Cmp(bn) <= 0; i.Add(i, one) {
        r.Mul(r, i)
    }
    return
}

//gets sum of the digits from the factorial value
func add(number *big.Int) *big.Int {
    ten := big.NewInt(10)
    sum := big.NewInt(0)
    mod := big.NewInt(0)
    for ten.Cmp(number)<0 {
      sum.Add(sum, mod.Mod(number,ten))
      number.Div(number,ten)
    }
    sum.Add(sum,number)
  return sum
}
func main() {
	//prints results
	fmt.Printf("The Sum of the digits from the factorial value of 100 is: ")
    fmt.Println(add(factoral(100)))

}