package apiHandler

import (
	"Database_Homework/dataUtil"
	"net/http"
)

func GetAllOrSpecifiedProject(w http.ResponseWriter, r *http.Request) {
	if !dataUtil.AutoCookieChecker(w, r) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	project_idname := r.URL.Query().Get("project_idname")

	res, err := dataUtil.GetAllOrSpecifiedProject(project_idname)
	if !dataUtil.HandleError(err, w) {
		return
	}

	w.WriteHeader(200)
	w.Write(append([]byte("{\"code\": 0, \"msg\": \"success\", \"data\": "), append(res.Str, []byte("}")...)...))
}

func AddOrUpdateOrDeleteProject(w http.ResponseWriter, r *http.Request) {
	if !dataUtil.AutoCookieChecker(w, r) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		project_id := r.URL.Query().Get("project_id")
		project_name := r.URL.Query().Get("project_name")
		project_detail := r.URL.Query().Get("project_detail")
		project_start_time := r.URL.Query().Get("project_start_time")
		project_end_time := r.URL.Query().Get("project_end_time")
		project_fund := r.URL.Query().Get("project_fund")
		worker_id := r.URL.Query().Get("worker_id")
		project_participant_id := r.URL.Query().Get("project_participant_id")
		project_supervisor_id := r.URL.Query().Get("project_supervisor_id")

		if project_name == "" || project_detail == "" || project_start_time == "" || project_end_time == "" || project_fund == "" || worker_id == "" || project_participant_id == "" {
			w.WriteHeader(400)
			w.Write([]byte("{\"code\": -1, \"msg\": \"project_name or project_detail or project_start_time or project_end_time or project_fund or worker_id or project_participant_id is empty\"}"))
			return
		}

		err := dataUtil.AddOrUpdateProject(project_id, project_name, project_detail, project_start_time, project_end_time, project_fund, worker_id, project_participant_id, project_supervisor_id)
		if !dataUtil.HandleError(err, w) {
			return
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))
	} else if r.Method == "DELETE" {
		project_id := r.URL.Query().Get("project_id")

		if project_id == "" {
			w.WriteHeader(400)
			w.Write([]byte("{\"code\": -1, \"msg\": \"project_id is empty\"}"))
			return
		}

		err := dataUtil.DeleteProject(project_id)
		if !dataUtil.HandleError(err, w) {
			return
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))
	} else {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"method is not GET or DELETE\"}"))
		return
	}
}
func FindAllWorkerInProject(w http.ResponseWriter, r *http.Request) {
	if !dataUtil.AutoCookieChecker(w, r) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	project_id := r.URL.Query().Get("idname")

	res, err := dataUtil.FindAllWorkerInProject(project_id)
	if !dataUtil.HandleError(err, w) {
		return
	}

	w.WriteHeader(200)
	w.Write(append([]byte("{\"code\": 0, \"msg\": \"success\", \"data\": "), append(res.Str, []byte("}")...)...))
}

func AddOrDeleteProjectWorker(w http.ResponseWriter, r *http.Request) {
	if !dataUtil.AutoCookieChecker(w, r) {
		return
	}
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		project_id := r.URL.Query().Get("project_id")
		worker_id := r.URL.Query().Get("worker_id")

		if project_id == "" || worker_id == "" {
			w.WriteHeader(400)
			w.Write([]byte("{\"code\": -1, \"msg\": \"project_id or worker_id is empty\"}"))
			return
		}

		err := dataUtil.AddProjectWorker(project_id, worker_id)
		if !dataUtil.HandleError(err, w) {
			return
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))
	} else if r.Method == "DELETE" {
		w.Header().Set("Content-Type", "application/json")
		project_id := r.URL.Query().Get("project_id")
		worker_id := r.URL.Query().Get("worker_id")

		if project_id == "" || worker_id == "" {
			w.WriteHeader(400)
			w.Write([]byte("{\"code\": -1, \"msg\": \"project_id or worker_id is empty\"}"))
			return
		}

		err := dataUtil.DeleteProjectWorker(project_id, worker_id)
		if !dataUtil.HandleError(err, w) {
			return
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))
	} else {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"method is not GET or DELETE\"}"))
		return
	}
}

