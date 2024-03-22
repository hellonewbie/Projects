package main

import (
	"io"
	"log"
	"net"
)

func handle(src net.Conn) {
	//website:127.0.0.1:80
	dst, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		log.Fatalln("Unable to connect to our unreachable host")
	}
	defer dst.Close()
	//告诉joe你已经通过代理连上了目标网站
	if _, err := src.Write([]byte("Success connect proxy")); err != nil {
		log.Fatalln("Unable to write data")
	}
	go func() {
		//把客户端流量导到目标网站
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()
	//
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}

}

func main() {
	listener, err := net.Listen("tcp", ":30000")
	if err != nil {
		log.Fatalln("Unable to bind port")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go handle(conn)
	}
}
