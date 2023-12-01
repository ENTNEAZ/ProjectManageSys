package dataStruct

import "encoding/json"

type ResearchRoom struct {
	ResearchRoomID        int
	ResearchRoomName      string
	ResearchRoomDirection string
	Worker_id             string
	Term                  string
	Join_date             string
	Worker_name           string
}

func (r ResearchRoom) ToJson() ([]byte, error) {
	return json.Marshal(r)
}
