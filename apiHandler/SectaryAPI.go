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
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("{\"code\": -1, \"msg\": \"get sectary by research room id failed\"}"))
			return
		}

		ret.CombineWith(*rs)
	}

	if id == "" {
		rs, err := dataUtil.GetAllSectary()
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("{\"code\": -1, \"msg\": \"get all sectary failed\"}"))
			return
		}

		ret.CombineWith(*rs)
	} else {
		rs, err := dataUtil.GetSectaryByResearchRoomName(id)
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("{\"code\": -1, \"msg\": \"get sectary by research room name failed\"}"))
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
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"worker id is not int\"}"))
		return
	}

	idi2, err := strconv.Atoi(roomID)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"research room id is not int\"}"))
		return
	}

	err = dataUtil.AddOrUpdateSectary(idi, idi2, jobDetail)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"add or update sectary failed\"}"))
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
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"worker id is not int\"}"))
		return
	}

	idi2, err := strconv.Atoi(roomID)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"research room id is not int\"}"))
		return
	}

	err = dataUtil.DeleteSectary(idi, idi2)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"delete sectary failed\"}"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))
}
