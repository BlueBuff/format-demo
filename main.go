package main

import (
	"hdg.com/tools/src/common-util"
	"fmt"
	"strings"
	"bytes"
	"io/ioutil"
)

func main() {
	reader := common_util.NewBufferFileReader("./resources/origin.data")
	data, err := reader.Read()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	var array = strings.Split(string(data), "\n")
	var b bytes.Buffer
	for _, item := range array {
		var temp = strings.Replace(item, "{", "[", -1)
		temp = strings.Replace(temp, "}", "]", -1)
		temp = strings.Replace(temp, "'", "\"", -1)
		b.WriteString(temp)
		b.WriteString("\n")
	}
	fmt.Println(b.String())
	WriteWithIoutil("./resources/content.out",b.String())
}

func WriteWithIoutil(name, content string) {
	data := []byte(content)
	if ioutil.WriteFile(name, data, 0644) == nil {
		fmt.Println("写入文件成功:", content)
	}
}
