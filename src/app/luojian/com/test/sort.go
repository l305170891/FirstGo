package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	//arr sort
	arr1 := [5]int{1, 3, 5, 6}
	sort.Ints(arr1[:])
	fmt.Println(arr1)

	//slice sort
	arr2 := []string{"aa", "111", "bb", "agb"}
	sort.Strings(arr2)
	fmt.Println(arr2)

}
