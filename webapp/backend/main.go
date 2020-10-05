package webapp

import (
	"net/http"

	"github.com/MaJloe3Jlo/HLexampleGO/webapp/backend/handlers"
)

func Server(app *handlers.App) {
	http.HandleFunc("/api/get", app.HomeHandler)
	http.HandleFunc("/api/post", app.ChangeHandler)
	http.HandleFunc("/api/list", app.AllHandler)

	http.ListenAndServe(":7777", nil)
}
