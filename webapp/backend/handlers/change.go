package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MaJloe3Jlo/HLexampleGO/cc/model"
)

func (a *App) ChangeHandler(w http.ResponseWriter, r *http.Request) {
	product := &model.Product{}

	if err := json.NewDecoder(r.Body).Decode(product); err != nil {
		fmt.Fprint(w, "error unmarshall json")
	}

	value, err := json.Marshal(product)
	if err != nil {
		fmt.Fprintf(w, "error marshal into string")
	}

	transaction, err := a.Hyperledger.ProductInvoke(string(value))
	if err != nil {
		http.Error(w, "Cannot post data in blockchain", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, transaction)
}
