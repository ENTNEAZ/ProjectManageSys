package apiHandler

import (
	"Database_Homework/dataUtil"
	"net/http"
	"strconv"
)

func GetAllResearchRoom(w http.ResponseWriter, r *http.Request) {
	researchRoom, err := dataUtil.GetAllResearchRoom()
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("{\"code\": -1, \"msg\": \"get all research room failed\"}"))
		return
	}

	var ret = []byte("[")

	for i := 0; i < len(researchRoom); i++ {
		j, err := researchRoom[i].ToJson()
		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("{\"code\": -1, \"msg\": \"research room to json failed\"}"))
			return
		}
		ret = append(ret, append(j, []byte(",")...)...)
	}
	if len(researchRoom) > 1 {
		ret = ret[:len(ret)-1]
	}

	ret = append(ret, []byte("]")...)

	w.WriteHeader(200)
	w.Write(append([]byte("{\"code\": 0, \"msg\": \"success\", \"data\": "), append(ret, []byte("}")...)...))

}

func AddOrUpdateResearchRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	direction := r.URL.Query().Get("direction")

	if name == "" || direction == "" {
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

	err := dataUtil.AddOrUpdateResearchRoom(idi, name, direction)
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