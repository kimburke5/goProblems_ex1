

package main

import (
    "fmt"
    "math/big"
)

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
	
	fmt.Printf("The Sum of the digits from the factorial value of 100 is: ")
    fmt.Println(add(factoral(100)))

}