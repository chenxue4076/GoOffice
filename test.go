package main

import (
	goOfficeCommon "GoOffice/go-office-common"
	"fmt"
	"strings"
)

func main() {
	xmlReader := goOfficeCommon.XMLReader{}
	files := []string{"[Content_Types].xml", "docProps/app.xml", "docProps/core.xml"}
	contents, err := xmlReader.GetMultiDomFromZip("/data/report/file_report/right.pptx", files)
	if err != nil {
		fmt.Println("main err")
		fmt.Println(err.Error())
	}
	for _, content := range contents {
		xmlReader.GetElements(content)
	}
}
func getDir(path string) string {
	return subString(path, 0, strings.LastIndex(path, "/"))
}

func subString(str string, start, end int) string {
	rs := []rune(str)
	length := len(rs)
	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < start || end > length {
		panic("end is wrong")
	}
	return string(rs[start:end])
}