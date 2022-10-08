package debug

import (
	"github.com/sirupsen/logrus"
	"net"
)

func debugNet(method string, params []string) {
	if method == "ParseIp" {
		parseIP := net.ParseIP(params[0])
		logrus.Info("parse ip is ", parseIP)
	}
}
