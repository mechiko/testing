package utm

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path"
	"strings"

	"testing/internal/entity"
)

// https://www.onlinetool.io/xmltogo/
type A struct {
	XMLName xml.Name `xml:"A"`
	Text    string   `xml:",chardata"`
	URL     string   `xml:"url"`
	Sign    string   `xml:"sign"`
	Ver     string   `xml:"ver"`
}

// <?xml version="1.0" encoding="UTF-8" standalone="no"?>
// <A>
// <url>3fbf9613-ddc3-4a6e-aa6f-3459466c2aa5</url>
// <sign>895B92CAD115B57B02C7D12ADC488066D99B60549D57A737B0CAC18E5E3E1C72E6B8D414C763CB58A5E67DE7C8C2ECE908451C4AE6838479A42ABBA8179D0CE8</sign>
// <ver>2</ver>
// </A>
// insert into state_txt (id, name) VALUES (0, 'Запрос создан');
// insert into state_txt (id, name) VALUES (1, 'Запрос отправлен');
// insert into state_txt (id, name) VALUES (-1, 'Ошибка');
// insert into state_txt (id, name) VALUES (2, 'Получен ответ');
// insert into state_txt (id, name) VALUES (3, 'Получен ответ частично');

func (u *utmService) Xml2Utm(valxml string, uri string) (entity.PostXmlReturn, error) {
	defer u.app.GetRecovery().RecoverLog("Xml2Utm()")
	var err error
	var doc entity.PostXmlReturn

	finalUri := `http://` + path.Join(u.GetUtm4Send(), uri)
	request, err := u.newfileUploadRequest(valxml, finalUri)
	if err != nil {
		return doc, fmt.Errorf("newfileUploadRequest() %w", err)
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return doc, fmt.Errorf("client.Do(request) %w", err)
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			return doc, fmt.Errorf("body.ReadFrom(resp.Body) %w", err)
		}
		resp.Body.Close()
		if err := xml.Unmarshal(body.Bytes(), &doc); err != nil {
			errUtm := &entity.UTMOutError{Err: err, Doc: body.String()}
			return doc, errUtm
		}
	}
	return doc, nil
}

func (u *utmService) newfileUploadRequest(xml string, uri string) (*http.Request, error) {
	defer u.app.GetRecovery().RecoverLog("newfileUploadRequest()")
	r := strings.NewReader(xml)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("xml_file", "file.xml")
	if err != nil {
		return nil, fmt.Errorf("writer.CreateFormFile() %w", err)
	}
	_, err = io.Copy(part, r)
	if err != nil {
		return nil, fmt.Errorf("io.Copy(part, r) %w", err)
	}
	err = writer.Close()
	if err != nil {
		return nil, fmt.Errorf("writer.Close() %w", err)
	}
	req, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest(POST, uri, body) %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, nil
}
