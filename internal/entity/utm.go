package entity

import (
	"encoding/xml"
)

type Db struct {
	CreateDate string `json:"createDate"`
	OwnerId    string `json:"ownerId"`
}

type RSA struct {
	CertType   string `json:"certType"`
	StartDate  string `json:"startDate"`
	ExpireDate string `json:"expireDate"`
	IsValid    string `json:"isValid"`
	Issuer     string `json:"issuer"`
}

type UTMInfo struct {
	Version   string `json:"Version"`
	Contour   string `json:"Contour"`
	RsaError  string `json:"RsaError"`
	CheckInfo string `json:"CheckInfo"`
	OwnerId   string `json:"OwnerId"`
	Db        Db     `json:"db"`
	Rsa       RSA    `json:"rsa"`
	Gost      RSA    `json:"gost"`
	License   bool   `json:"license"`
}

type PostXmlReturn struct {
	XMLName xml.Name `xml:"A"`
	Text    string   `xml:",chardata"`
	URL     string   `xml:"url"`
	Sign    string   `xml:"sign"`
	Ver     string   `xml:"ver"`
}

type TicketItem struct {
	Id          int64
	RequestId   int64
	TicketDate  string
	DocType     string
	TransportId string
	RegId       string
	OpResult    string
	Oname       string
	Oresult     string
	Odate       string
	Ocomment    string
	Result      string
	Conclusion  string
	Cdate       string
	Comments    string
	XmlRaw      string
}
type TicketList struct {
	Total int64        `json:"total"`
	Rows  []TicketItem `json:"rows"`
}
