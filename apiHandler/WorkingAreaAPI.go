package apiHandler

import (
	"Database_Homework/dataUtil"
	"Database_Homework/jsonHelper"
	"net/http"
	"strconv"
)

func GetAllOrSpecifiedWorkingArea(w http.ResponseWriter, r *http.Request) {
	if !dataUtil.AutoCookieChecker(w, r) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id_or_name")
	idi, err := strconv.Atoi(id)
	isIDInt := err == nil

	var ret jsonHelper.JsonStr
	if isIDInt {
		rs, err := dataUtil.GetWorkingAreaByResearchRoomID(idi)
		if !dataUtil.HandleError(err, w) {
			return
		}
		ret.CombineWith(*rs)
	}
	if id == "" {
		rs, err := dataUtil.GetAllWorkingArea()
		if !dataUtil.HandleError(err, w) {
			return
		}

		ret.CombineWith(*rs)
	} else {
		rs, err := dataUtil.GetWorkingAreaByResearchRoomName(id)
		if !dataUtil.HandleError(err, w) {
			return
		}

		ret.CombineWith(*rs)
	}

	w.WriteHeader(200)
	w.Write(append([]byte("{\"code\": 0, \"msg\": \"success\", \"data\": "), append(ret.Str, []byte("}")...)...))
}

func AddOrUpdateWorkingAreaSubmit(w http.ResponseWriter, r *http.Request) {
	if !dataUtil.AutoCookieChecker(w, r) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	size := r.URL.Query().Get("size")
	address := r.URL.Query().Get("address")

	if size == "" || address == "" {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"size or address is empty\"}"))
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

	sizei, err := strconv.Atoi(size)
	if !dataUtil.HandleError(err, w) {
		return
	}

	err = dataUtil.AddOrUpdateWorkingArea(idi, sizei, address)

	if !dataUtil.HandleError(err, w) {
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))

}

func AddWorkingAreaForResearchRoom(w http.ResponseWriter, r *http.Request) {
	if !dataUtil.AutoCookieChecker(w, r) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	researchRoomID := r.URL.Query().Get("research_room_id")
	workingAreaID := r.URL.Query().Get("working_area_id")

	if researchRoomID == "" || workingAreaID == "" {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"research_room_id or working_area_id is empty\"}"))
		return
	}

	researchRoomIDi, err := strconv.Atoi(researchRoomID)
	if !dataUtil.HandleError(err, w) {
		return
	}

	workingAreaIDi, err := strconv.Atoi(workingAreaID)
	if !dataUtil.HandleError(err, w) {
		return
	}

	err = dataUtil.AddWorkingAreaForResearchRoom(workingAreaIDi, researchRoomIDi)
	if !dataUtil.HandleError(err, w) {
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))

}

func DeleteWorkingAreaForResearchRoom(w http.ResponseWriter, r *http.Request) {
	if !dataUtil.AutoCookieChecker(w, r) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	researchRoomID := r.URL.Query().Get("research_room_id")
	workingAreaID := r.URL.Query().Get("working_area_id")

	if researchRoomID == "" || workingAreaID == "" {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"research_room_id or working_area_id is empty\"}"))
		return
	}

	researchRoomIDi, err := strconv.Atoi(researchRoomID)
	if !dataUtil.HandleError(err, w) {
		return
	}

	workingAreaIDi, err := strconv.Atoi(workingAreaID)
	if !dataUtil.HandleError(err, w) {
		return
	}

	err = dataUtil.DeleteWorkingAreaForResearchRoom(workingAreaIDi, researchRoomIDi)
	if !dataUtil.HandleError(err, w) {
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))
}
