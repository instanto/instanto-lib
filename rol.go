package instanto_lib_db

import (
	_ "github.com/go-sql-driver/mysql"
)

type Rol struct {
	Id          string `json:"id"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
}

func RolCreate(id, displayName, description string) (verr *ValidationError, err error) {
	verr = RolValidate(displayName, description)
	if verr != nil {
		return
	}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO rol(id,display_name,description) VALUES(?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, displayName, description)
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
func RolUpdate(id, displayName, description string) (numRows int64, verr *ValidationError, err error) {
	verr = RolValidate(displayName, description)
	if verr != nil {
		return
	}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "UPDATE rol SET display_name=?,description=? WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(displayName, description, id)
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
func RolDelete(id string) (numRows int64, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM rol WHERE id=?"
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
func RolGetAll() (rols []*Rol, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM rol"
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
		p := Rol{}
		err = rows.Scan(&p.Id, &p.DisplayName, &p.Description)
		if err != nil {
			return
		}
		rols = append(rols, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func RolGetById(id string) (rol *Rol, err error) {
	rol = &Rol{}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM rol WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&rol.Id, &rol.DisplayName, &rol.Description)
	if err != nil {
		return
	}
	return
}
func RolCount() (count int64, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM rol"
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
func RolExists(id string) (exists bool, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM rol WHERE id=?"
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
func RolGetColumns() []string {
	columns := []string{
		"id",
		"display_name",
		"description",
	}
	return columns
}
func RolValidateDisplayName(displayName string) (verr *ValidationError) {
	if len(displayName) == 0 {
		verr = &ValidationError{"display_name", "cannot be empty"}
		return
	}
	if len(displayName) > 45 {
		verr = &ValidationError{"display_name", "length cannot be greater than 45"}
		return
	}
	return
}
func RolValidateDescription(description string) (verr *ValidationError) {
	if len(description) == 0 {
		verr = &ValidationError{"description", "cannot be empty"}
		return
	}
	if len(description) > 200 {
		verr = &ValidationError{"description", "length cannot be greater than 200"}
		return
	}
	return
}
func RolValidate(displayName, description string) (verr *ValidationError) {
	verr = RolValidateDisplayName(displayName)
	if verr != nil {
		return
	}
	verr = RolValidateDescription(description)
	if verr != nil {
		return
	}
	return
}
