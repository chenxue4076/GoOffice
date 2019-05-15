package powerpoint_writer

import (
	powerpoint2007 "GoOffice/go-office-powerpoint/writer/PowerPoint2007"
	"archive/zip"
	"bytes"
	"log"
	"os"
)

type PowerPointWriter struct{
	powerpoint2007.PowerPoint2007
}

func (w *PowerPointWriter) Build() {
	w.WriteContentTypes()
	w.WriteRelsRels()
}

func (w *PowerPointWriter) Write(path string) {
	//build a buffer
	buf := new(bytes.Buffer)
	//create zip file
	zipFile := zip.NewWriter(buf)

	for _, file := range w.Files {
		f, err := zipFile.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_,err = f.Write(file.Content)
		if err != nil {
			log.Fatal(err)
		}
	}
	//close zip file
	err := zipFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	toFile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	_, err = buf.WriteTo(toFile)
	if err != nil {
		log.Fatal(err)
	}
}