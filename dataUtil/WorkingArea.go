package dataUtil

import (
	"Database_Homework/databaseAccess"
	"Database_Homework/jsonHelper"
	sql2 "database/sql"
)

func GetWorkingAreaByResearchRoomName(name string) (*jsonHelper.JsonStr, error) {
	sql := "SELECT working_area.working_area_id , research_room_name ,working_area_size , working_area_address FROM (working_area LEFT JOIN research_room_working_area ON working_area.working_area_id = research_room_working_area.working_area_id) LEFT JOIN research_room ON research_room.research_room_id = research_room_working_area.research_room_id WHERE research_room_name LIKE ?"
	db := databaseAccess.DatabaseConn()
	defer db.Close()

	stmt, err := db.Prepare(sql)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	name = "%" + name + "%"
	rows, err := stmt.Query(name)

	if err != nil {
		return nil, err
	}

	ret := new(jsonHelper.JsonStr)
	ret.JsonArrayInit()

	for rows.Next() {
		var workingAreaID int
		var researchRoomName sql2.NullString
		var workingAreaSize int
		var workingAreaAddress string
		err := rows.Scan(&workingAreaID, &researchRoomName, &workingAreaSize, &workingAreaAddress)
		if err != nil {
			return nil, err
		}
		var temp jsonHelper.JsonStr
		temp.JsonDictInit()
		if researchRoomName.Valid {
			temp.JsonDictAddStrInt("WorkingAreaID", workingAreaID)
			temp.JsonDictAddStrStr("ResearchRoomName", researchRoomName.String)
			temp.JsonDictAddStrInt("WorkingAreaSize", workingAreaSize)
			temp.JsonDictAddStrStr("WorkingAreaAddress", workingAreaAddress)
		} else {
			temp.JsonDictAddStrInt("WorkingAreaID", workingAreaID)
			temp.JsonDictAddStrStr("ResearchRoomName", "暂无")
			temp.JsonDictAddStrInt("WorkingAreaSize", workingAreaSize)
			temp.JsonDictAddStrStr("WorkingAreaAddress", workingAreaAddress)
		}
		temp.JsonDictEnd()
		ret.JsonArrayAddJson(temp)
	}
	ret.JsonArrayEnd()
	return ret, nil
}

func GetWorkingAreaByResearchRoomID(id int) (*jsonHelper.JsonStr, error) {
	sql := "SELECT working_area.working_area_id, research_room_name ,working_area_size , working_area_address FROM (working_area LEFT JOIN research_room_working_area ON working_area.working_area_id = research_room_working_area.working_area_id) LEFT JOIN research_room ON research_room.research_room_id = research_room_working_area.research_room_id WHERE research_room_working_area.research_room_id = ?"
	db := databaseAccess.DatabaseConn()
	defer db.Close()

	stmt, err := db.Prepare(sql)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)

	if err != nil {
		return nil, err
	}

	ret := new(jsonHelper.JsonStr)
	ret.JsonArrayInit()

	for rows.Next() {
		var workingAreaID int
		var researchRoomName sql2.NullString
		var workingAreaSize int
		var workingAreaAddress string
		err := rows.Scan(&workingAreaID, &researchRoomName, &workingAreaSize, &workingAreaAddress)
		if err != nil {
			return nil, err
		}

		var temp jsonHelper.JsonStr
		temp.JsonDictInit()
		if researchRoomName.Valid {
			temp.JsonDictAddStrInt("WorkingAreaID", workingAreaID)
			temp.JsonDictAddStrStr("ResearchRoomName", researchRoomName.String)
			temp.JsonDictAddStrInt("WorkingAreaSize", workingAreaSize)
			temp.JsonDictAddStrStr("WorkingAreaAddress", workingAreaAddress)
		} else {
			temp.JsonDictAddStrInt("WorkingAreaID", workingAreaID)
			temp.JsonDictAddStrStr("ResearchRoomName", "暂无")
			temp.JsonDictAddStrInt("WorkingAreaSize", workingAreaSize)
			temp.JsonDictAddStrStr("WorkingAreaAddress", workingAreaAddress)
		}
		temp.JsonDictEnd()
		ret.JsonArrayAddJson(temp)
	}
	ret.JsonArrayEnd()
	return ret, nil
}

func GetAllWorkingArea() (*jsonHelper.JsonStr, error) {
	sql := "SELECT working_area.working_area_id  ,research_room_name ,working_area_size , working_area_address FROM (working_area LEFT JOIN research_room_working_area ON working_area.working_area_id = research_room_working_area.working_area_id) LEFT JOIN research_room ON research_room.research_room_id = research_room_working_area.research_room_id"
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
		var workingAreaID int
		var researchRoomName sql2.NullString
		var workingAreaSize int
		var workingAreaAddress string
		err := rows.Scan(&workingAreaID, &researchRoomName, &workingAreaSize, &workingAreaAddress)
		if err != nil {
			return nil, err
		}
		var temp jsonHelper.JsonStr
		temp.JsonDictInit()
		if researchRoomName.Valid {
			temp.JsonDictAddStrInt("WorkingAreaID", workingAreaID)
			temp.JsonDictAddStrStr("ResearchRoomName", researchRoomName.String)
			temp.JsonDictAddStrInt("WorkingAreaSize", workingAreaSize)
			temp.JsonDictAddStrStr("WorkingAreaAddress", workingAreaAddress)
		} else {
			temp.JsonDictAddStrInt("WorkingAreaID", workingAreaID)
			temp.JsonDictAddStrStr("ResearchRoomName", "暂无")
			temp.JsonDictAddStrInt("WorkingAreaSize", workingAreaSize)
			temp.JsonDictAddStrStr("WorkingAreaAddress", workingAreaAddress)
		}
		temp.JsonDictEnd()
		ret.JsonArrayAddJson(temp)
	}
	ret.JsonArrayEnd()
	return ret, nil
}

func AddOrUpdateWorkingArea(id int, size int, address string) error {
	if id == -1 {
		// insert
		sql := "INSERT INTO working_area (working_area_size, working_area_address) VALUES (?, ?)"
		db := databaseAccess.DatabaseConn()
		defer db.Close()

		stmt, err := db.Prepare(sql)
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(size, address)
		if err != nil {
			return err
		}

		return nil
	} else {
		// update
		sql := "UPDATE working_area SET working_area_size = ?, working_area_address = ? WHERE working_area_id = ?"
		db := databaseAccess.DatabaseConn()
		defer db.Close()

		stmt, err := db.Prepare(sql)
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(size, address, id)
		if err != nil {
			return err
		}

		return nil
	}
}

func AddWorkingAreaForResearchRoom(workingAreaID int, researchRoomID int) error {
	sql := "INSERT INTO research_room_working_area (research_room_id, working_area_id) VALUES (?, ?)"
	db := databaseAccess.DatabaseConn()
	defer db.Close()

	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(researchRoomID, workingAreaID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteWorkingAreaForResearchRoom(workingAreaID int, researchRoomID int) error {
	sql := "DELETE FROM research_room_working_area WHERE research_room_id = ? AND working_area_id = ?"
	db := databaseAccess.DatabaseConn()
	defer db.Close()

	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(researchRoomID, workingAreaID)
	if err != nil {
		return err
	}

	return nil
}
