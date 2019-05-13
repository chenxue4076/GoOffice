package main

import (
	powerpoint_reader "GoOffice/go-office-powerpoint/reader"
	"fmt"
)

func main() {
	powerPoint := powerpoint_reader.PowerPointReader{}
	powerPoint.LoadPowerPoint("/data/report/file_report/right.pptx")
	for _,v := range powerPoint.Files {
		fmt.Println(v.Name)
	}
}
