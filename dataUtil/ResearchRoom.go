package dataUtil

import (
	"Database_Homework/dataStruct"
	"Database_Homework/databaseAccess"
	"strconv"
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

func GetAllOrSpecifiedResearchRoomWorker(name_or_id string) ([]byte, error) {
	sql := "select worker_id,worker_name,research_room_name,direction from worker natural join research_room_worker natural join research_room where research_room.research_room_id = ? or research_room.research_room_name LIKE ?"
	if name_or_id == "" {
		sql = "select worker_id,worker_name,research_room_name,direction from worker natural join research_room_worker natural join research_room where 1 = 1 or ? = ?"
	}
	db := databaseAccess.DatabaseConn()

	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(name_or_id, "%"+name_or_id+"%")
	if err != nil {
		return nil, err
	}

	var ret []byte
	ret = append(ret, []byte("[")...)
	for rows.Next() {
		var workerID int
		var workerName string
		var researchRoomName string
		var direction string
		err := rows.Scan(&workerID, &workerName, &researchRoomName, &direction)
		if err != nil {
			return nil, err
		}

		ret = append(ret, []byte("{\"worker_id\": ")...)
		ret = append(ret, []byte(strconv.Itoa(workerID))...)
		ret = append(ret, []byte(", \"worker_name\": \"")...)
		ret = append(ret, []byte(workerName)...)
		ret = append(ret, []byte("\", \"research_room_name\": \"")...)
		ret = append(ret, []byte(researchRoomName)...)
		ret = append(ret, []byte("\", \"direction\": \"")...)
		ret = append(ret, []byte(direction)...)
		ret = append(ret, []byte("\"},")...)
	}

	if len(ret) > 1 {
		ret = ret[:len(ret)-1]
	}

	ret = append(ret, []byte("]")...)

	return ret, nil
}

func AddOrUpdateResearchRoomWorker(worker_id int, research_room_id int, direction string) error {
	// find if exist
	sql := "SELECT * FROM research_room_worker WHERE worker_id = ? AND research_room_id = ?"
	db := databaseAccess.DatabaseConn()
	defer db.Close()

	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query(worker_id, research_room_id)
	if err != nil {
		return err
	}

	if rows.Next() {
		// update
		sql = "UPDATE research_room_worker SET direction = ? WHERE worker_id = ? AND research_room_id = ?"
		stmt, err := db.Prepare(sql)
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(direction, worker_id, research_room_id)
		if err != nil {
			return err
		}
	} else {
		// insert
		sql = "INSERT INTO research_room_worker (worker_id, research_room_id, direction) VALUES (?, ?, ?)"
		stmt, err := db.Prepare(sql)
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(worker_id, research_room_id, direction)
		if err != nil {
			return err
		}
	}

	return nil
}

func DeleteResearchRoomWorker(worker_id int, research_room_id int) error {
	sql := "DELETE FROM research_room_worker WHERE worker_id = ? AND research_room_id = ?"
	db := databaseAccess.DatabaseConn()
	defer db.Close()

	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(worker_id, research_room_id)
	if err != nil {
		return err
	}

	return nil
}
