package utm

import (
	"path"
	"strings"

	"testing/internal/entity"
)

type UtmService interface {
	Xml2Utm(valxml string, uri string) (entity.PostXmlReturn, error)
	GetApiUri() string
	Ping() bool
}

type utmService struct {
	app    entity.App
	utmUri string
	layout string
}

// создаем пустышку для проверки всех методов интерфейса
var _ UtmService = &utmService{}

func NewUtmService(a entity.App) (UtmService, error) {
	s := &utmService{
		app:    a,
		layout: a.GetConfiguration().TimeLayout,
	}
	cfg := a.GetConfiguration()
	s.utmUri = `http://` + cfg.UtmHost + ":" + cfg.UtmPort
	if s.utmUri == "" {
		s.utmUri = `http://localhost:8080`
	}

	return s, nil
}

func (u *utmService) GetApiUri() string {
	return path.Join(u.GetUtmUri(), "/opt/in/")
}

func (u *utmService) GetUtmUri() string {
	cfg := u.app.GetConfiguration()
	retUri := cfg.UtmHost + ":" + cfg.UtmPort
	if retUri == "" {
		retUri = "localhost:8080"
	}
	if !strings.HasPrefix(retUri, `http://`) {
		retUri = `http://` + retUri
		u.app.DebugLog().Str("API", retUri).Send()
	}
	return retUri
}

func (u *utmService) GetUtmAddr() string {
	cfg := u.app.GetConfiguration()
	retUri := cfg.UtmHost + ":" + cfg.UtmPort
	if retUri == "" {
		retUri = "localhost:8080"
	}
	return retUri
}

func (u *utmService) GetUtm4Send() string {
	cfg := u.app.GetConfiguration()
	retUri := cfg.UtmHost + ":" + cfg.UtmPort
	if retUri == "" {
		retUri = "localhost:8080"
	}
	retUri = path.Join(retUri, `opt/in`)
	// if !strings.HasPrefix(retUri, `http://`) {
	// 	retUri = `http://` + retUri
	// 	application.Get().DebugLog().Str("API", retUri).Send()
	// }
	return retUri
}

func (u *utmService) GetUtm4Query(url string) string {
	cfg := u.app.GetConfiguration()
	retUri := cfg.UtmHost + ":" + cfg.UtmPort
	if retUri == "" {
		retUri = "localhost:8080"
	}
	retUri = path.Join(retUri, url)
	if !strings.HasPrefix(retUri, `http://`) {
		retUri = `http://` + retUri
	}
	return retUri
}
