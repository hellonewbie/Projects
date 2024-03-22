package IO

import (
	"black-hat/IoAndParse/parse"
	"fmt"
	"log"
	"os"
	"strings"
)

// FooReader defines an io.Reader to read from stdin.
type FooReader struct{}

// Read reads data from stdin.
func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

func Input() {
	// Instantiate reader and writer.
	var (
		reader FooReader
	)

	// Create buffer to hold input/output.
	input := make([]byte, 4096)

	// Use reader to read input.
	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	fmt.Printf("Read %d bytes from stdin\n", s)
	// Use writer to write output.
	//这里要求传递的参数是字符串
	//这个地方卡了好久
	//问题可能出现在字符串的格式上。
	//因为你直接将读取的字节数据转换为字符串，
	//但是用户输入的数据格式可能不是严格的文本字符串
	//而是包含了换行符或其他非可见字符。
	//为了解决这个问题，你可以对读取的字节数据进行处理，去除可能的换行符等非可见字符，然后再将其转换为字符串。
	//可以使用 strings.TrimSpace 函数来去除字符串两端的空白字符，确保得到的字符串是干净的。
	//因为是以字节的形式读取的所以可能换行符什么也读进去了
	str := strings.TrimSpace(string(input[:s]))
	fmt.Println(str)
	parsePort, err := parse.Parse(str)
	if err != nil {
		log.Fatalln("Unable to parse input")
	}
	for index, port := range parsePort {
		ports := fmt.Sprintf("%d port:%d", index, port)
		fmt.Println(ports)
	}
}
