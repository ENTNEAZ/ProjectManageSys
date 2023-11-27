package dataUtil

import (
	"Database_Homework/dataStruct"
	"Database_Homework/databaseAccess"
)

func GetAllSectary() ([]dataStruct.Sectary, error) {
	sql := "SELECT worker_id,worker_name,job_detail,research_room_name FROM research_room NATURAL JOIN research_room_sectary NATURAL JOIN worker"
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
	sql := "SELECT worker_id,worker_name,job_detail,research_room_name FROM research_room NATURAL JOIN research_room_sectary NATURAL JOIN worker WHERE research_room_id = ?"
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
