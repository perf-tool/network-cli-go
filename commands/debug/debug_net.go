package debug

import (
	"github.com/sirupsen/logrus"
	"net"
)

func debugNet(method string, params []string) {
	switch method {
	case "ParseIp":
		parseIP := net.ParseIP(params[0])
		logrus.Info("parse ip is ", parseIP)
	case "LookupIp":
		lookupIp, err := net.LookupIP(params[0])
		if err != nil {
			panic(err)
		}
		logrus.Info("Lookup ip is ", lookupIp)
	}
}
