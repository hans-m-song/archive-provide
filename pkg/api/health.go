package api

import (
	"fmt"
	"net/http"
)

func healthHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "OK")
	rw.WriteHeader(200)
}
