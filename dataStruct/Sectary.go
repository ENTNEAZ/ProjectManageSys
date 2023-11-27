package dataStruct

import (
	"encoding/json"
)

type Sectary struct {
	WorkerID         int
	WorkerName       string
	JobDetail        string
	ResearchRoomName string
}

func (s Sectary) ToJson() ([]byte, error) {
	return json.Marshal(s)
}
