package utils

import (
	"github.com/gorilla/securecookie"
	"net/http"
	"strconv"
)

// cookie handling
var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func GetUser(request *http.Request) (userName string, userId int) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
			userId, _ = strconv.Atoi(cookieValue["user_id"])
		}
	}
	return userName, userId
}

func setSession(userName string, userId int, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
		"user_id": strconv.Itoa(userId),
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
	passwd := request.FormValue("password")
	redirectTarget := "/login"
	if name != "" && passwd != "" {
		// .. check credentials ..
		user_id := ReadUserId(name, passwd)
		// match name or password
		if user_id != 0{
			setSession(name, user_id, response)
			redirectTarget = "/"
		}	
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
