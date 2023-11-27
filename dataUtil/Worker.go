package dataUtil

import (
	"Database_Homework/databaseAccess"
	"errors"
)
import "Database_Homework/dataStruct"

func GetWorkerByID(id int) (w dataStruct.Worker, retErr error) {
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
		//fmt.Println(workerId, workerName, workerGender, workerBirth, workerJoinTime, workerJob)
		w = dataStruct.Worker{
			WorkerId:       workerId,
			WorkerName:     workerName,
			WorkerGender:   workerGender,
			WorkerBirth:    workerBirth,
			WorkerJoinTime: workerJoinTime,
			WorkerJob:      workerJob,
		}
	}
	if !isExist {
		retErr = errors.New("worker not exist")
		return
	} else {
		return
	}

}

func GetAllWorker() (w []dataStruct.Worker, retErr error) {
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
		w = append(w, dataStruct.Worker{
			WorkerId:       workerId,
			WorkerName:     workerName,
			WorkerGender:   workerGender,
			WorkerBirth:    workerBirth,
			WorkerJoinTime: workerJoinTime,
			WorkerJob:      workerJob,
		})
	}
	return
}

func AddOrUpdateWorker(w dataStruct.Worker) error {
	if w.WorkerId == -1 {
		// add a worker
		s := "INSERT INTO worker(worker_name, worker_gender, worker_birth, worker_join_time, worker_job) VALUES (?, ?, ?, ?, ?)"
		db := databaseAccess.DatabaseConn()
		defer db.Close()

		stmt, err := db.Prepare(s)
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(w.WorkerName, w.WorkerGender, w.WorkerBirth, w.WorkerJoinTime, w.WorkerJob)
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

		_, err = stmt.Exec(w.WorkerName, w.WorkerGender, w.WorkerBirth, w.WorkerJoinTime, w.WorkerJob, w.WorkerId)
		if err != nil {
			return err
		}

		return nil

	}
}
