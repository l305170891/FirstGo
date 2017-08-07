package std

import "fmt"

func VarRun() {
    fmt.Println("变量学习--开始")

    //int
    var i int
    i = 100
    fmt.Printf("i typeof int  ,  i = %d \n" , i)

    //float
    var f float32 = 3.1415
    fmt.Printf("f typeof float  ,  f = %f \n" , f)

    //string
    s := "string string"
    fmt.Printf("s typeof string  ,  s = %s \n" , s)

    fmt.Println("变量学习--结束")
}
