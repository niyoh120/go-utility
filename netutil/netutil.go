package netutil

import (
	"net"
	"strings"

	"github.com/niyoh120/go-utility/debug/assert"
)

func JoinHostPort(host, port string) string {
	assert.NotZero(host)
	if port != "" {
		return net.JoinHostPort(host, port)
	}
	return host
}

func SplitHostPort(addr string) (host, port string) {
	var err error
	if strings.Contains(addr, ":") {
		host, port, err = net.SplitHostPort(addr)
		assert.Nil(err)
		return
	}
	return addr, ""
}
