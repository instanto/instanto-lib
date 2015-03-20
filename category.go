package instantolib

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Category struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

func (dbp *DBProvider) CategoryCreate(name, description, createdBy string) (id int64, verr *ValidationError, err error) {
	verr = categoryValidate(name, description)
	if verr != nil {
		return
	}
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO category(name,description,created_by,updated_by,created_at,updated_at) VALUES(?,?,?,?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	result, err := stmt.Exec(name, description, createdBy, createdBy, ts, ts)
	if err != nil {
		if is, field := IsDbError1452(err); is {
			verr = &ValidationError{field, "not exist"}
			err = nil
			return
		}
		return
	}
	id, err = result.LastInsertId()
	if err != nil {
		return
	}
	return
}
func (dbp *DBProvider) CategoryUpdate(id int64, name, description, updatedBy string) (numRows int64, verr *ValidationError, err error) {
	verr = categoryValidate(name, description)
	if verr != nil {
		return
	}
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "UPDATE category SET name=?,description=?,updated_by=?,updated_at=? WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	result, err := stmt.Exec(name, description, updatedBy, ts, id)
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
func (dbp *DBProvider) CategoryDelete(id int64) (numRows int64, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM category WHERE id=?"
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
func (dbp *DBProvider) CategoryGetAll() (categorys []*Category, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM category"
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
		p := Category{}
		err = rows.Scan(&p.Id, &p.Name, &p.Description, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return
		}
		categorys = append(categorys, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func (dbp *DBProvider) CategoryGetById(id int64) (category *Category, err error) {
	category = &Category{}
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM category WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&category.Id, &category.Name, &category.Description, &category.CreatedBy, &category.UpdatedBy, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		return
	}
	return
}
func (dbp *DBProvider) CategoryCount() (count int64, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM category"
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
func (dbp *DBProvider) CategoryExists(id int64) (exists bool, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM category WHERE id=?"
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
func (dbp *DBProvider) CategoryGetColumns() []string {
	columns := []string{
		"id",
		"name",
		"description",
		"created_by",
		"updated_by",
		"created_at",
		"updated_at",
	}
	return columns
}
func categoryValidateName(name string) (verr *ValidationError) {
	if len(name) == 0 {
		verr = &ValidationError{"name", "cannot be empty"}
		return
	}
	if len(name) > 45 {
		verr = &ValidationError{"name", "length cannot be greater than 45"}
		return
	}
	return
}
func categoryValidateDescription(description string) (verr *ValidationError) {
	if len(description) > 45 {
		verr = &ValidationError{"description", "length cannot be greater than 45"}
	}
	return
}
func categoryValidate(name, description string) (verr *ValidationError) {
	verr = categoryValidateName(name)
	if verr != nil {
		return
	}
	verr = categoryValidateDescription(description)
	if verr != nil {
		return
	}
	return
}
