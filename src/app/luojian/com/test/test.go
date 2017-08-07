package main

import "fmt"

func main() {

	//var
	//变量声明必须使用，不使用会报错误
	var as string
	as = "abc"
	fmt.Println("hello, world", len(as))

	//default int
	var intdef int
	fmt.Println("intdef is :", intdef)

	//const
	//常量声明可以不使用，
	const const1 = "this is a const"
	const name1, name2 = "123123", 111
	fmt.Println(len(const1))

	//条件语句
	// if else
	const inta int = 2
	if inta <= 1 {
		fmt.Println("inta <= 1")
	} else if inta == 2 {
		fmt.Println("inta == 2")
	} else {
		fmt.Println("inta >2")
	}

	//switch 语句不需要在 case 后边加上 break

	// for
	var b int = 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	/* for 循环 */
	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值为: %d\n", a)
	}

	for a < b {
		a++
		fmt.Printf("a 的值为: %d\n", a)
	}

	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

	var switchInt = 1
	switch {
	case true:
		fmt.Println(1)
	case switchInt == 1:
		fmt.Println(2)

	}

	fmt.Println(sum(1, 2, 3, 4))

	fmt.Println(sum([]int{1, 2, 3, 4}...))

}

func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}
