package powerpoint_struct

import "encoding/xml"

//docProps/custom.xml
type DocPropsCustom struct {
	XMLName		xml.Name		`xml:"Properties"`
	Xmlns		xml.Attr		`xml:"xmlns,attr"`
	Xmlns_vt	xml.Attr		`xml:"xmlns:vt,attr"`
}

//docProps/core.xml
type DocPropsCore struct {
	XMLName			xml.Name		`xml:"cp:coreProperties"`
	XmlnsCp			xml.Attr		`xml:"xmlns:cp,attr"`
	XmlnsDc			xml.Attr		`xml:"xmlns:dc,attr"`
	XmlnsDcterms	xml.Attr		`xml:"xmlns:dcterms,attr"`
	XmlnsDcmitype	xml.Attr		`xml:"xmlns:dcmitype,attr"`
	XmlnsXsi		xml.Attr		`xml:"xmlns:xsi,attr"`

	DcCreator		xml.Attr		`xml:"dc:creator"`
}

//docProps/app.xml
type DocPropsApp struct {

}