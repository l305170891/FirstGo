package main

import (
    "net/http"
    "fmt"
    "io/ioutil"
    "html/template"
    "reflect"
)

const lenPath = len("/view/")


// 数据
type Page struct {
    Title string
    Body  []byte
}

//保存方法
func (p *Page) save() error {
    filename := "./src/web/views/data/" + p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

//加载页面数据
func loadPage(title string) (*Page, error) {
    filename := "./src/web/views/data/" + title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Body: body, Title: title}, nil

}
/* handler */
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[lenPath:]
    p,err := loadPage(title)

    if err != nil {
        p = &Page{Title: title}
    }

    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
    //title := r.URL.Path[lenPath:]
    //p,err := loadPage(title)
    //if err != nil{
    //    p = page{title:title}
    //}
    //
    //fmt.Fprintf(w,
    //    "<h1>Editing %s</h1>"+
    //    "<form action=\"/save/%s\" method=\"POST\" >"+
    //    "<textarea name=\"body\">%s</textarea><br>"+
    //    "<input type=\"submit\" value=\"Save\">"+
    //    "</form>",
    //    p.title, p.title, p.body)

    title := r.URL.Path[lenPath:]
    p,err := loadPage(title)
    if err != nil{
        p = &Page{Title:title}
    }

    t, _  := template.ParseFiles("./src/web/views/edit.html")
    //t.Execute(w,  map[string]string{"Title":"Title1111", "Body":"Body11111"})
    t.Execute(w, Struct2Map(*p))
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[lenPath:]
    body := []byte(r.FormValue("body"))
    page := Page{Title:title, Body:body}
    page.save()

    http.Redirect(w, r, "/edit/"+title, http.StatusFound)
}

func Struct2Map(obj interface{}) map[string]interface{} {
    t := reflect.TypeOf(obj)
    v := reflect.ValueOf(obj)

    var data = make(map[string]interface{})
    for i := 0; i < t.NumField(); i++ {
        data[t.Field(i).Name] = v.Field(i).Interface()
    }
    return data
}

func main() {

    //handler  可以定义一个入口， 然后再分发
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/edit/", editHandler)
    http.HandleFunc("/save/", saveHandler)

    //通过读取配置文件， 或者 读取flag  参数 设置端口
    http.ListenAndServe(":8080", nil)
}
