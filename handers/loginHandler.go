package handers

import (
	_ "github.com/gorilla/securecookie"
	"net/http"
	"github.com/gorilla/sessions"
)


var (
	store *sessions.CookieStore = sessions.NewCookieStore()
	users map[string]string     = make(map[string]string)
)


func Register() {

}

func setSession(name string, w http.ResponseWriter) {
}
