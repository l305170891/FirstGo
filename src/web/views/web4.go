package main

import (
    "net/http"
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

const lenPath = len("/view/")

/* handler */
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[lenPath:]

    db, err := sql.Open("mysql", "root:@/code2015?charset=utf8")
    checkErr(err)

    rows, err := db.Query("SELECT * FROM abstract")
    checkErr(err)

    //字典类型
    //构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
    columns, _ := rows.Columns()
    scanArgs := make([]interface{}, len(columns))
    values  := make([]interface{}, len(columns))
    for i := range values {
        scanArgs[i] = &values[i]
    }

    for rows.Next() {
        //将行数据保存到record字典
        err = rows.Scan(scanArgs...)
        record := make(map[string]string)
        for i, col := range values {
            if col != nil {
                record[columns[i]] = string(col.([]byte))
            }
        }
        fmt.Fprintln(w, record)
    }

    fmt.Fprintf(w, "<h1>%s</h1>", title)
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {

    //handler  可以定义一个入口， 然后再分发
    http.HandleFunc("/view/", viewHandler)


    //通过读取配置文件， 或者 读取flag  参数 设置端口
    http.ListenAndServe(":8081", nil)
}
