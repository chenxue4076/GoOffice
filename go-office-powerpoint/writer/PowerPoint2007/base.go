package powerpoint2007

import (
	powerpoint_struct "GoOffice/go-office-powerpoint/struct"
	"encoding/xml"
	"fmt"
)

type PowerPoint2007 struct {
	powerpoint_struct.PowerPoint
}

func SaveBuffFile(w *PowerPoint2007, contentStruct interface{}, fileName string)  {
	var fileContent powerpoint_struct.FileContent
	//将数据生成到临时文件列表中
	xmlStream, err := xml.MarshalIndent(contentStruct, "", "    ")
	if err != nil {
		fmt.Println(err.Error())
	}
	xmlStream = append([]byte(xml.Header), xmlStream ...)
	fileContent = powerpoint_struct.FileContent{fileName, xmlStream}
	w.Files = append(w.Files, fileContent)
}