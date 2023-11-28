package dataUtil

import (
	"Database_Homework/databaseAccess"
	sql2 "database/sql"
	"strconv"
)

func GetAllOrSpecified3rdPartInfo(project_idname string) ([]byte, error) {
	sql := "SELECT project_participant_id,project_participant_name,project_participant_address,project_participant.project_participant_worker_id,project_participant_worker_telephone,project_participant_worker_mobile,project_participant_worker_email from project_participant join project_participant_worker on  project_participant.project_participant_worker_id = project_participant_worker.project_participant_worker_id"
	if project_idname == "" {
		sql += " where 1 = 1 or project_participant_id = ? or project_participant_name like ?"
	} else {
		sql += " where project_participant_id = ? or project_participant_name LIKE ?"
	}
	db := databaseAccess.DatabaseConn()
	defer db.Close()

	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}

	idnamei, err := strconv.Atoi(project_idname)
	if err != nil {
		idnamei = 0
	}

	rows, err := stmt.Query(idnamei, "%"+project_idname+"%")

	if err != nil {
		return nil, err
	}

	var ret []byte

	ret = append(ret, []byte("[")...)

	for rows.Next() {
		var project_participant_id int
		var project_participant_name string
		var project_participant_address string
		var project_participant_worker_id int
		var project_participant_worker_telephone string
		var project_participant_worker_mobile string
		var project_participant_worker_email string

		err = rows.Scan(&project_participant_id, &project_participant_name, &project_participant_address, &project_participant_worker_id, &project_participant_worker_telephone, &project_participant_worker_mobile, &project_participant_worker_email)
		if err != nil {
			return nil, err
		}

		ret = append(ret, []byte("{\"project_participant_id\": "+strconv.Itoa(project_participant_id)+", \"project_participant_name\": \""+project_participant_name+"\", \"project_participant_address\": \""+project_participant_address+"\", \"project_participant_worker_id\": "+strconv.Itoa(project_participant_worker_id)+", \"project_participant_worker_telephone\": \""+project_participant_worker_telephone+"\", \"project_participant_worker_mobile\": \""+project_participant_worker_mobile+"\", \"project_participant_worker_email\": \""+project_participant_worker_email+"\"},")...)
	}

	if len(ret) > 1 {
		ret = ret[:len(ret)-1]
	}
	ret = append(ret, []byte("]")...)

	return ret, nil
}

func AddOrUpdate3rdPartInfo(project_participant_id string, project_participant_name string, project_participant_address string, project_participant_worker_id string) error {
	sql := ""
	if project_participant_id == "" {
		sql = "INSERT INTO project_participant (project_participant_name, project_participant_address, project_participant_worker_id) VALUES (?, ?, ?)"
	} else {
		sql = "UPDATE project_participant SET project_participant_name = ?, project_participant_address = ?, project_participant_worker_id = ? WHERE project_participant_id = ?"
	}
	project_participant_idi, _ := strconv.Atoi(project_participant_id)

	db := databaseAccess.DatabaseConn()
	defer db.Close()

	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}

	if project_participant_id == "" {
		_, err = stmt.Exec(project_participant_name, project_participant_address, project_participant_worker_id)
	} else {
		_, err = stmt.Exec(project_participant_name, project_participant_address, project_participant_worker_id, project_participant_idi)
	}

	if err != nil {
		return err
	}

	return nil

}

func Delete3rdPartInfo(project_participant_id string) error {
	sql := "DELETE FROM project_participant WHERE project_participant_id = ?"
	project_participant_idi, err := strconv.Atoi(project_participant_id)
	if err != nil {
		return err
	}
	db := databaseAccess.DatabaseConn()
	defer db.Close()

	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(project_participant_idi)

	if err != nil {
		return err
	}

	return nil
}

