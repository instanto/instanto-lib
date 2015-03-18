package instanto_lib_db

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type PublicationType struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func PublicationTypeCreate(name, createdBy string) (id int64, verr *ValidationError, err error) {
	verr = PublicationTypeValidate(name)
	if verr != nil {
		return
	}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO publication_type(name,created_by,updated_by,created_at,updated_at) VALUES(?,?,?,?,?)"
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
func PublicationTypeUpdate(id int64, name, updatedBy string) (numRows int64, verr *ValidationError, err error) {
	verr = PublicationTypeValidate(name)
	if verr != nil {
		return
	}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "UPDATE publication_type SET name=?,updated_by=?,updated_at=? WHERE id=?"
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
func PublicationTypeDelete(id int64) (numRows int64, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM publication_type WHERE id=?"
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
func PublicationTypeGetAll() (publicationTypes []*PublicationType, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM publication_type"
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
		p := PublicationType{}
		err = rows.Scan(&p.Id, &p.Name, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return
		}
		publicationTypes = append(publicationTypes, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func PublicationTypeGetById(id int64) (publicationType *PublicationType, err error) {
	publicationType = &PublicationType{}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM publication_type WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&publicationType.Id, &publicationType.Name, &publicationType.CreatedBy, &publicationType.UpdatedBy, &publicationType.CreatedAt, &publicationType.UpdatedAt)
	if err != nil {
		return
	}
	return
}
func PublicationTypeCount() (count int64, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM publication_type"
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
func PublicationTypeExists(id int64) (exists bool, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM publication_type WHERE id=?"
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
func PublicationTypeGetColumns() []string {
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
func PublicationTypeValidateName(name string) (verr *ValidationError) {
	if verr = ValidateNotEmpty("name", name); verr != nil {
		return verr
	}
	return ValidateLength("name", name, 200)
}

func PublicationTypeValidate(name string) (verr *ValidationError) {
	verr = PublicationTypeValidateName(name)
	if verr != nil {
		return
	}
	return
}
