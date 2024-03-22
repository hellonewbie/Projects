package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		//打印错误信息并导致程序以 panic 的方式终止执行
		log.Panicln(err)
	}
	// Print HTTP Status
	fmt.Println(resp.Status)

	// Read and display response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(body))
	resp.Body.Close()

	resp, err = http.Head("https://www.baidu.com")
	if err != nil {
		log.Panicln(err)
	}
	resp.Body.Close()
	fmt.Println(resp.Status)
	//在Go语言中，url.Values{}是一个用于表示URL查询参数的数据结构，它类似于一个字典或映射，可以用来存储键值对。
	form := url.Values{}
	//URL查询参数中添加一个键值对，其中键为"foo"，值为"bar"。这样做的效果是将键值对"foo=bar"添加到URL查询参数中
	//在这个例子中，我们向https://www.baidu.com/s?这个URL后面添加了查询参数wd=王者
	//第一个参数就是网站、第二个参数就是指定了请求体的数据类型
	//第三个参数是请求体的内容，使用strings.NewReader(form.Encode())来创建一个包含URL编码表单数据的io.Reader。
	//这里的form.Encode()是将url.Values类型的表单数据编码成URL编码的字符串，以便于发送到服务器。
	form.Add("wd", "王者")
	resp, err = http.Post(
		"https://www.baidu.com/s?",
		"application/x-www-form-urlencoded",
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		log.Panicln(err)
	}
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("Failed post")
	}
	fmt.Println(string(body))
	resp.Body.Close()
	fmt.Println(resp.Status)

	req, err := http.NewRequest("DELETE", "https://www.baidu.com", nil)
	if err != nil {
		log.Panicln(err)
	}
	var client http.Client
	resp, err = client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	resp.Body.Close()
	fmt.Println(resp.Status)

	req, err = http.NewRequest("PUT", "https://www.baidu.com", strings.NewReader(form.Encode()))

	if err != nil {
		log.Panicln(err)
	}
	resp, err = client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	resp.Body.Close()
	fmt.Println(resp.Status)

}
