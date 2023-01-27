package utm

import (
	"bytes"
	"fmt"
	"net/http"
)

func (u *utmService) Delete(num string) error {
	uri := u.GetUtm4Query(`opt/out/` + num)
	if req, err := http.NewRequest(http.MethodDelete, uri, &bytes.Buffer{}); err != nil {
		return fmt.Errorf("utm.Delete %w", err)
	} else {
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("client.Do(req) %w", err)
		}
		defer resp.Body.Close()
	}
	return nil
}
