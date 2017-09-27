package std

import (
	"fmt"
	"reflect"
)

func ArrRun() {

	/*
	   1.数组是值类型，将一个数组赋值给另一个数组时将复制一份新的元素(切片是引用类型)
	   2.数组的长度不可改变


	*/

	fmt.Println("数组学习--开始")

	//未赋值
	var arr1 [4]int
	for _, v := range arr1 {
		fmt.Printf("arr1 is %d \n", v)
	}

	fmt.Printf("arr1 len is %d \n", len(arr1))
	fmt.Printf("arr1 cap is %d \n", cap(arr1))

	//赋值
	//var arr2 = [3] int { 1,2,3}
	var arr2 = [...]int{1, 2, 3}
	for i, v := range arr2 {
		fmt.Printf("arr2 index is %d, value is %d \n", i, v)
	}

	arr2[2] = 4
	for i, v := range arr2 {
		fmt.Printf("arr2 index is %d, value is %d \n", i, v)
	}

	fmt.Printf("arr2 len is %d \n", len(arr2))
	fmt.Printf("arr2 cap is %d \n", cap(arr2))

	//二维数组
	//var arr3  = [...][...] int{
	//    {1,2,3},
	//    {4,5,6},
	//}
	var arr3 [2][3]int
	arr3 = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	for i, v1 := range arr3 {
		for j, v2 := range v1 {
			fmt.Printf("arr3 index is %d  %d, value is %d \n", i, j, v2)
		}
	}

	fmt.Printf("arr3 len is %d \n", len(arr3))
	fmt.Printf("arr3 cap is %d \n", cap(arr3))

	fmt.Println("arr3 type is ", reflect.TypeOf(arr3))

	//指定值
	var arr4 = [5]int{2: 1, 3: 2, 4: 3}
	for i, v := range arr4 {
		fmt.Printf("arr4 index is %d, value is %d \n", i, v)
	}

	fmt.Printf("arr4 len is %d \n", len(arr4))
	fmt.Printf("arr4 cap is %d \n", cap(arr4))

	fmt.Println("数组学习--结束")
}
