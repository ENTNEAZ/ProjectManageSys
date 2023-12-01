package dataUtil

import (
	"Database_Homework/dataStruct"
	"Database_Homework/databaseAccess"
)

func GetAllSectary() ([]dataStruct.Sectary, error) {
	sql := "SELECT worker.worker_id,worker_name,job_detail,research_room_name FROM research_room join research_room_sectary on research_room.research_room_id = research_room_sectary.research_room_id join worker on research_room_sectary.worker_id = worker.worker_id"
	db := databaseAccess.DatabaseConn()

	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var ret []dataStruct.Sectary
	for rows.Next() {
		var workerID int
		var workerName string
		var jobDetail string
		var researchRoomName string
		err := rows.Scan(&workerID, &workerName, &jobDetail, &researchRoomName)
		if err != nil {
			return nil, err
		}
		ret = append(ret, dataStruct.Sectary{
			WorkerID:         workerID,
			WorkerName:       workerName,
			JobDetail:        jobDetail,
			ResearchRoomName: researchRoomName,
		})
	}
	return ret, nil
}

func GetSectaryByResearchRoomName(name string) ([]dataStruct.Sectary, error) {
	sql := "SELECT worker_id,worker_name,job_detail,research_room_name FROM research_room NATURAL JOIN research_room_sectary NATURAL JOIN worker WHERE research_room_name LIKE ?"
	db := databaseAccess.DatabaseConn()

	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query("%" + name + "%")
	if err != nil {
		return nil, err
	}

	var ret []dataStruct.Sectary
	for rows.Next() {
		var workerID int
		var workerName string
		var jobDetail string
		var researchRoomName string
		err := rows.Scan(&workerID, &workerName, &jobDetail, &researchRoomName)
		if err != nil {
			return nil, err
		}
		ret = append(ret, dataStruct.Sectary{
			WorkerID:         workerID,
			WorkerName:       workerName,
			JobDetail:        jobDetail,
			ResearchRoomName: researchRoomName,
		})
	}
	return ret, nil
}

func GetSectaryByResearchRoomID(id int) ([]dataStruct.Sectary, error) {
	sql := "SELECT worker.worker_id,worker_name,job_detail,research_room_name FROM research_room join research_room_sectary on research_room.research_room_id = research_room_sectary.research_room_id join worker on research_room_sectary.worker_id = worker.worker_id WHERE research_room_sectary.research_room_id = ?"

	db := databaseAccess.DatabaseConn()

	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	var ret []dataStruct.Sectary
	for rows.Next() {
		var workerID int
		var workerName string
		var jobDetail string
		var researchRoomName string
		err := rows.Scan(&workerID, &workerName, &jobDetail, &researchRoomName)
		if err != nil {
			return nil, err
		}
		ret = append(ret, dataStruct.Sectary{
			WorkerID:         workerID,
			WorkerName:       workerName,
			JobDetail:        jobDetail,
			ResearchRoomName: researchRoomName,
		})
	}
	return ret, nil
}

func AddOrUpdateSectary(workerID int, researchRoomID int, info string) error {
	if exist, err := checkSectaryExist(workerID, researchRoomID); err != nil {
		return err
	} else {
		if exist {
			return updateSectary(workerID, researchRoomID, info)
		} else {
			return addSectary(workerID, researchRoomID, info)
		}
	}
}

func checkSectaryExist(workerID int, researchRoomID int) (bool, error) {
	sql := "SELECT * FROM research_room_sectary WHERE worker_id = ? AND research_room_id = ?"
	db := databaseAccess.DatabaseConn()

	stmt, err := db.Prepare(sql)
	if err != nil {
		return false, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(workerID, researchRoomID)
	if err != nil {
		return false, err
	}

	if rows.Next() {
		return true, nil
	} else {
		return false, nil
	}
}

func updateSectary(workerID int, researchRoomID int, info string) error {
	sql := "UPDATE research_room_sectary SET job_detail = ? WHERE worker_id = ? AND research_room_id = ?"
	db := databaseAccess.DatabaseConn()

	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(info, workerID, researchRoomID)
	if err != nil {
		return err
	}

	return nil
}

func addSectary(workerID int, researchRoomID int, info string) error {
	sql := "INSERT INTO research_room_sectary (worker_id, research_room_id, job_detail) VALUES (?, ?, ?)"
	db := databaseAccess.DatabaseConn()

	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(workerID, researchRoomID, info)
	if err != nil {
		return err
	}

	return nil
}

func DeleteSectary(workerID int, researchRoomID int) error {
	sql := "DELETE FROM research_room_sectary WHERE worker_id = ? AND research_room_id = ?"
	db := databaseAccess.DatabaseConn()

	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(workerID, researchRoomID)
	if err != nil {
		return err
	}

	return nil
}
