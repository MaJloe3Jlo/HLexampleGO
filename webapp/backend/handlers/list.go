package handlers

import (
	"fmt"
	"net/http"
)

func (a *App) AllHandler(w http.ResponseWriter, r *http.Request) {
	value, err := a.Hyperledger.ProductAll()
	if err != nil {
		http.Error(w, "Cannot get an value from blockchain", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(value))
}
