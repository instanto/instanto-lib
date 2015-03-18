package instanto_lib_db

import (
	_ "github.com/go-sql-driver/mysql"
)

type UGroup struct {
	Id          string `json:"id"`
	DisplayName string `json:"display_name"`
}

func UGroupCreate(id, displayName string) (verr *ValidationError, err error) {
	verr = UGroupValidate(displayName)
	if verr != nil {
		return
	}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO ugroup(id,display_name) VALUES(?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, displayName)
	if err != nil {
		if IsDbError1062(err) {
			verr = &ValidationError{"id", "this id is taken, use another"}
			err = nil
			return
		}
		if is, field := IsDbError1452(err); is {
			verr = &ValidationError{field, "not exist"}
			err = nil
			return
		}
		return
	}
	return
}
func UGroupUpdate(id, display_name string) (numRows int64, verr *ValidationError, err error) {
	verr = UGroupValidate(display_name)
	if verr != nil {
		return
	}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "UPDATE ugroup SET display_name=? WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(display_name, id)
	if err != nil {
		if is, field := IsDbError1452(err); is {
			verr = &ValidationError{field, "not exists"}
			err = nil
			return
		}
		return
	}
	numRows, err = result.RowsAffected()
	if err != nil {
		return
	}
	return
}
func UGroupDelete(id string) (numRows int64, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM ugroup WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(id)
	if err != nil {
		return
	}
	numRows, err = result.RowsAffected()
	if err != nil {
		return
	}
	return
}
func UGroupGetAll() (groups []*UGroup, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM ugroup"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := UGroup{}
		err = rows.Scan(&p.Id, &p.DisplayName)
		if err != nil {
			return
		}
		groups = append(groups, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func UGroupGetById(id string) (group *UGroup, err error) {
	group = &UGroup{}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM ugroup WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&group.Id, &group.DisplayName)
	if err != nil {
		return
	}
	return
}
func UGroupCount() (count int64, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM ugroup"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow().Scan(&count)
	if err != nil {
		return
	}
	return
}
func UGroupExists(id string) (exists bool, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM ugroup WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	var count int64
	err = stmt.QueryRow(id).Scan(&count)
	if err != nil {
		return
	}
	if count != 1 {
		return
	}
	exists = true
	return
}
func UGroupGetColumns() []string {
	columns := []string{
		"id",
		"display_name",
	}
	return columns
}
func UGroupValidateDisplayName(displayName string) (verr *ValidationError) {
	if verr = ValidateNotEmpty("display_name", displayName); verr != nil {
		return verr
	}
	return ValidateLength("display_name", displayName, 200)
}
func UGroupValidate(displayName string) (verr *ValidationError) {
	verr = UGroupValidateDisplayName(displayName)
	if verr != nil {
		return
	}
	return
}
