package apiHandler

import (
	"Database_Homework/dataUtil"
	"Database_Homework/jsonHelper"
	"net/http"
	"strconv"
)

func GetAllOrSpecifiedSectary(w http.ResponseWriter, r *http.Request) {
	if !dataUtil.AutoCookieChecker(w, r) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id_or_name")
	idi, err := strconv.Atoi(id)
	isIDInt := err == nil

	var ret jsonHelper.JsonStr

	if isIDInt {
		rs, err := dataUtil.GetSectaryByResearchRoomID(idi)
		if !dataUtil.HandleError(err, w) {
			return
		}

		ret.CombineWith(*rs)
	}

	if id == "" {
		rs, err := dataUtil.GetAllSectary()
		if !dataUtil.HandleError(err, w) {
			return
		}

		ret.CombineWith(*rs)
	} else {
		rs, err := dataUtil.GetSectaryByResearchRoomName(id)
		if !dataUtil.HandleError(err, w) {
			return
		}

		ret.CombineWith(*rs)
	}

	w.WriteHeader(200)
	w.Write(append([]byte("{\"code\": 0, \"msg\": \"success\", \"data\": "), append(ret.Str, []byte("}")...)...))

}

func AddOrUpdateSectary(w http.ResponseWriter, r *http.Request) {
	if !dataUtil.AutoCookieChecker(w, r) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("worker_id")
	roomID := r.URL.Query().Get("research_room_id")
	jobDetail := r.URL.Query().Get("job_detail")

	idi, err := strconv.Atoi(id)
	if !dataUtil.HandleError(err, w) {
		return
	}

	idi2, err := strconv.Atoi(roomID)
	if !dataUtil.HandleError(err, w) {
		return
	}

	err = dataUtil.AddOrUpdateSectary(idi, idi2, jobDetail)
	if !dataUtil.HandleError(err, w) {
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))

}

func DeleteSectary(w http.ResponseWriter, r *http.Request) {
	if !dataUtil.AutoCookieChecker(w, r) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("worker_id")
	roomID := r.URL.Query().Get("research_room_id")

	idi, err := strconv.Atoi(id)
	if !dataUtil.HandleError(err, w) {
		return
	}

	idi2, err := strconv.Atoi(roomID)
	if !dataUtil.HandleError(err, w) {
		return
	}

	err = dataUtil.DeleteSectary(idi, idi2)
	if !dataUtil.HandleError(err, w) {
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))
}
