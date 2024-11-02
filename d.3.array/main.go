package main

import "fmt"

func main() {
	// ARRAYS - length cannot change
	var arr [3]int = [3]int{} // stores {0,0,0}
	arr2 := [3]int{1, 2, 3}
	arr2[0] = 0
	fmt.Println(arr, arr2)

	// mostly used
	// ****Slices (uses array under the hood) - length can change

	scores := []int{10, 12, 14}
	// append returns a new slice with the ele
	scores = append(scores, 16)
	fmt.Println(scores)

	names := []string{"nishant", "sarvesh", "dixit", "kumar"}

	names = names[1:3] // [sarvesh dixit] not includes the end limit
	// names = names[1:] // [sarvesh dixit kumar]
	// names = names[:2] // [nishant sarvesh dixit]
	fmt.Println(names)

}
