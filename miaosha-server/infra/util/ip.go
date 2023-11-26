package util

import (
	"net"
	"net/http"
	"strings"
)

func GetIP(r *http.Request) (ip string) {
	// 尝试从 X-Forwarded-For 中获取
	xForwardedFor := r.Header.Get(`X-Forwarded-For`)
	ip = strings.TrimSpace(strings.Split(xForwardedFor, `,`)[0])
	if ip == `` {
		// 尝试从 X-Real-Ip 中获取
		ip = strings.TrimSpace(r.Header.Get(`X-Real-Ip`))
		if ip == `` {
			// 直接从 Remote Addr 中获取
			_ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
			if err != nil {
				panic(err)
			} else {
				ip = _ip
			}
		}
	}
	return
}
