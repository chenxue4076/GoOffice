package powerpoint_reader

import (
	powerpoint_struct "GoOffice/go-office-powerpoint/struct"
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
)

type PowerPointReader struct{
	powerpoint_struct.PowerPoint
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

		fileContent := powerpoint_struct.FileContent{file.Name, b}
		r.Files = append(r.Files, fileContent)
	}
	r.LoadPowerPointContents()
	return nil
}

func (r *PowerPointReader) LoadPowerPointContents()  {
	for _, file := range r.Files {
		if file.Name == "[Content_Types].xml" {
			 err := xml.Unmarshal(file.Content, &r.ContentTypes.ContentTypesXml)
			 if err != nil {
			 	log.Fatal(err)
			 }
		} else if file.Name == "_rels/.rels" {
			err := xml.Unmarshal(file.Content, &r.ContentTypes.ContentTypesXmlRels)
			if err != nil {
				log.Fatal(err)
			}
		} else if file.Name == "docProps/custom.xml" {
			err := xml.Unmarshal(file.Content, &r.DocProps.Custom)
			if err != nil {
				log.Fatal(err)
			}
		} /*else if file.Name == "docProps/app.xml" {
			err := xml.Unmarshal(file.Content, &r.DocProps.App)
			if err != nil {
				log.Fatal(err)
			}
		} else if file.Name == "docProps/core.xml" {
			fmt.Println(string(file.Content))
			err := xml.Unmarshal(file.Content, &r.DocProps.Core)
			if err != nil {
				log.Fatal(err)
			}
			err := go_office_common.XmlWalk(file.Content, &r.DocProps.Core)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(r.DocProps.Core)
		} else if file.Name == "ppt/presentation.xml" {
			err := xml.Unmarshal(file.Content, &r.Ppts.Presentation)
			if err != nil {
				log.Fatal(err)
			}
		} else if file.Name == "ppt/_rels/presentation.xml.rels" {
			err := xml.Unmarshal(file.Content, &r.Ppts.PresentationRels)
			if err != nil {
				log.Fatal(err)
			}
		} else if file.Name == "ppt/viewProps.xml" {
			err := xml.Unmarshal(file.Content, &r.Ppts.ViewProps)
			if err != nil {
				log.Fatal(err)
			}
		} else if file.Name == "ppt/tableStyles.xml" {
			err := xml.Unmarshal(file.Content, &r.Ppts.TableStyles)
			if err != nil {
				log.Fatal(err)
			}
		} else if file.Name == "ppt/presProps.xml" {
			err := xml.Unmarshal(file.Content, &r.Ppts.PresProps)
			if err != nil {
				log.Fatal(err)
			}
		}*/
	}
	xmlStream, err := xml.MarshalIndent(r.DocProps.Custom, "", "    ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(xmlStream))
}

//end file
