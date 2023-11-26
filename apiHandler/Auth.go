package apiHandler

import (
	"Database_Homework/dataUtil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("hashPassword")

	success, cookie, err := dataUtil.Auth(username, password, false)
	if err != nil || !success {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"login failed\"}"))
		return
	}

	w.Header().Set("Set-Cookie", cookie)
	w.Header().Set("Location", "/js/")
	w.WriteHeader(302)
	w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))
	return
}

func Register(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("hashPassword")

	success, cookie, err := dataUtil.Auth(username, password, true)
	if err != nil || !success {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"register failed\"}"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Set-Cookie", cookie)
	w.Header().Add("Location", "/js/")
	w.WriteHeader(302)
	w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))
	return
}
