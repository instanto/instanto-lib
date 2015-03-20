package instantolib

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Publisher struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func (dbp *DBProvider) PublisherCreate(name, createdBy string) (id int64, verr *ValidationError, err error) {
	verr = publisherValidate(name)
	if verr != nil {
		return
	}
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO publisher(name,created_by,updated_by,created_at,updated_at) VALUES(?,?,?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	result, err := stmt.Exec(name, createdBy, createdBy, ts, ts)
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
func (dbp *DBProvider) PublisherUpdate(id int64, name, updatedBy string) (numRows int64, verr *ValidationError, err error) {
	verr = publisherValidate(name)
	if verr != nil {
		return
	}
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "UPDATE publisher SET name=?,updated_by=?,updated_at=? WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	result, err := stmt.Exec(name, updatedBy, ts, id)
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
func (dbp *DBProvider) PublisherDelete(id int64) (numRows int64, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM publisher WHERE id=?"
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
func (dbp *DBProvider) PublisherGetAll() (publishers []*Publisher, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM publisher"
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
		p := Publisher{}
		err = rows.Scan(&p.Id, &p.Name, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return
		}
		publishers = append(publishers, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func (dbp *DBProvider) PublisherGetById(id int64) (publisher *Publisher, err error) {
	publisher = &Publisher{}
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM publisher WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&publisher.Id, &publisher.Name, &publisher.CreatedBy, &publisher.UpdatedBy, &publisher.CreatedAt, &publisher.UpdatedAt)
	if err != nil {
		return
	}
	return
}
func (dbp *DBProvider) PublisherCount() (count int64, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM publisher"
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
func (dbp *DBProvider) PublisherExists(id int64) (exists bool, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM publisher WHERE id=?"
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
func (dbp *DBProvider) PublisherGetColumns() []string {
	columns := []string{
		"id",
		"name",
		"created_by",
		"updated_by",
		"created_at",
		"updated_at",
	}
	return columns
}
func publisherValidateName(name string) (verr *ValidationError) {
	if verr = validateNotEmpty("name", name); verr != nil {
		return verr
	}
	return validateLength("name", name, 200)
}

func publisherValidate(name string) (verr *ValidationError) {
	verr = publisherValidateName(name)
	if verr != nil {
		return
	}
	return
}
