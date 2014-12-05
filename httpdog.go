package httpdog

import (
	"net/http"
	"strings"
)

type HttpDog struct {
}

func NewHttpDog() *HttpDog {
	return &HttpDog{}
}

func (d *HttpDog) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ip := r.Header.Get("X-Real-IP")
	if !(strings.HasPrefix(ip, "172.16.") || strings.HasPrefix(ip, "192.168.")) {
		w.Write([]byte("Invalid client ip!\n"))
		return
	}
	next(w, r)
}
