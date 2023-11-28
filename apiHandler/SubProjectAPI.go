package apiHandler

import (
	"Database_Homework/dataUtil"
	"net/http"
)

func FindAllSubProjectInProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idname := r.URL.Query().Get("idname")

	res, err := dataUtil.FindAllSubProjectInProject(idname)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"find all sub project in project error\"}"))
		return
	}

	w.WriteHeader(200)
	w.Write(append([]byte("{\"code\": 0, \"msg\": \"success\", \"data\": "), append(res, []byte("}")...)...))

}

func AddOrUpdateOrDeleteSpecifiedSubProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		project_id := r.URL.Query().Get("project_id")
		sub_project_id := r.URL.Query().Get("sub_project_id")
		worker_id := r.URL.Query().Get("worker_id")
		sub_project_end_time := r.URL.Query().Get("sub_project_end_time")
		sub_project_fund := r.URL.Query().Get("sub_project_fund")
		sub_project_tech_detail := r.URL.Query().Get("sub_project_tech_detail")

		if project_id == "" || worker_id == "" || sub_project_end_time == "" || sub_project_fund == "" || sub_project_tech_detail == "" {
			w.WriteHeader(400)
			w.Write([]byte("{\"code\": -1, \"msg\": \"project_id or worker_id or sub_project_end_time or sub_project_fund or sub_project_tech_detail is empty\"}"))
			return
		}

		err := dataUtil.AddOrUpdateSubProject(project_id, sub_project_id, worker_id, sub_project_end_time, sub_project_fund, sub_project_tech_detail)

		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("{\"code\": -1, \"msg\": \"add or update sub project error\"}"))
			return
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))
	} else if r.Method == "DELETE" {
		sub_project_id := r.URL.Query().Get("sub_project_id")

		if sub_project_id == "" {
			w.WriteHeader(400)
			w.Write([]byte("{\"code\": -1, \"msg\": \"sub_project_id is empty\"}"))
			return
		}

		err := dataUtil.DeleteSubProject(sub_project_id)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("{\"code\": -1, \"msg\": \"delete sub project error\"}"))
			return
		}

		w.WriteHeader(200)
		w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))
	}
}
