package dataUtil

import (
	"Database_Homework/databaseAccess"
	"strconv"
)

func FindAllSubProjectInProject(idname string) ([]byte, error) {
	sql := "select project.project_id,project.project_name,sub_project_id,sub_project.worker_id,worker_name,sub_project_end_time,sub_project_fund,sub_project_tech_detail from sub_project join project on sub_project.project_id = project.project_id join worker on sub_project.worker_id = worker.worker_id"
	if idname != "" {
		sql += " where project.project_id = ? or project.project_name like ?"
	} else {
		sql += " where 1=1 or project.project_id = ? or project.project_name like ?"
	}

	db := databaseAccess.DatabaseConn()

	stmt, err := db.Prepare(sql)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	idnamei, _ := strconv.Atoi(idname)

	rows, err := stmt.Query(idnamei, "%"+idname+"%")

	if err != nil {
		return nil, err
	}

	var ret []byte
	ret = append(ret, '[')

	for rows.Next() {
		var projectID string
		var projectName string
		var subProjectID string
		var workerID string
		var workerName string
		var subProjectEndTime string
		var subProjectFund string
		var subProjectTechDetail string

		err := rows.Scan(&projectID, &projectName, &subProjectID, &workerID, &workerName, &subProjectEndTime, &subProjectFund, &subProjectTechDetail)
		if err != nil {
			return nil, err
		}

		ret = append(ret, '{')
		ret = append(ret, []byte("\"project_id\":\"")...)
		ret = append(ret, []byte(projectID)...)
		ret = append(ret, []byte("\",")...)
		ret = append(ret, []byte("\"project_name\":\"")...)
		ret = append(ret, []byte(projectName)...)
		ret = append(ret, []byte("\",")...)
		ret = append(ret, []byte("\"sub_project_id\":\"")...)
		ret = append(ret, []byte(subProjectID)...)
		ret = append(ret, []byte("\",")...)
		ret = append(ret, []byte("\"worker_id\":\"")...)
		ret = append(ret, []byte(workerID)...)
		ret = append(ret, []byte("\",")...)
		ret = append(ret, []byte("\"worker_name\":\"")...)
		ret = append(ret, []byte(workerName)...)
		ret = append(ret, []byte("\",")...)
		ret = append(ret, []byte("\"sub_project_end_time\":\"")...)
		ret = append(ret, []byte(subProjectEndTime)...)
		ret = append(ret, []byte("\",")...)
		ret = append(ret, []byte("\"sub_project_fund\":\"")...)
		ret = append(ret, []byte(subProjectFund)...)
		ret = append(ret, []byte("\",")...)
		ret = append(ret, []byte("\"sub_project_tech_detail\":\"")...)
		ret = append(ret, []byte(subProjectTechDetail)...)
		ret = append(ret, []byte("\"},")...)
	}

	if len(ret) > 1 {
		ret = ret[:len(ret)-1]
	}

	ret = append(ret, ']')

	return ret, nil

}

func AddOrUpdateSubProject(projectID string, subProjectID string, workerID string, subProjectEndTime string, subProjectFund string, subProjectTechDetail string) error {
	sql := "update sub_project set worker_id=?,sub_project_end_time=?,sub_project_fund=?,sub_project_tech_detail=? where project_id=? and sub_project_id=?"
	if subProjectID == "" {
		sql = "insert into sub_project(project_id,worker_id,sub_project_end_time,sub_project_fund,sub_project_tech_detail) values(?,?,?,?,?)"
	}

	db := databaseAccess.DatabaseConn()

	stmt, err := db.Prepare(sql)

	if err != nil {
		return err
	}

	defer stmt.Close()

	if subProjectID == "" {
		_, err = stmt.Exec(projectID, workerID, subProjectEndTime, subProjectFund, subProjectTechDetail)
	} else {
		_, err = stmt.Exec(workerID, subProjectEndTime, subProjectFund, subProjectTechDetail, projectID, subProjectID)
	}

	if err != nil {
		return err
	}

	return nil
}

