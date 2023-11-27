package dataStruct

import "encoding/json"

type ResearchRoom struct {
	ResearchRoomID        int
	ResearchRoomName      string
	ResearchRoomDirection string
}

func (r ResearchRoom) ToJson() ([]byte, error) {
	return json.Marshal(r)
}