func GetAllOrSpecified3rdPartContact(idname string) ([]byte, error) {
	sql := "select project_participant_worker.project_participant_worker_id,project_participant_worker_telephone,project_participant_worker_mobile,project_participant_worker_email,project_participant.project_participant_id,project_participant.project_participant_name,T.project_participant_id,T.project_participant_name\n    from project_participant_worker\n    left join project_participant_project_participant_worker_contact on project_participant_worker.project_participant_worker_id = project_participant_project_participant_worker_contact.project_participant_worker_id\n    left join project_participant on project_participant_worker.project_participant_worker_id = project_participant.project_participant_worker_id\n    left join project_participant as T on project_participant_project_participant_worker_contact.project_participant_id = T.project_participant_id"

	if idname == "" {
		sql += " where 1=1 or T.project_participant_id = ? or T.project_participant_name like ? or project_participant.project_participant_id = ? or project_participant.project_participant_name like ?"
	} else {
		sql += " where T.project_participant_id = ? or T.project_participant_name like ? or project_participant.project_participant_id = ? or project_participant.project_participant_name like ?"
	}

	idnamei, err := strconv.Atoi(idname)
	if err != nil {
		idnamei = 0
	}

	db := databaseAccess.DatabaseConn()
	defer db.Close()

	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(idnamei, "%"+idname+"%", idnamei, "%"+idname+"%")

	if err != nil {
		return nil, err
	}

	var ret []byte

	ret = append(ret, []byte("[")...)

	for rows.Next() {
		var project_participant_worker_id int
		var project_participant_worker_telephone string
		var project_participant_worker_mobile string
		var project_participant_worker_email string
		var project_participant_id sql2.NullInt32
		var project_participant_name sql2.NullString
		var T_project_participant_id sql2.NullInt32
		var T_project_participant_name sql2.NullString

		err = rows.Scan(&project_participant_worker_id, &project_participant_worker_telephone, &project_participant_worker_mobile, &project_participant_worker_email, &project_participant_id, &project_participant_name, &T_project_participant_id, &T_project_participant_name)
		if err != nil {
			return nil, err
		}
		currentWork := "暂无"
		if project_participant_id.Valid {
			currentWork = "负责人"
		}
		if T_project_participant_id.Valid {
			currentWork = "联系人"
		}
		if !project_participant_id.Valid {
			project_participant_id.Int32 = T_project_participant_id.Int32
			project_participant_id.Valid = T_project_participant_id.Valid

		}

		if !project_participant_name.Valid {
			project_participant_name.String = T_project_participant_name.String
			project_participant_name.Valid = T_project_participant_name.Valid
		}

		ret = append(ret, []byte("{\"project_participant_worker_id\":")...)
		ret = append(ret, []byte(strconv.Itoa(project_participant_worker_id))...)
		ret = append(ret, []byte(", \"project_participant_worker_work\": \"")...)
		ret = append(ret, []byte(""+currentWork+"\"")...)
		ret = append(ret, []byte(", \"project_participant_worker_telephone\": \"")...)
		ret = append(ret, []byte(project_participant_worker_telephone)...)
		ret = append(ret, []byte("\", \"project_participant_worker_mobile\": \"")...)
		ret = append(ret, []byte(project_participant_worker_mobile)...)
		ret = append(ret, []byte("\", \"project_participant_worker_email\": \"")...)
		ret = append(ret, []byte(project_participant_worker_email)...)
		ret = append(ret, []byte("\", \"project_participant_id\": ")...)
		if project_participant_id.Valid {
			ret = append(ret, []byte(strconv.Itoa(int(project_participant_id.Int32)))...)
		} else {
			ret = append(ret, []byte("\"暂无\"")...)
		}
		ret = append(ret, []byte(", \"project_participant_name\": \"")...)
		if project_participant_name.Valid {
			ret = append(ret, []byte(project_participant_name.String)...)
		} else {
			ret = append(ret, []byte("暂无")...)
		}
		ret = append(ret, []byte("\"},")...)
	}

	if len(ret) > 1 {
		ret = ret[:len(ret)-1]
	}

	ret = append(ret, []byte("]")...)

	return ret, nil
}

func AddOrUpdate3rdPartContact(project_participant_worker_id string, project_participant_worker_telephone string, project_participant_worker_mobile string, project_participant_worker_email string) error {
	sql := ""
	if project_participant_worker_id == "" {
		sql = "INSERT INTO project_participant_worker (project_participant_worker_telephone, project_participant_worker_mobile, project_participant_worker_email) VALUES (?, ?, ?)"
	} else {
		sql = "UPDATE project_participant_worker SET project_participant_worker_telephone = ?, project_participant_worker_mobile = ?, project_participant_worker_email = ? WHERE project_participant_worker_id = ?"
	}
	project_participant_worker_idi, _ := strconv.Atoi(project_participant_worker_id)

	db := databaseAccess.DatabaseConn()
	defer db.Close()

	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}

	if project_participant_worker_id == "" {
		_, err = stmt.Exec(project_participant_worker_telephone, project_participant_worker_mobile, project_participant_worker_email)
	} else {
		_, err = stmt.Exec(project_participant_worker_telephone, project_participant_worker_mobile, project_participant_worker_email, project_participant_worker_idi)
	}

	if err != nil {
		return err
	}

	return nil

}

func Delete3rdPartContact(project_participant_worker_id string) error {
	sql := "DELETE FROM project_participant_worker WHERE project_participant_worker_id = ?"
	project_participant_worker_idi, err := strconv.Atoi(project_participant_worker_id)
	if err != nil {
		return err
	}
	db := databaseAccess.DatabaseConn()
	defer db.Close()

	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(project_participant_worker_idi)

	if err != nil {
		return err
	}

	return nil
}

func AddContactRelation(project_participant_worker_id string, project_participant_id string) error {
	sql := "INSERT INTO project_participant_project_participant_worker_contact (project_participant_id, project_participant_worker_id) VALUES (?, ?)"
	project_participant_worker_idi, _ := strconv.Atoi(project_participant_worker_id)
	project_participant_idi, _ := strconv.Atoi(project_participant_id)

	db := databaseAccess.DatabaseConn()
	defer db.Close()

	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(project_participant_idi, project_participant_worker_idi)

	if err != nil {
		return err
	}

	return nil

}

func DeleteContactRelation(project_participant_worker_id, project_participant_id string) error {
	sql := "DELETE FROM project_participant_project_participant_worker_contact WHERE project_participant_worker_id = ? AND project_participant_id = ?"
	project_participant_worker_idi, err := strconv.Atoi(project_participant_worker_id)
	if err != nil {
		return err
	}
	project_participant_idi, err := strconv.Atoi(project_participant_id)
	if err != nil {
		return err
	}
	db := databaseAccess.DatabaseConn()
	defer db.Close()

	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(project_participant_worker_idi, project_participant_idi)

	if err != nil {
		return err
	}

	return nil
}
