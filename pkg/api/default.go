package api

import (
	"fmt"
	"net/http"
)

func defaultServerErrorResponse(rw http.ResponseWriter, err error) {
	rw.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(rw, err)
}

func defaultNotFoundHandler(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusAccepted)
}
