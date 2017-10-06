// Kimberly Burke
// G00269948
//https://rosettacode.org/wiki/Sorting_algorithms/Merge_sort#Go
//https://stackoverflow.com/questions/16248241/concatenate-two-slices-in-go

package main
 
import "fmt"

//initializing variable
var a = []int{2, 1}
var b = []int{4, 3}
var s = make([]int, len(a)/2+1) // scratch space for merge step
 
func main() {
	//print original arrays
    fmt.Println("before:", a, b)
	//merging the arrays
    a := append([]int{2, 1}, []int{4, 3}...)
	//running the merged arrays through mergeSort function
    mergeSort(a)
    fmt.Println("after: ", a)
}

//mergeSort function to sort list starting from lowest
//and ending with the highest value
func mergeSort(a []int) {
    if len(a) < 2 {
        return
    }
    mid := len(a) / 2
    mergeSort(a[:mid])
    mergeSort(a[mid:])
    if a[mid-1] <= a[mid] {
        return
    }
    // merge step, with the copy-half optimization
    copy(s, a[:mid])
    l, r := 0, mid
    for i := 0; ; i++ {
        if s[l] <= a[r] {
            a[i] = s[l]
            l++
            if l == mid {
                break
            }
        } else {
            a[i] = a[r]
            r++
            if r == len(a) {
                copy(a[i+1:], s[l:mid])
                break
            }
        }
    }
    return
}