package powerpoint2007

import (
	powerpoint_struct "GoOffice/go-office-powerpoint/struct"
	"strconv"
)

func (w *PowerPoint2007) WriteContentTypes()  {
	var contentTypesXml powerpoint_struct.ContentTypesXml
	// Types
	contentTypesXml.XMLName = "Types"
	contentTypesXml.Xmlns = "http://schemas.openxmlformats.org/package/2006/content-types"

	var xmlDefault []powerpoint_struct.ContentTypesXmlDefault
	// Rels
	xmlDefault = append(xmlDefault, writeDefaultContentType("rels", "application/vnd.openxmlformats-package.relationships+xml"))
	// XML
	xmlDefault = append(xmlDefault, writeDefaultContentType("xml", "application/xml"))

	var xmlOverride []powerpoint_struct.ContentTypesXmlOverride
	// Presentation
	xmlOverride = append(xmlOverride, writeOverrideContentType("/ppt/presentation.xml", "application/vnd.openxmlformats-officedocument.presentationml.presentation.main+xml"))
	// PptProps
	xmlOverride = append(xmlOverride, writeOverrideContentType("/ppt/presProps.xml", "application/vnd.openxmlformats-officedocument.presentationml.presProps+xml"))
	xmlOverride = append(xmlOverride, writeOverrideContentType("/ppt/tableStyles.xml","application/vnd.openxmlformats-officedocument.presentationml.tableStyles+xml"))
	xmlOverride = append(xmlOverride, writeOverrideContentType("/ppt/viewProps.xml","application/vnd.openxmlformats-officedocument.presentationml.viewProps+xml"))
	// DocProps
	xmlOverride = append(xmlOverride, writeOverrideContentType("/docProps/app.xml","application/vnd.openxmlformats-officedocument.extended-properties+xml"))
	xmlOverride = append(xmlOverride, writeOverrideContentType("/docProps/core.xml","application/vnd.openxmlformats-package.core-properties+xml"))
	xmlOverride = append(xmlOverride, writeOverrideContentType("/docProps/custom.xml","application/vnd.openxmlformats-officedocument.custom-properties+xml"))
	// Slide masters
	//sldLayoutNr := 0
	//sldLayoutId := time.Now().Unix() + 689016272	// requires minimum value of 2 147 483 648
	for im, _ := range w.PowerPoint.Ppts.Masters {
		xmlOverride = append(xmlOverride, writeOverrideContentType("/ppt/slideMasters/slideMaster"+ strconv.Itoa(im + 1) +".xml","application/vnd.openxmlformats-officedocument.presentationml.slideMaster+xml"))
		xmlOverride = append(xmlOverride, writeOverrideContentType("/ppt/theme/theme"+ strconv.Itoa(im + 1) +".xml","application/vnd.openxmlformats-officedocument.theme+xml"))
	}
	// Slides
	for is, _ := range w.PowerPoint.Ppts.Slides {
		xmlOverride = append(xmlOverride, writeOverrideContentType("/ppt/slides/slides"+ strconv.Itoa(is + 1) +".xml","application/vnd.openxmlformats-officedocument.presentationml.slide+xml"))
	}

	// Add media content-types
	xmlDefault = append(xmlDefault, writeDefaultContentType("gif", "image/gif"))
	xmlDefault = append(xmlDefault, writeDefaultContentType("jpg", "image/jpeg"))
	xmlDefault = append(xmlDefault, writeDefaultContentType("jpeg", "image/jpeg"))
	xmlDefault = append(xmlDefault, writeDefaultContentType("png", "image/png"))
	xmlDefault = append(xmlDefault, writeDefaultContentType("xlsx", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"))

	contentTypesXml.Default = xmlDefault
	contentTypesXml.Override = xmlOverride
	//save file template content
	SaveBuffFile(w, contentTypesXml, "[Content_Types].xml")
}

func (w *PowerPoint2007) WriteRelsRels() {
	var contentTypesXmlRels powerpoint_struct.ContentTypesXmlRels
	// Relationships
	contentTypesXmlRels.XMLName = "Relationships"
	contentTypesXmlRels.Xmlns = "http://schemas.openxmlformats.org/package/2006/relationships"

	var relationship []powerpoint_struct.ContentTypesXmlRelsRelationship
	relationship = append(relationship, writeRelsRelationship("rId1", "http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument", "ppt/presentation.xml"))
	relationship = append(relationship, writeRelsRelationship("rId2", "http://schemas.openxmlformats.org/package/2006/relationships/metadata/core-properties", "docProps/core.xml"))
	relationship = append(relationship, writeRelsRelationship("rId3", "http://schemas.openxmlformats.org/officeDocument/2006/relationships/extended-properties", "docProps/app.xml"))
	relationship = append(relationship, writeRelsRelationship("rId4", "http://schemas.openxmlformats.org/officeDocument/2006/relationships/custom-properties", "docProps/custom.xml"))

	contentTypesXmlRels.Relationship = relationship
	//save file template content
	SaveBuffFile(w, contentTypesXmlRels, "_rels/.rels")

}

func writeDefaultContentType(extention, contentType string) (defaultContent powerpoint_struct.ContentTypesXmlDefault) {
	defaultContent = powerpoint_struct.ContentTypesXmlDefault{extention,contentType}
	return defaultContent
}
func writeOverrideContentType(partName, contentType string) (overrideContent powerpoint_struct.ContentTypesXmlOverride) {
	overrideContent = powerpoint_struct.ContentTypesXmlOverride{partName,contentType}
	return overrideContent
}

func writeRelsRelationship(relationId, relationType, target string) (relationship powerpoint_struct.ContentTypesXmlRelsRelationship) {
	relationship = powerpoint_struct.ContentTypesXmlRelsRelationship{relationId, relationType, target}
	return relationship
}