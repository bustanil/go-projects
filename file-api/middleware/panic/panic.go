package panic

import (
	"log"
	"net/http"

	"runtime/debug"
)

type PanicLogger struct {
}

func (m *PanicLogger) Intercept(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Panic caught %+v", r)
			log.Printf("Stacktrace: %+v", string(debug.Stack()))
		}
	}()
	next(w, r)
}
