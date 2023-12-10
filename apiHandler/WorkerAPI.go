package apiHandler

import (
	"Database_Homework/dataUtil"
	"net/http"
	"strconv"
)

func GetWorkerByID(w http.ResponseWriter, r *http.Request) {
	if !dataUtil.AutoCookieChecker(w, r) {
		return
	}
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

	w.WriteHeader(200)
	w.Write(append([]byte("{\"code\": 0, \"msg\": \"success\", \"data\": "), append(worker.Str, []byte("}")...)...))

}

func GetAllWorker(w http.ResponseWriter, r *http.Request) {
	if !dataUtil.AutoCookieChecker(w, r) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	workers, err := dataUtil.GetAllWorker()
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"get all worker failed\"}"))
		return
	}

	w.WriteHeader(200)
	w.Write(append([]byte("{\"code\": 0, \"msg\": \"success\", \"data\": "), append(workers.Str, []byte("}")...)...))

}

func AddOrUpdateWorker(w http.ResponseWriter, r *http.Request) {
	if !dataUtil.AutoCookieChecker(w, r) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	name := r.URL.Query().Get("WorkerName")
	gender := r.URL.Query().Get("WorkerGender")
	birth := r.URL.Query().Get("WorkerBirth")
	joinTime := r.URL.Query().Get("WorkerJoinTime")
	job := r.URL.Query().Get("WorkerJob")
	id := r.URL.Query().Get("WorkerId")

	// check empty
	if name == "" || gender == "" || birth == "" || joinTime == "" || job == "" {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"parameter is empty\"}"))
		return
	}

	idi := -1

	if id != "" {
		var err error
		idi, err = strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("{\"code\": -1, \"msg\": \"id is not a number\"}"))
			return
		}
	}

	err := dataUtil.AddOrUpdateWorker(idi, name, gender, birth, joinTime, job)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"add worker failed\"}"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))
}

func DeleteWorkerByID(w http.ResponseWriter, r *http.Request) {
	if !dataUtil.AutoCookieChecker(w, r) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("WorkerId")
	idi, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"id is not a number\"}"))
		return
	}
	err = dataUtil.DeleteWorker(idi)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"delete worker by id failed\"}"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))
}
