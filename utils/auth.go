package utils

import (
	"github.com/gorilla/securecookie"
	"net/http"
)

// cookie handling

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func GetUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}

func setSession(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

// login handler

func LoginHandler(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	pass := request.FormValue("password")
	redirectTarget := "/login"
	if name != "" && pass != "" {
		// .. check credentials ..
		setSession(name, response)
		redirectTarget = "/"
	}
	http.Redirect(response, request, redirectTarget, 302)
}

// logout handler

func LogoutHandler(response http.ResponseWriter, request *http.Request) {
	clearSession(response)
	http.Redirect(response, request, "/login", 302)
}


// internal page

// const internalPage = `
// <h1>Internal</h1>
// <hr>
// <small>User: %s</small>
// <form method="post" action="/logout">
//     <button type="submit">Logout</button>
// </form>
// `
