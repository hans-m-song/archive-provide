package api

import "net/http"

func headerJsonContent(rw http.ResponseWriter) {
	rw.Header().Set("Content-Type", "application/json")
}
