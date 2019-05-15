package main

import (
	powerpoint_reader "GoOffice/go-office-powerpoint/reader"
	powerpoint_writer "GoOffice/go-office-powerpoint/writer"
)

func main() {
	powerPoint := powerpoint_reader.PowerPointReader{}
	powerPoint.LoadPowerPoint("/data/report/file_report/right.pptx")
	/*for _,v := range powerPoint.Files {
		fmt.Println(v.Name)
	}*/
	w := powerpoint_writer.PowerPointWriter{}
	w.ContentTypes = powerPoint.ContentTypes
	w.DocProps = powerPoint.DocProps
	w.Ppts = powerPoint.Ppts

	w.Build()
	path := "/data/report/file_report/go_test.pptx"
	w.Write(path)
}
