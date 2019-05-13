package powerpoint_reader

import (
	powerpoint_struct "GoOffice/go-office-powerpoint/struct"
	"archive/zip"
	"encoding/xml"
	"github.com/pkg/errors"
	"io/ioutil"
)

type PowerPointReader struct{
	ContentTypes 	ContentTypes
	Ppt				Ppt
	DocProps	DocProps
	Files		[] FileContent
}
type ContentTypes struct {
	ContentTypesXml		powerpoint_struct.ContentTypesXml
	ContentTypesXmlRels	powerpoint_struct.ContentTypesXmlRels
}
type DocProps struct {
	App			powerpoint_struct.DocPropsApp
	Core		powerpoint_struct.DocPropsCore
	Custom		powerpoint_struct.DocPropsCustom
}
type Ppt struct {
	Medias	[]interface{}
	Slides	[]interface{}
	Themes	[]interface{}
	Masters	[]interface{}
	Layouts	[]interface{}
	Presentation	interface{}
	PresentationRels	interface{}
	PresProps		interface{}
	TableStyles		interface{}
	ViewProps		interface{}
}
type FileContent struct {
	Name		string
	Content		[]byte
}

func (r *PowerPointReader) LoadPowerPoint(path string) (err error) {
	reader, err := zip.OpenReader(path)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()
		b, err := ioutil.ReadAll(rc)
		if err != nil {
			return err
		}
		fileContent := FileContent{file.Name, b}
		r.Files = append(r.Files, fileContent)
	}
	return nil
}

func ReadXmlStruct(content []byte, fileName string) ( result interface{}, err error) {
	if fileName == "[Content_Types].xml" {
		a := powerpoint_struct.ContentTypesXml{}
		result = &a
	} else if fileName == "_rels/.rels" {
		a := powerpoint_struct.ContentTypesXmlRels{}
		result = &a
	} else {
		return nil, errors.New("No struct to read file " + fileName)
	}
	err = xml.Unmarshal(content, result)
	if err != nil {
		return nil, err
	}
	//fmt.Println(result)
	return result, nil
}
