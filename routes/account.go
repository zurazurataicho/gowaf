package route

import (
	"net/http"
	handler "zura.org/gowaf/internal/account/presentation"
)

func AccountRouter(mux *http.ServeMux) {
	mux.HandleFunc("GET /account", handler.GetAccountHandler)
}
