package dataUtil

import (
	"Database_Homework/databaseAccess"
	"Database_Homework/jsonHelper"
	"errors"
)

func GetWorkerByID(id int) (str *jsonHelper.JsonStr, retErr error) {
	db := databaseAccess.DatabaseConn()
	defer db.Close()

	s := "select * from worker where worker_id = ?"
	stmt, err := db.Prepare(s)
	if err != nil {
		retErr = err
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		retErr = err
		return
	}
	defer rows.Close()

	var isExist bool = false
	for rows.Next() {
		isExist = true
		var workerId int
		var workerName string
		var workerGender string
		var workerBirth string
		var workerJoinTime string
		var workerJob string
		err := rows.Scan(&workerId, &workerName, &workerGender, &workerBirth, &workerJoinTime, &workerJob)
		if err != nil {
			retErr = err
			return
		}
		str.JsonDictInit()
		str.JsonDictAddStrInt("WorkerId", workerId)
		str.JsonDictAddStrStr("WorkerName", workerName)
		str.JsonDictAddStrStr("WorkerGender", workerGender)
		str.JsonDictAddStrStr("WorkerBirth", workerBirth)
		str.JsonDictAddStrStr("WorkerJoinTime", workerJoinTime)
		str.JsonDictAddStrStr("WorkerJob", workerJob)
		str.JsonDictEnd()

	}
	if !isExist {
		retErr = errors.New("worker not exist")
		return
	} else {
		retErr = nil
		return
	}

}

func GetAllWorker() (j *jsonHelper.JsonStr, retErr error) {
	db := databaseAccess.DatabaseConn()
	defer db.Close()

	s := "select * from worker"
	stmt, err := db.Prepare(s)
	if err != nil {
		retErr = err
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query()

	if err != nil {
		retErr = err
		return
	}

	defer rows.Close()
	j = new(jsonHelper.JsonStr)
	j.JsonArrayInit()
	for rows.Next() {
		var workerId int
		var workerName string
		var workerGender string
		var workerBirth string
		var workerJoinTime string
		var workerJob string
		err := rows.Scan(&workerId, &workerName, &workerGender, &workerBirth, &workerJoinTime, &workerJob)
		if err != nil {
			retErr = err
			return
		}
		//fmt.Println(workerId, workerName, workerGender, workerBirth, workerJoinTime, workerJob)
		var temp jsonHelper.JsonStr
		temp.JsonDictInit()
		temp.JsonDictAddStrInt("WorkerId", workerId)
		temp.JsonDictAddStrStr("WorkerName", workerName)
		temp.JsonDictAddStrStr("WorkerGender", workerGender)
		temp.JsonDictAddStrStr("WorkerBirth", workerBirth)
		temp.JsonDictAddStrStr("WorkerJoinTime", workerJoinTime)
		temp.JsonDictAddStrStr("WorkerJob", workerJob)
		temp.JsonDictEnd()
		j.JsonArrayAddJson(temp)
	}
	j.JsonArrayEnd()
	return
}

func AddOrUpdateWorker(WorkerId int, WorkerName, WorkerGender, WorkerBirth, WorkerJoinTime, WorkerJob string) error {
	if WorkerId == -1 {
		// add a worker
		s := "INSERT INTO worker(worker_name, worker_gender, worker_birth, worker_join_time, worker_job) VALUES (?, ?, ?, ?, ?)"
		db := databaseAccess.DatabaseConn()
		defer db.Close()

		stmt, err := db.Prepare(s)
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(WorkerName, WorkerGender, WorkerBirth, WorkerJoinTime, WorkerJob)
		if err != nil {
			return err
		}

		return nil
	} else {
		// update
		s := "UPDATE worker SET worker_name = ?,worker_gender = ?,worker_birth = ?,worker_join_time = ?,worker_job = ? WHERE worker_id = ?"
		db := databaseAccess.DatabaseConn()
		defer db.Close()

		stmt, err := db.Prepare(s)
		if err != nil {
			return err
		}

		defer stmt.Close()

		_, err = stmt.Exec(WorkerName, WorkerGender, WorkerBirth, WorkerJoinTime, WorkerJob, WorkerId)
		if err != nil {
			return err
		}

		return nil

	}
}

func DeleteWorker(id int) error {
	s := "DELETE FROM worker WHERE worker_id = ?"
	db := databaseAccess.DatabaseConn()
	defer db.Close()

	stmt, err := db.Prepare(s)
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
