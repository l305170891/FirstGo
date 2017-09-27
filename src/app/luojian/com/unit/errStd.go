package std

import (
	"errors"
	"fmt"
	"math"
)

func ErrRun() {
	fmt.Println("错误处理学习--开始")

	//  correct
	f, err := sqrt(4.0)
	if err == nil {
		fmt.Println("4 sqrt is ", f)
	} else {
		fmt.Println("error is ", err)
	}

	//  error
	f2, err2 := sqrt(-4.0)
	if err2 == nil {
		fmt.Println("-4 sqrt is ", f2)
	} else {
		fmt.Println("error is ", err2)
	}

	fmt.Println("错误处理学习--结束")
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("必须大于 0 ")
	}
	return math.Sqrt(f), nil
}
