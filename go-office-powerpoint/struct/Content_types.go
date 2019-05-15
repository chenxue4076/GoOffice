package powerpoint_struct

//[Content_types].xml
type ContentTypesXml struct {
	XMLName		string		`xml:"Types"`
	Xmlns		string		`xml:"xmlns,attr"`
	Default 	[]ContentTypesXmlDefault	`xml:"Default"`
	Override	[]ContentTypesXmlOverride	`xml:"Override"`
	//Media 		[]ContentTypesXmlDefault	`xml:"Default"`
}
type ContentTypesXmlDefault struct {
	Extension	string		`xml:"Extension,attr"`
	ContentType	string		`xml:"ContentType,attr"`
}
type ContentTypesXmlOverride struct {
	PartName	string		`xml:"PartName,attr"`
	ContentType	string		`xml:"ContentType,attr"`
}

//_rels/.rels
type ContentTypesXmlRels struct {
	XMLName		string		`xml:"Relationships"`
	Xmlns		string		`xml:"xmlns,attr"`
	Relationship 	[]ContentTypesXmlRelsRelationship	`xml:"Relationship"`
}
type ContentTypesXmlRelsRelationship struct {
	Id			string		`xml:"Id,attr"`
	Type		string		`xml:"Type,attr"`
	Target		string		`xml:"Target,attr"`
}