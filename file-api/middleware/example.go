package middleware

import (
	"fmt"
	"net/http"
)

type HTTPMiddleware struct {
	Counter int
}

func (m *HTTPMiddleware) Intercept(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("HTTP middleware", m.Counter)
	next(w, r)
}
