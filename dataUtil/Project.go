package dataUtil

import (
	"Database_Homework/databaseAccess"
	"strconv"
)

func GetAllOrSpecifiedProject(idname string) ([]byte, error) {
	sql := "select project_id,project_name,project_detail,project_start_time,project_end_time,project_fund,worker_name,A.project_participant_id,A.project_participant_name,T.project_participant_id,T.project_participant_name from (project natural join worker natural join project_participant AS A natural join project_participant_worker)  JOIN project_participant AS T ON project.project_supervisor_id = T.project_participant_id where project_id = ? or project_name like ?"
	if idname == "" {
		sql = "select project_id,project_name,project_detail,project_start_time,project_end_time,project_fund,worker_name,A.project_participant_id,A.project_participant_name,T.project_participant_id,T.project_participant_name from (project natural join worker natural join project_participant AS A natural join project_participant_worker)  JOIN project_participant AS T ON project.project_supervisor_id = T.project_participant_id where 1 = 1 or project_id = ? or project_name like ?"
	}

	idnamei, err := strconv.Atoi(idname)
	if err != nil {
		idnamei = 0
	}
	db := databaseAccess.DatabaseConn()
	rows, err := db.Query(sql, idnamei, "%"+idname+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ret []byte
	ret = append(ret, '[')
	for rows.Next() {
		var project_id, project_name, project_detail, project_start_time, project_end_time, project_fund, worker_name, project_participant_id, project_participant_name, project_supervisor_id, project_supervisor_name string
		err := rows.Scan(&project_id, &project_name, &project_detail, &project_start_time, &project_end_time, &project_fund, &worker_name, &project_participant_id, &project_participant_name, &project_supervisor_id, &project_supervisor_name)
		if err != nil {
			return nil, err
		}
		ret = append(ret, []byte("{\"project_id\":")...)
		ret = append(ret, []byte(project_id)...)
		ret = append(ret, []byte(",\"project_name\":\"")...)
		ret = append(ret, []byte(project_name)...)
		ret = append(ret, []byte("\",\"project_detail\":\"")...)
		ret = append(ret, []byte(project_detail)...)
		ret = append(ret, []byte("\",\"project_start_time\":\"")...)
		ret = append(ret, []byte(project_start_time)...)
		ret = append(ret, []byte("\",\"project_end_time\":\"")...)
		ret = append(ret, []byte(project_end_time)...)
		ret = append(ret, []byte("\",\"project_fund\":")...)
		ret = append(ret, []byte(project_fund)...)
		ret = append(ret, []byte(",\"worker_name\":\"")...)
		ret = append(ret, []byte(worker_name)...)
		ret = append(ret, []byte("\",\"project_participant_id\":")...)
		ret = append(ret, []byte(project_participant_id)...)
		ret = append(ret, []byte(",\"project_participant_name\":\"")...)
		ret = append(ret, []byte(project_participant_name)...)
		ret = append(ret, []byte("\",\"project_supervisor_id\":")...)
		ret = append(ret, []byte(project_supervisor_id)...)
		ret = append(ret, []byte(",\"project_supervisor_name\":\"")...)
		ret = append(ret, []byte(project_supervisor_name)...)
		ret = append(ret, []byte("\"},")...)
	}
	if len(ret) > 1 {
		ret = ret[:len(ret)-1]
	}
	ret = append(ret, ']')
	return ret, nil
}

func AddOrUpdateProject(id, name, detail, start_time, end_time, fund, worker_id, participant_id, project_supervisor_id string) error {
	sql := "INSERT INTO project (project_name,project_detail,project_start_time,project_end_time,project_fund,worker_id,project_supervisor_id) VALUES (?,?,?,?,?,?,?)"
	if id != "" {
		sql = "UPDATE project SET project_name = ?,project_detail = ?,project_start_time = ?,project_end_time = ?,project_fund = ?,worker_id = ?,project_supervisor_id = ?,project_participant_id =? WHERE project_id = ?"
	}

	db := databaseAccess.DatabaseConn()
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	idi, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	idwi, err := strconv.Atoi(worker_id)
	if err != nil {
		return err
	}
	idpi, err := strconv.Atoi(participant_id)
	if err != nil {
		return err
	}
	idpsi, err := strconv.Atoi(project_supervisor_id)
	if err != nil {
		return err
	}

	if id != "" {
		_, err = stmt.Exec(name, detail, start_time, end_time, fund, idwi, idpsi, idpi, idi)
	} else {
		_, err = stmt.Exec(name, detail, start_time, end_time, fund, idwi, idpsi)
	}

	if err != nil {
		return err
	}
	return nil
}

func DeleteProject(id string) error {
	sql := "DELETE FROM project WHERE project_id = ?"
	db := databaseAccess.DatabaseConn()
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	idi, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(idi)
	if err != nil {
		return err
	}
	return nil
}

func FindAllWorkerInProject(idname string) ([]byte, error) {
	sql := "select project.project_id,project_name,worker.worker_id,worker.worker_name from ((project join project_worker on project.project_id = project_worker.project_id) join worker on project_worker.worker_id = worker.worker_id) where project.project_id = ? or project_name like ?"
	if idname == "" {
		sql = "select project.project_id,project_name,worker.worker_id,worker.worker_name from ((project join project_worker on project.project_id = project_worker.project_id) join worker on project_worker.worker_id = worker.worker_id) where 1 = 1 or project.project_id = ? or project_name like ?"
	}

	idnamei, err := strconv.Atoi(idname)
	if err != nil {
		idnamei = 0
	}

	db := databaseAccess.DatabaseConn()
	rows, err := db.Query(sql, idnamei, "%"+idname+"%")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var ret []byte
	ret = append(ret, '[')
	for rows.Next() {
		var project_id, project_name, worker_id, worker_name string
		err := rows.Scan(&project_id, &project_name, &worker_id, &worker_name)
		if err != nil {
			return nil, err
		}
		ret = append(ret, []byte("{\"project_id\":")...)
		ret = append(ret, []byte(project_id)...)
		ret = append(ret, []byte(",\"project_name\":\"")...)
		ret = append(ret, []byte(project_name)...)
		ret = append(ret, []byte("\",\"worker_id\":")...)
		ret = append(ret, []byte(worker_id)...)
		ret = append(ret, []byte(",\"worker_name\":\"")...)
		ret = append(ret, []byte(worker_name)...)
		ret = append(ret, []byte("\"},")...)
	}

	if len(ret) > 1 {
		ret = ret[:len(ret)-1]
	}

	ret = append(ret, ']')
	return ret, nil

}

func AddProjectWorker(project_id, worker_id string) error {
	sql := "INSERT INTO project_worker (project_id,worker_id) VALUES (?,?)"
	db := databaseAccess.DatabaseConn()
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	project_idi, err := strconv.Atoi(project_id)
	if err != nil {
		return err
	}
	worker_idi, err := strconv.Atoi(worker_id)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(project_idi, worker_idi)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProjectWorker(project_id, worker_id string) error {
	sql := "DELETE FROM project_worker WHERE project_id = ? AND worker_id = ?"
	db := databaseAccess.DatabaseConn()
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	project_idi, err := strconv.Atoi(project_id)
	if err != nil {
		return err
	}
	worker_idi, err := strconv.Atoi(worker_id)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(project_idi, worker_idi)
	if err != nil {
		return err
	}
	return nil
}

func FindAllParticipantInProject(projectidname string) ([]byte, error) {
	sql := "select project.project_id,project_name,project_participant.project_participant_id,project_participant_name from (project join project_project_participant on project.project_id = project_project_participant.project_id) join project_participant on project_project_participant.project_participant_id = project_participant.project_participant_id where project.project_id = ? or project_name like ?"
	if projectidname == "" {
		sql = "select project.project_id,project_name,project_participant.project_participant_id,project_participant_name from (project join project_project_participant on project.project_id = project_project_participant.project_id) join project_participant on project_project_participant.project_participant_id = project_participant.project_participant_id where 1 = 1 or project.project_id = ? or project_name like ?"
	}

	projectidnamei, err := strconv.Atoi(projectidname)
	if err != nil {
		projectidnamei = 0
	}

	db := databaseAccess.DatabaseConn()
	rows, err := db.Query(sql, projectidnamei, "%"+projectidname+"%")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var ret []byte
	ret = append(ret, '[')
	for rows.Next() {
		var project_id, project_name, project_participant_id, project_participant_name string
		err := rows.Scan(&project_id, &project_name, &project_participant_id, &project_participant_name)
		if err != nil {
			return nil, err
		}
		ret = append(ret, []byte("{\"project_id\":")...)
		ret = append(ret, []byte(project_id)...)
		ret = append(ret, []byte(",\"project_name\":\"")...)
		ret = append(ret, []byte(project_name)...)
		ret = append(ret, []byte("\",\"project_participant_id\":")...)
		ret = append(ret, []byte(project_participant_id)...)
		ret = append(ret, []byte(",\"project_participant_name\":\"")...)
		ret = append(ret, []byte(project_participant_name)...)
		ret = append(ret, []byte("\"},")...)
	}

	if len(ret) > 1 {
		ret = ret[:len(ret)-1]
	}

	ret = append(ret, ']')
	return ret, nil
}

func AddProjectParticipant(project_id, project_participant_id string) error {
	sql := "INSERT INTO project_project_participant (project_id,project_participant_id) VALUES (?,?)"
	db := databaseAccess.DatabaseConn()
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	project_idi, err := strconv.Atoi(project_id)
	if err != nil {
		return err
	}
	project_participant_idi, err := strconv.Atoi(project_participant_id)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(project_idi, project_participant_idi)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProjectParticipant(project_id, project_participant_id string) error {
	sql := "DELETE FROM project_project_participant WHERE project_id = ? AND project_participant_id = ?"
	db := databaseAccess.DatabaseConn()
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	project_idi, err := strconv.Atoi(project_id)
	if err != nil {
		return err
	}
	project_participant_idi, err := strconv.Atoi(project_participant_id)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(project_idi, project_participant_idi)
	if err != nil {
		return err
	}
	return nil
}
