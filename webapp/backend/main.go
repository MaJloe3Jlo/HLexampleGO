package webapp

import (
	"net/http"

	"github.com/MaJloe3Jlo/HLexampleGO/webapp/backend/handlers"
)

func Server(app *handlers.App) {
	http.HandleFunc("/", app.HomeHandler)
	http.HandleFunc("/post", app.ChangeHandler)

	http.ListenAndServe(":7777", nil)
}
