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
	if !dataUtil.HandleError(err, w) {
		return
	}
	worker, err := dataUtil.GetWorkerByID(idi)

	if !dataUtil.HandleError(err, w) {
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
	if !dataUtil.HandleError(err, w) {
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
		if !dataUtil.HandleError(err, w) {
			return
		}
	}

	err := dataUtil.AddOrUpdateWorker(idi, name, gender, birth, joinTime, job)

	if !dataUtil.HandleError(err, w) {
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
	if !dataUtil.HandleError(err, w) {
		return
	}
	err = dataUtil.DeleteWorker(idi)

	if !dataUtil.HandleError(err, w) {
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))
}
