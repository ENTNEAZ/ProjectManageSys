package apiHandler

import (
	"Database_Homework/dataUtil"
	"net/http"
	"strconv"
)

func GetAllResearchRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	researchRoom, err := dataUtil.GetAllResearchRoom()
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"get all research room failed\"}"))
		return
	}

	w.WriteHeader(200)
	w.Write(append([]byte("{\"code\": 0, \"msg\": \"success\", \"data\": "), append(researchRoom, []byte("}")...)...))

}

func AddOrUpdateResearchRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	direction := r.URL.Query().Get("direction")
	work_id := r.URL.Query().Get("worker_id")
	term := r.URL.Query().Get("term")
	join_date := r.URL.Query().Get("join_date")

	if name == "" || direction == "" || work_id == "" || term == "" || join_date == "" {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"name or direction is empty\"}"))
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

	err := dataUtil.AddOrUpdateResearchRoom(idi, name, direction, work_id, term, join_date)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"add or update research room failed\"}"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))
}

func DeleteResearchRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")

	if id == "" {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"id is empty\"}"))
		return
	}

	idi, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"id is not a number\"}"))
		return
	}

	err = dataUtil.DeleteResearchRoom(idi)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"delete research room failed\"}"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))
}

func GetAllOrSpecifiedResearchRoomWorker(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name_or_id := r.URL.Query().Get("name_or_id")

	res, err := dataUtil.GetAllOrSpecifiedResearchRoomWorker(name_or_id)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"get all or specified research room worker failed\"}"))
		return
	}

	w.WriteHeader(200)
	w.Write(append([]byte("{\"code\": 0, \"msg\": \"success\", \"data\": "), append(res, []byte("}")...)...))
}

func AddOrUpdateResearchRoomWorker(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	workerID := r.URL.Query().Get("worker_id")
	roomID := r.URL.Query().Get("research_room_id")
	direction := r.URL.Query().Get("direction")

	if workerID == "" || roomID == "" || direction == "" {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"worker id or research room id or direction is empty\"}"))
		return
	}

	workerIDI, err := strconv.Atoi(workerID)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"worker id is not a number\"}"))
		return
	}

	roomIDI, err := strconv.Atoi(roomID)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"research room id is not a number\"}"))
		return
	}

	err = dataUtil.AddOrUpdateResearchRoomWorker(workerIDI, roomIDI, direction)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"add or update research room worker failed\"}"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))

}

func DeleteResearchRoomWorker(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	workerID := r.URL.Query().Get("worker_id")
	roomID := r.URL.Query().Get("research_room_id")

	if workerID == "" || roomID == "" {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"worker id or research room id is empty\"}"))
		return
	}

	workerIDI, err := strconv.Atoi(workerID)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"worker id is not a number\"}"))
		return
	}

	roomIDI, err := strconv.Atoi(roomID)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"research room id is not a number\"}"))
		return
	}

	err = dataUtil.DeleteResearchRoomWorker(workerIDI, roomIDI)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"delete research room worker failed\"}"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("{\"code\": 0, \"msg\": \"success\"}"))
}
