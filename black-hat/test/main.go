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
	form := url.Values{}
	form.Add("wd", "王者")
	req, err := http.NewRequest("POST", "https://www.baidu.com/s", strings.NewReader(form.Encode()))
	if err != nil {
		log.Fatal("POST FAILED SEND:", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("POST SEND FAILED:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("FAILED TO READ RESPONSE BODY:", err)
	}
	fmt.Println(string(body))
}
