package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hans-m-song/archive-provide/pkg/provide"
)

type HandlerFunction func(rw http.ResponseWriter, r *http.Request)

func listNamesHandler(provider *provide.Provider, table string) HandlerFunction {
	return func(rw http.ResponseWriter, r *http.Request) {
		data, err := provider.ListNames(table)
		if err != nil {
			defaultServerErrorResponse(rw, err)
			return
		}

		content, err := json.Marshal(*data)
		if err != nil {
			defaultServerErrorResponse(rw, err)
			return
		}

		rw.WriteHeader(http.StatusOK)
		fmt.Fprintf(rw, "%s", content)
	}
}
