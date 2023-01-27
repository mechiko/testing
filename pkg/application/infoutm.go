package application

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"testing/internal/entity"
)

var myClient = &http.Client{Timeout: 2 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func (a *applicationType) GetUtmInfo() (*entity.UTMInfo, error) {
	cfg := a.configuration
	apiUrl := `http://` + cfg.UtmHost + ":" + cfg.UtmPort + "/api/info/list"
	utmInfo := &entity.UTMInfo{
		OwnerId: "",
	}

	if err := getJson(apiUrl, utmInfo); err != nil {
		a.GetConfig().Set("application.disconnected", true, true)
		a.GetConfig().Set("application.scanutm", false, true)
		return utmInfo, fmt.Errorf("getJson() %w", err)
	}
	a.GetConfig().Set("application.disconnected", false, true)
	return utmInfo, nil
}
