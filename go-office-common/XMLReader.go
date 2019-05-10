package go_office_common

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"strings"
)

type XMLReader struct {

}

func (r *XMLReader) GetMultiDomFromZip(zipFile string, xmlFiles []string) (contents [][]byte, err error) {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return nil,err
	}
	defer reader.Close()
	for _, file := range reader.File {
		if StringInArray(file.Name, xmlFiles) {
			rc, err := file.Open()
			if err != nil {
				return nil, err
			}
			defer rc.Close()
			b, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, err
			}
			contents = append(contents,[]byte(b))
		}
	}
	return contents, nil
}


func (r *XMLReader) GetDomFromZip(zipFile string, xmlFile string) (content []byte, err error) {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return nil,err
	}
	defer reader.Close()
	for _, file := range reader.File {
		if file.Name == xmlFile {
			rc, err := file.Open()
			if err != nil {
				return nil, err
			}
			defer rc.Close()
			b, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, err
			}
			//return string(b), nil
			return b, nil
		}
	}
	return nil, errors.New("No "+xmlFile+" File")
}

func (r *XMLReader) GetDomFromString(content []byte) (a interface{}) {
	err := xml.Unmarshal(content, &a)
	fmt.Println(a)
	fmt.Println(err)
	if err != nil {
		fmt.Println("test")
		fmt.Println(err.Error())
	}
	return a
}

func (r *XMLReader) GetNodes(contentNode string) (a interface{}) {
	err := xml.Unmarshal([]byte(contentNode), &a)
	fmt.Println(a)
	fmt.Println(err)
	if err != nil {
		fmt.Println("test")
		fmt.Println(err.Error())
	}
	return a
}

func (r *XMLReader) GetElements(content []byte)  {
	inputReader := strings.NewReader(string(content))
	decoder := xml.NewDecoder(inputReader)
	var t xml.Token
	var err error
	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
			// 处理元素开始（标签）
			case xml.StartElement:
				name := token.Name.Local
				fmt.Printf("Token name: %s\n", name)
				for _, attr := range token.Attr {
					attrName := attr.Name.Local
					attrValue := attr.Value
					fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
				}
			// 处理元素结束（标签）
			case xml.EndElement:
				fmt.Printf("Token of '%s' end\n", token.Name.Local)
			// 处理字符数据（这里就是元素的文本）
			case xml.CharData:
				content := string([]byte(token))
				fmt.Printf("This is the content: %v\n", content)
			default:
				//
		}
	}
}


func StringInArray(need string, needArray []string) bool {
	for _, v := range needArray {
		if need == v {
			return true
		}
	}
	return false
}