func FindAllParticipantInProject(w http.ResponseWriter, r *http.Request) {
	if !dataUtil.AutoCookieChecker(w, r) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	project_id := r.URL.Query().Get("idname")

	res, err := dataUtil.FindAllParticipantInProject(project_id)
	if !dataUtil.HandleError(err, w) {
		return
	}

	w.WriteHeader(200)
	w.Write(append([]byte("{\"code\": 0, \"msg\": \"success\", \"data\": "), append(res.Str, []byte("}")...)...))
}

func AddOrDeleteProjectParticipant(w http.ResponseWriter, r *http.Request) {
	if !dataUtil.AutoCookieChecker(w, r) {
		return
	}
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		project_id := r.URL.Query().Get("project_id")
		project_participant_id := r.URL.Query().Get("project_participant_id")

		if project_id == "" || project_participant_id == "" {
			w.WriteHeader(400)
			w.Write([]byte("{\"code\": -1, \"msg\": \"project_id or project_participant_id is empty\"}"))
			return
		}

		err := dataUtil.AddProjectParticipant(project_id, project_participant_id)
		if !dataUtil.HandleError(err, w) {
			return
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))
	} else if r.Method == "DELETE" {
		w.Header().Set("Content-Type", "application/json")
		project_id := r.URL.Query().Get("project_id")
		project_participant_id := r.URL.Query().Get("project_participant_id")

		if project_id == "" || project_participant_id == "" {
			w.WriteHeader(400)
			w.Write([]byte("{\"code\": -1, \"msg\": \"project_id or project_participant_id is empty\"}"))
			return
		}

		err := dataUtil.DeleteProjectParticipant(project_id, project_participant_id)
		if !dataUtil.HandleError(err, w) {
			return
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))
	} else {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"method is not GET or DELETE\"}"))
		return
	}
}

func GetAllOrSpecifiedProjectFruit(w http.ResponseWriter, r *http.Request) {
	if !dataUtil.AutoCookieChecker(w, r) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	project_id := r.URL.Query().Get("idname")

	res, err := dataUtil.GetAllOrSpecifiedProjectFruit(project_id)
	if !dataUtil.HandleError(err, w) {
		return
	}

	w.WriteHeader(200)
	w.Write(append([]byte("{\"code\": 0, \"msg\": \"success\", \"data\": "), append(res.Str, []byte("}")...)...))
}

func AddOrUpdateOrDeleteProjectFruit(w http.ResponseWriter, r *http.Request) {
	if !dataUtil.AutoCookieChecker(w, r) {
		return
	}
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		project_id := r.URL.Query().Get("project_id")
		worker_id := r.URL.Query().Get("worker_id")
		project_fruit_id := r.URL.Query().Get("project_fruit_id")
		project_fruit_get_time := r.URL.Query().Get("project_fruit_get_time")
		project_fruit_master_rank := r.URL.Query().Get("project_fruit_master_rank")
		project_fruit_type := r.URL.Query().Get("project_fruit_type")
		project_fruit_detail := r.URL.Query().Get("project_fruit_detail")

		if project_id == "" || worker_id == "" || project_fruit_get_time == "" || project_fruit_master_rank == "" || project_fruit_type == "" || project_fruit_detail == "" {
			w.WriteHeader(400)
			w.Write([]byte("{\"code\": -1, \"msg\": \"missing params\"}"))
			return
		}

		err := dataUtil.AddOrUpdateProjectFruit(project_id, worker_id, project_fruit_id, project_fruit_get_time, project_fruit_master_rank, project_fruit_type, project_fruit_detail)

		if !dataUtil.HandleError(err, w) {
			return
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))

	} else if r.Method == "DELETE" {
		w.Header().Set("Content-Type", "application/json")
		project_fruit_id := r.URL.Query().Get("project_fruit_id")

		if project_fruit_id == "" {
			w.WriteHeader(400)
			w.Write([]byte("{\"code\": -1, \"msg\": \"project_fruit_id is empty\"}"))
			return
		}

		err := dataUtil.DeleteProjectFruit(project_fruit_id)
		if !dataUtil.HandleError(err, w) {
			return
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))
	} else {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"method is not GET or DELETE\"}"))
		return
	}
}
