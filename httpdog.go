package httpdog

import (
	"fmt"
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
		w.Write([]byte(fmt.Sprintf("Invalid client ip(%s)\n", ip)))
		return
	}
	next(w, r)
}
