package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Status struct {
	Message string
	Status  string
}

func main() {
	res, err := http.Get(
		"http://localhost:8082/hello",
	)
	if err != nil {
		log.Fatalln(err)
	}

	var status Status
	//创建一个解码器，解码器含有缓冲，并将解码后的内容存到status结构体中
	if err := json.NewDecoder(res.Body).Decode(&status); err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	log.Printf("%s -> %s\n", status.Status, status.Message)
}
