package account

import (
	"net/http"
)

func GetAccountHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Account data"))
}
