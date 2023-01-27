package utm

import (
	"net"
	"time"
)

func (u *utmService) Ping() bool {
	timeout := 1

	target := u.GetUtmAddr()
	conn, err := net.DialTimeout("tcp", target, time.Duration(timeout)*time.Second)
	if err != nil {
		u.app.GetConfig().Set("application.disconnected", true, true)
		return false
	}
	defer conn.Close()
	if u.app.GetConfiguration().Application.Disconnected {
		if err := u.app.InitUtm(); err != nil {
			u.app.ErrorLog().AnErr("Ping()", err).Send()
		}
	}
	return true
}
