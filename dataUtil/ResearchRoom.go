package dataUtil

import (
	"Database_Homework/databaseAccess"
	"Database_Homework/jsonHelper"
)

func GetAllResearchRoom() (*jsonHelper.JsonStr, error) {
	sql := "SELECT research_room_id,research_room_name,research_room_direction,worker.worker_id,worker.worker_name,term,join_date FROM research_room join worker on worker.worker_id = research_room.worker_id"
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

	ret := new(jsonHelper.JsonStr)
	ret.JsonArrayInit()
	for rows.Next() {
		var researchRoomID int
		var researchRoomName string
		var researchRoomDirection string
		var worker_id string
		var term string
		var join_date string
		var worker_name string
		err := rows.Scan(&researchRoomID, &researchRoomName, &researchRoomDirection, &worker_id, &worker_name, &term, &join_date)
		if err != nil {
			return nil, err
		}
		var temp jsonHelper.JsonStr
		temp.JsonDictInit()
		temp.JsonDictAddStrInt("ResearchRoomID", researchRoomID)
		temp.JsonDictAddStrStr("ResearchRoomName", researchRoomName)
		temp.JsonDictAddStrStr("ResearchRoomDirection", researchRoomDirection)
		temp.JsonDictAddStrStr("Worker_id", worker_id)
		temp.JsonDictAddStrStr("Worker_name", worker_name)
		temp.JsonDictAddStrStr("Term", term)
		temp.JsonDictAddStrStr("Join_date", join_date)
		temp.JsonDictEnd()
		ret.JsonArrayAddJson(temp)
	}
	ret.JsonArrayEnd()
	return ret, nil
}

func AddOrUpdateResearchRoom(id int, name string, direction string, work_id, term, join_date string) error {
	if id == -1 {
		// insert
		sql := "INSERT INTO research_room (research_room_name, research_room_direction, worker_id, term, join_date) VALUES (?, ?, ?, ?, ?)"
		db := databaseAccess.DatabaseConn()
		defer db.Close()

		stmt, err := db.Prepare(sql)
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(name, direction, work_id, term, join_date)
		if err != nil {
			return err
		}

		return nil
	} else {
		// update
		sql := "UPDATE research_room SET research_room_name = ?, research_room_direction = ? , worker_id = ? , term = ? , join_date = ? WHERE research_room_id = ?"
		db := databaseAccess.DatabaseConn()
		defer db.Close()

		stmt, err := db.Prepare(sql)
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(name, direction, work_id, term, join_date, id)
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

func GetAllOrSpecifiedResearchRoomWorker(name_or_id string) (*jsonHelper.JsonStr, error) {
	sql := "select worker.worker_id,worker_name,research_room_name,direction from worker join research_room_worker on worker.worker_id = research_room_worker.worker_id join research_room on research_room_worker.research_room_id = research_room.research_room_id where research_room.research_room_id = ? or research_room.research_room_name LIKE ?"
	if name_or_id == "" {
		sql = "select worker.worker_id,worker_name,research_room_name,direction from worker join research_room_worker on worker.worker_id = research_room_worker.worker_id join research_room on research_room_worker.research_room_id = research_room.research_room_id where 1 = 1 or ? = ?"
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

	ret := new(jsonHelper.JsonStr)
	ret.JsonArrayInit()
	for rows.Next() {
		var workerID int
		var workerName string
		var researchRoomName string
		var direction string
		err := rows.Scan(&workerID, &workerName, &researchRoomName, &direction)
		if err != nil {
			return nil, err
		}

		var temp jsonHelper.JsonStr
		temp.JsonDictInit()
		temp.JsonDictAddStrInt("worker_id", workerID)
		temp.JsonDictAddStrStr("worker_name", workerName)
		temp.JsonDictAddStrStr("research_room_name", researchRoomName)
		temp.JsonDictAddStrStr("direction", direction)
		temp.JsonDictEnd()
		ret.JsonArrayAddJson(temp)
	}
	ret.JsonArrayEnd()
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
