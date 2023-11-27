package dataStruct

import "encoding/json"

type WorkingArea struct {
	WorkingAreaID      int
	ResearchRoomName   string
	WorkingAreaSize    int
	WorkingAreaAddress string
}

func (w WorkingArea) ToJson() ([]byte, error) {
	return json.Marshal(w)
}