func DeleteSubProject(subProjectID string) error {
	sql := "delete from sub_project where sub_project_id = ?"

	db := databaseAccess.DatabaseConn()

	stmt, err := db.Prepare(sql)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(subProjectID)

	if err != nil {
		return err
	}

	return nil
}

func FindAllSubProjectInProjectForWorker(idname string) ([]byte, error) {
	sql := "select sub_project_worker.sub_project_id,sub_project_tech_detail,sub_project_worker.worker_id,worker_name,join_time,sub_project_worker_fund,workload from sub_project_worker join sub_project on sub_project_worker.sub_project_id = sub_project.sub_project_id join worker on sub_project_worker.worker_id = worker.worker_id"
	if idname != "" {
		sql += " where sub_project_worker.sub_project_id = ? or worker_name like ?"
	} else {
		sql += " where 1=1 or sub_project_worker.sub_project_id = ? or worker_name like ?"
	}

	db := databaseAccess.DatabaseConn()

	stmt, err := db.Prepare(sql)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	idnamei, _ := strconv.Atoi(idname)

	rows, err := stmt.Query(idnamei, "%"+idname+"%")

	if err != nil {
		return nil, err
	}

	var ret []byte
	ret = append(ret, '[')

	for rows.Next() {
		var subProjectID string
		var subProjectTechDetail string
		var workerID string
		var workerName string
		var joinTime string
		var subProjectWorkerFund string
		var workload string

		err := rows.Scan(&subProjectID, &subProjectTechDetail, &workerID, &workerName, &joinTime, &subProjectWorkerFund, &workload)
		if err != nil {
			return nil, err
		}

		ret = append(ret, '{')
		ret = append(ret, []byte("\"sub_project_id\":\"")...)
		ret = append(ret, []byte(subProjectID)...)
		ret = append(ret, []byte("\",")...)
		ret = append(ret, []byte("\"sub_project_tech_detail\":\"")...)
		ret = append(ret, []byte(subProjectTechDetail)...)
		ret = append(ret, []byte("\",")...)
		ret = append(ret, []byte("\"worker_id\":\"")...)
		ret = append(ret, []byte(workerID)...)
		ret = append(ret, []byte("\",")...)
		ret = append(ret, []byte("\"worker_name\":\"")...)
		ret = append(ret, []byte(workerName)...)
		ret = append(ret, []byte("\",")...)
		ret = append(ret, []byte("\"join_time\":\"")...)
		ret = append(ret, []byte(joinTime)...)
		ret = append(ret, []byte("\",")...)
		ret = append(ret, []byte("\"sub_project_worker_fund\":\"")...)
		ret = append(ret, []byte(subProjectWorkerFund)...)
		ret = append(ret, []byte("\",")...)
		ret = append(ret, []byte("\"workload\":\"")...)
		ret = append(ret, []byte(workload)...)
		ret = append(ret, []byte("\"},")...)
	}

	if len(ret) > 1 {
		ret = ret[:len(ret)-1]
	}

	ret = append(ret, ']')

	return ret, nil
}

func AddOrUpdateSubProjectWorker(subProjectID string, workerID string, joinTime string, subProjectWorkerFund string, workload string) error {
	sql := "replace into sub_project_worker(sub_project_id,worker_id,join_time,sub_project_worker_fund,workload) values(?,?,?,?,?)"

	db := databaseAccess.DatabaseConn()

	stmt, err := db.Prepare(sql)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(subProjectID, workerID, joinTime, subProjectWorkerFund, workload)

	if err != nil {
		return err
	}

	return nil
}

func DeleteSubProjectWorker(subProjectID string, workerID string) error {
	sql := "delete from sub_project_worker where sub_project_id = ? and worker_id = ?"

	db := databaseAccess.DatabaseConn()

	stmt, err := db.Prepare(sql)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(subProjectID, workerID)

	if err != nil {
		return err
	}

	return nil
}
