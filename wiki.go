package main
import (
	"fmt"
	"io/ioutil"
	"os"
)

//页面数据结构
type page struct {
	title string
	body []byte
}

//页面的保存方法
func (p *page) save() os.Error {
	filename := p.title + ".txt"
	return ioutil.WriteFile(filename, p.body, 0600)
}

//加载页面方法
func loadPage(title string) *page {
	filename := title + ".txt"
	body, _ := ioutil.ReadFile(filename)
	return &page{title: title, body: body}
}

func loadPage(title string) (*page, os.Error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &page{title: title, body: body}, nil
}

func main() {
	p1 := &page{title: "TestPage", body: []byte("This is a sample page.")}
	p1.save()
	p2, _ := loadPage("TestPageTwo")
	fmt.println(string(p2.body))
}