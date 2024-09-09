package middleware

import "net/http"

type Middleware interface {
	Intercept(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}
