package handlers

import (
	"fmt"
	"net/http"
)

func (a *App) ChangeHandler(w http.ResponseWriter, r *http.Request) {
	product := r.FormValue("product")
	transaction, err := a.Hyperledger.ProductInvoke(product)
	if err != nil {
		http.Error(w, "Cannot post data in blockchain", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, transaction)
}
