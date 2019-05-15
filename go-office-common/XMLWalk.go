package go_office_common

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

func XmlWalk(content []byte, v interface{}) (err error) {
	//inputReader := strings.NewReader(string(content))
	decoder := xml.NewDecoder(bytes.NewReader(content))
	var t xml.Token
	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {

		switch token := t.(type) {
			// 处理元素开始（标签）
			case xml.StartElement:
				fmt.Println("Start Element")
				fmt.Println(token)
				name := token.Name.Local
				fmt.Printf("Token name: %s\n", name)
				fmt.Printf("Token name space: %s\n", token.Name.Space)
				for _, attr := range token.Attr {
					attrName := attr.Name.Local
					attrValue := attr.Value
					fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
				}
			// 处理元素结束（标签）
			case xml.EndElement:
				fmt.Println("End Element")
				fmt.Println(token)
				fmt.Printf("Token of '%s' end\n", token.Name.Local)
			// 处理字符数据（这里就是元素的文本）
			case xml.CharData:
				fmt.Println("CharData")
				fmt.Println(string(token))
				content := string([]byte(token))
				fmt.Printf("This is the content: %v\n", content)
			default:
				fmt.Println("Default::::")
				fmt.Println(token)
				//
		}
	}
	return err
}