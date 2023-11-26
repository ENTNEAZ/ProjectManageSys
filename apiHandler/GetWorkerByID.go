package apiHandler

import (
	"Database_Homework/dataUtil"
	"net/http"
	"strconv"
)

func GetWorkerByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	idi, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"id is not a number\"}"))
		return
	}
	worker, err := dataUtil.GetWorkerByID(idi)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"get worker by id failed\"}"))
		return
	}
	j, err := worker.ToJson()
	if err != nil {
		w.WriteHeader(400)

		w.Write([]byte("{\"code\": -1, \"msg\": \"worker to json failed\"}"))
		return
	}
	w.WriteHeader(200)
	w.Write(append([]byte("{\"code\": 0, \"msg\": \"success\", \"data\": "), append(j, []byte("}")...)...))

}
