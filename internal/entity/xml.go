package entity

type XmlService interface {
	SetFsrarId(fsrarid string)
	SetRequestId(requestid int64)
	SetXml(xml string)
}
