package dataUtil

import (
	"Database_Homework/dataStruct"
	"Database_Homework/databaseAccess"
)

func GetAllResearchRoom() ([]dataStruct.ResearchRoom, error) {
	sql := "SELECT * FROM research_room"
	db := databaseAccess.DatabaseConn()
	defer db.Close()

	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var ret []dataStruct.ResearchRoom
	for rows.Next() {
		var researchRoomID int
		var researchRoomName string
		var researchRoomDirection string
		err := rows.Scan(&researchRoomID, &researchRoomName, &researchRoomDirection)
		if err != nil {
			return nil, err
		}
		ret = append(ret, dataStruct.ResearchRoom{
			ResearchRoomID:        researchRoomID,
			ResearchRoomName:      researchRoomName,
			ResearchRoomDirection: researchRoomDirection,
		})

	}
	return ret, nil
}

func AddOrUpdateResearchRoom(id int, name string, direction string) error {
	if id == -1 {
		// insert
		sql := "INSERT INTO research_room (research_room_name, research_room_direction) VALUES (?, ?)"
		db := databaseAccess.DatabaseConn()
		defer db.Close()

		stmt, err := db.Prepare(sql)
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(name, direction)
		if err != nil {
			return err
		}

		return nil
	} else {
		// update
		sql := "UPDATE research_room SET research_room_name = ?, research_room_direction = ? WHERE research_room_id = ?"
		db := databaseAccess.DatabaseConn()
		defer db.Close()

		stmt, err := db.Prepare(sql)
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(name, direction, id)
		if err != nil {
			return err
		}

		return nil
	}
}

func DeleteResearchRoom(id int) error {
	sql := "DELETE FROM research_room WHERE research_room_id = ?"
	db := databaseAccess.DatabaseConn()
	defer db.Close()

	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
