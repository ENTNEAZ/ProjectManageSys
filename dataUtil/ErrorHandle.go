package dataUtil

import (
	"net/http"
	"strings"
)

func HandleError(e error, w http.ResponseWriter) bool {
	if e != nil {
		strErr := e.Error()
		strErr = strings.Replace(strErr, "\"", "\\\"", -1)
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"" + strErr + "\"}"))
		return false
	} else {
		return true
	}
}
