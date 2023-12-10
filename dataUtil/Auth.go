package dataUtil

import (
	"Database_Homework/databaseAccess"
	"math/rand"
	"net/http"
	"strconv"
)

func Auth(username string, password string, isRegister bool) (bool, string, error) {
	if isRegister {
		db := databaseAccess.DatabaseConn()
		sql := "select * from registered_user where username = ?"
		rows, err := db.Query(sql, username)
		if err != nil {
			return false, "", err
		}
		defer rows.Close()
		if rows.Next() {
			return false, "", nil
		} else {
			sql = "insert into registered_user(username, password) values(?, ?)"
			_, err := db.Exec(sql, username, password)
			if err != nil {
				return false, "", err
			}
			cookie := getRandomCookie()
			sql = "insert into cookie values(?, ?)"
			_, err = db.Exec(sql, cookie, username)
			if err != nil {
				return false, "", err
			}
			return true, cookie, nil
		}
	} else {
		db := databaseAccess.DatabaseConn()
		sql := "select * from registered_user where username = ? and password = ?"
		rows, err := db.Query(sql, username, password)
		if err != nil {
			return false, "", err
		}
		defer rows.Close()
		if rows.Next() {
			cookie := getRandomCookie()
			sql = "insert into cookie values(?, ?)"
			_, err = db.Exec(sql, cookie, username)
			if err != nil {
				return false, "", err
			}
			return true, cookie, nil
		} else {
			return false, "", nil
		}
	}
}

func getRandomCookie() (s string) {
	for i := 0; i < 32; i++ {
		s += strconv.Itoa('a' + rand.Int()%26)
	}
	return
}

func GetCookieUser(cookie string) (s string, err error) {
	db := databaseAccess.DatabaseConn()
	sql := "select username from cookie where cookie = ?"
	rows, err := db.Query(sql, cookie)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&s)
		return s, nil
	} else {
		return "", nil
	}
}

func CheckCookieValid(r *http.Request) bool {
	cookie, err := r.Cookie("SessionID")
	if err != nil || cookie.Value == "" {
		return false
	}

	res, err := GetCookieUser(cookie.Value)
	if err != nil || res == "" {
		return false
	}
	return true
}

func AutoCookieChecker(w http.ResponseWriter, r *http.Request) bool {
	valid := CheckCookieValid(r)
	if valid {
		return true
	} else {
		w.Header().Set("Location", "/")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(401)
		w.Write([]byte("{\"code\": -1, \"msg\": \"not login\"}"))
		return false
	}
}
