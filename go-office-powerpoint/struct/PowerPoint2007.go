package powerpoint_struct

type PowerPoint struct{
	ContentTypes 	ContentTypes
	Ppts			Ppts
	DocProps	DocProps
	Files		[] FileContent
}
type ContentTypes struct {
	ContentTypesXml		ContentTypesXml
	ContentTypesXmlRels	ContentTypesXmlRels
}
type DocProps struct {
	App			DocPropsApp
	Core		DocPropsCore
	Custom		DocPropsCustom
}
type Ppts struct {
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