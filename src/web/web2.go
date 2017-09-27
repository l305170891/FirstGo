package main

import (
	"fmt"
	"io/ioutil"
)

// 数据
type page struct {
	title string
	body  []byte
}

//保存方法
func (p *page) save() error {
	filename := p.title + ".txt"
	return ioutil.WriteFile(filename, p.body, 0600)
}

//加载页面数据
func loadPage(title string) (*page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &page{body: body, title: title}, nil

}

func main() {

	p1 := &page{title: "testPage", body: []byte("this is a simple page.")}
	p1.save()

	p2, err := loadPage("testPage2")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(p2.body))
	}

}
