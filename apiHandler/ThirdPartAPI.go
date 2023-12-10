package apiHandler

import (
	"Database_Homework/dataUtil"
	"net/http"
)

func GetAllOrSpecified3rdPartInfo(w http.ResponseWriter, r *http.Request) {
	if !dataUtil.AutoCookieChecker(w, r) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	third_part_idname := r.URL.Query().Get("idname")

	res, err := dataUtil.GetAllOrSpecified3rdPartInfo(third_part_idname)
	if !dataUtil.HandleError(err, w) {
		return
	}

	w.WriteHeader(200)
	w.Write(append([]byte("{\"code\": 0, \"msg\": \"success\", \"data\": "), append(res.Str, []byte("}")...)...))
}

func AddOrUpdateOrDelete3rdPartInfo(w http.ResponseWriter, r *http.Request) {
	if !dataUtil.AutoCookieChecker(w, r) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		project_participant_id := r.URL.Query().Get("project_participant_id")
		project_participant_name := r.URL.Query().Get("project_participant_name")
		project_participant_address := r.URL.Query().Get("project_participant_address")
		project_participant_worker_id := r.URL.Query().Get("project_participant_worker_id")

		if project_participant_name == "" || project_participant_address == "" || project_participant_worker_id == "" {
			w.WriteHeader(400)
			w.Write([]byte("{\"code\": -1, \"msg\": \"project_participant_name or project_participant_address or project_participant_worker_id is empty\"}"))
			return
		}

		err := dataUtil.AddOrUpdate3rdPartInfo(project_participant_id, project_participant_name, project_participant_address, project_participant_worker_id)

		if !dataUtil.HandleError(err, w) {
			return
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))
	} else if r.Method == "DELETE" {
		project_participant_id := r.URL.Query().Get("project_participant_id")

		if project_participant_id == "" {
			w.WriteHeader(400)
			w.Write([]byte("{\"code\": -1, \"msg\": \"project_participant_id is empty\"}"))
			return
		}

		err := dataUtil.Delete3rdPartInfo(project_participant_id)
		if !dataUtil.HandleError(err, w) {
			return
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))
	}
}

func GetAllOrSpecified3rdPartContact(w http.ResponseWriter, r *http.Request) {
	if !dataUtil.AutoCookieChecker(w, r) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	third_part_idname := r.URL.Query().Get("idname")

	res, err := dataUtil.GetAllOrSpecified3rdPartContact(third_part_idname)
	if !dataUtil.HandleError(err, w) {
		return
	}

	w.WriteHeader(200)
	w.Write(append([]byte("{\"code\": 0, \"msg\": \"success\", \"data\": "), append(res.Str, []byte("}")...)...))
}

func AddOrUpdateOrDelete3rdPartContact(w http.ResponseWriter, r *http.Request) {
	if !dataUtil.AutoCookieChecker(w, r) {
		return
	}
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		project_participant_worker_id := r.URL.Query().Get("project_participant_worker_id")
		project_participant_worker_telephone := r.URL.Query().Get("project_participant_worker_telephone")
		project_participant_worker_mobile := r.URL.Query().Get("project_participant_worker_mobile")
		project_participant_worker_email := r.URL.Query().Get("project_participant_worker_email")

		if project_participant_worker_telephone == "" || project_participant_worker_mobile == "" || project_participant_worker_email == "" {
			w.WriteHeader(400)
			w.Write([]byte("{\"code\": -1, \"msg\": \"project_participant_worker_telephone or project_participant_worker_mobile or project_participant_worker_email or project_participant_id is empty\"}"))
			return
		}

		err := dataUtil.AddOrUpdate3rdPartContact(project_participant_worker_id, project_participant_worker_telephone, project_participant_worker_mobile, project_participant_worker_email)

		if !dataUtil.HandleError(err, w) {
			return
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))
	} else if r.Method == "DELETE" {
		w.Header().Set("Content-Type", "application/json")
		project_participant_worker_id := r.URL.Query().Get("project_participant_worker_id")

		if project_participant_worker_id == "" {
			w.WriteHeader(400)
			w.Write([]byte("{\"code\": -1, \"msg\": \"project_participant_worker_id or project_participant_id is empty\"}"))
			return
		}

		err := dataUtil.Delete3rdPartContact(project_participant_worker_id)
		if !dataUtil.HandleError(err, w) {
			return
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))
	}
}

func AddOrDeleteContactRelation(w http.ResponseWriter, r *http.Request) {
	if !dataUtil.AutoCookieChecker(w, r) {
		return
	}
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		project_participant_worker_id := r.URL.Query().Get("project_participant_worker_id")
		project_participant_id := r.URL.Query().Get("project_participant_id")

		if project_participant_worker_id == "" || project_participant_id == "" {
			w.WriteHeader(400)
			w.Write([]byte("{\"code\": -1, \"msg\": \"project_participant_worker_id or project_participant_id is empty\"}"))
			return
		}

		err := dataUtil.AddContactRelation(project_participant_worker_id, project_participant_id)
		if !dataUtil.HandleError(err, w) {
			return
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))
	} else if r.Method == "DELETE" {
		w.Header().Set("Content-Type", "application/json")
		project_participant_worker_id := r.URL.Query().Get("project_participant_worker_id")
		project_participant_id := r.URL.Query().Get("project_participant_id")

		if project_participant_worker_id == "" || project_participant_id == "" {
			w.WriteHeader(400)
			w.Write([]byte("{\"code\": -1, \"msg\": \"project_participant_worker_id or project_participant_id is empty\"}"))
			return
		}

		err := dataUtil.DeleteContactRelation(project_participant_worker_id, project_participant_id)
		if !dataUtil.HandleError(err, w) {
			return
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))
	}
}
