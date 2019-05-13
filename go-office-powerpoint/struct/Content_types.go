package powerpoint_struct

import "encoding/xml"

//[Content_types].xml
type ContentTypesXml struct {
	XMLName		xml.Name		`xml:"Types"`
	Xmlns		xml.Attr		`xml:"xmlns,attr"`
	Default 	[]ContentTypesXmlDefault	`xml:"Default"`
	Override	[]ContentTypesXmlOverride	`xml:"Override"`
}
type ContentTypesXmlDefault struct {
	Extension	xml.Attr		`xml:"Extension,attr"`
	ContentType	xml.Attr		`xml:"ContentType,attr"`
}
type ContentTypesXmlOverride struct {
	PartName	xml.Attr		`xml:"PartName,attr"`
	ContentType	xml.Attr		`xml:"ContentType,attr"`
}

//_rels/.rels
type ContentTypesXmlRels struct {
	XMLName		xml.Name		`xml:"Relationships"`
	Xmlns		xml.Attr		`xml:"xmlns,attr"`
	Relationship 	[]ContentTypesXmlRelsRelationship	`xml:"Relationship"`
}
type ContentTypesXmlRelsRelationship struct {
	Id			xml.Attr		`xml:"Id,attr"`
	Type		xml.Attr		`xml:"Type,attr"`
	Target		xml.Attr		`xml:"Target,attr"`
}