package powerpoint_struct

import "encoding/xml"

//docProps/custom.xml
type DocPropsCustom struct {
	XMLName		string		`xml:"Properties"`
	Xmlns		string		`xml:"xmlns,attr"`
	XmlnsVt		string		`xml:"xmlns:vt,attr"`
}

//docProps/core.xml
type DocPropsCore struct {	//for this content ,must use xml.Token, can not use Unmarshal
	XMLName			xml.Name	`xml:"cp:coreProperties"`
	XmlnsCp			string		`xml:"xmlns cp,attr"`
	XmlnsDc			string		`xml:"xmlns dc,attr"`
	XmlnsDcterms	string		`xml:"xmlns dcterms,attr"`
	XmlnsDcmitype	string		`xml:"xmlns dcmitype,attr"`
	XmlnsXsi		string		`xml:"xmlns xsi,attr"`

	DcCreator			xml.CharData		`xml:"dc:creator"`
	CpLastModifiedBy	string		`xml:"cp:lastModifiedBy"`
	DctermsCreated		DocPropsCoreDctermsCreated		`xml:"dcterms:created"`
	DctermsModified		DocPropsCoreDctermsModified		`xml:"dcterms:modified"`
	DcTitle				string		`xml:"dc:title"`
	DcDescription		string		`xml:"dc:description"`
	DcSubject			string		`xml:"dc:subject"`
	CpKeywords			string		`xml:"cp:keywords"`
	CpCategory			string		`xml:"cp:category"`
}
type DocPropsCoreDctermsCreated struct {
	XMLName			xml.Name	`xml:"dcterms:created"`
	XsiType			string		`xml:"xsi:type,attr"`
	DctermsCreated	string			`xml:",innerxml"`
}
type DocPropsCoreDctermsModified struct {
	XMLName			xml.Name	`xml:"dcterms:modified"`
	XsiType			string		`xml:"xsi:type,attr"`
	DctermsCreated	string			`xml:",innerxml"`
}

//docProps/app.xml
type DocPropsApp struct {
	XMLName		string		`xml:"Properties"`
	Xmlns		string		`xml:"xmlns,attr"`
	XmlnsVt		string		`xml:"xmlns:vt,attr"`
	Application	string			`xml:"Application"`
	Slides		int				`xml:"Slides"`
	ScaleCrop	bool			`xml:"ScaleCrop"`
	HeadingPairs	DocPropsAppHeadingPairs
	TitlesOfParts	DocPropsAppTitlesOfParts
	Company				string	`xml:"Company"`
	LinksUpToDate		bool	`xml:"LinksUpToDate"`
	SharedDoc			bool	`xml:"SharedDoc"`
	HyperlinksChanged	bool	`xml:"HyperlinksChanged"`
	AppVersion			string	`xml:"AppVersion"`
}

type DocPropsAppHeadingPairs struct {
	Size		int		`xml:"size,attr"`
	BaseType	string		`xml:"baseType,attr"`
	VtVariant	[]DocPropsAppHeadingPairsVtVariant	`xml:"vt:variant"`
}
type DocPropsAppTitlesOfParts struct {
	Size		int		`xml:"size,attr"`
	BaseType	string		`xml:"baseType,attr"`
	VtLpstr		string			`xml:"vt:lpstr"`
}
type DocPropsAppHeadingPairsVtVariant struct {
	VtLpstr		string			`xml:"vt:lpstr"`
	VtI4		int				`xml:"vt:i4"`
}