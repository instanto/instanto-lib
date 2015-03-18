package instanto_lib_db

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Status struct {
	Id                 int64  `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	CreatedBy          string `json:"created_by"`
	UpdatedBy          string `json:"updated_by"`
	CreatedAt          int64  `json:"created_at"`
	UpdatedAt          int64  `json:"updated_at"`
	RelMemberCreatedBy string `json:"member_created_by,omitempty"`
	RelMemberCreatedAt string `json:"member_created_at,omitempty"`
}

func StatusCreate(name, description, createdBy string) (id int64, verr *ValidationError, err error) {
	verr = StatusValidate(name, description)
	if verr != nil {
		return
	}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO status(name,description,created_by,updated_by,created_at,updated_at) VALUES(?,?,?,?,?,?)"
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
func StatusUpdate(id int64, name, description, updatedBy string) (numRows int64, verr *ValidationError, err error) {
	verr = StatusValidate(name, description)
	if verr != nil {
		return
	}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "UPDATE status SET name=?,description=?,updated_by=?,updated_at=? WHERE id=?"
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
func StatusDelete(id int64) (numRows int64, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM status WHERE id=?"
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
func StatusGetAll() (statuss []*Status, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM status"
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
		p := Status{}
		err = rows.Scan(&p.Id, &p.Name, &p.Description, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return
		}
		statuss = append(statuss, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func StatusGetById(id int64) (status *Status, err error) {
	status = &Status{}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM status WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&status.Id, &status.Name, &status.Description, &status.CreatedBy, &status.UpdatedBy, &status.CreatedAt, &status.UpdatedAt)
	if err != nil {
		return
	}
	return
}
func StatusGetByMember(memberId int64) (statuses []*Status, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT status.*,member_status.created_by,member_status.created_at FROM member_status INNER JOIN status ON member_status.status=status.id  WHERE member=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(memberId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := Status{}
		err = rows.Scan(&p.Id, &p.Name, &p.Description, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.RelMemberCreatedBy, &p.RelMemberCreatedAt)
		if err != nil {
			return
		}
		statuses = append(statuses, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func StatusCount() (count int64, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM status"
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
func StatusExists(id int64) (exists bool, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM status WHERE id=?"
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
func StatusAddMember(id, memberId int64, createdBy string) (verr *ValidationError, err error) {
	member, err := MemberGetById(memberId)
	if err != nil {
		return
	}
	if member.PrimaryStatus == id {
		verr = &ValidationError{"member", "this member has this status as primary"}
		return
	}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO member_status(member,status,created_by,created_at) VALUES(?,?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	_, err = stmt.Exec(memberId, id, createdBy, ts)
	if err != nil {
		if IsDbError1062(err) {
			verr = &ValidationError{"member", "this member has already been added"}
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
func StatusRemoveMember(id, memberId int64) (removed bool, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM member_status WHERE member=? AND status=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(memberId, id)
	if err != nil {
		return
	}
	numRows, err := result.RowsAffected()
	if err != nil {
		return
	}
	if numRows != 0 {
		removed = true
		return
	}
	return
}

func StatusGetMembers(id int64) (members []*Member, err error) {
	members, err = MemberGetByStatus(id)
	return
}
func StatusGetColumns() []string {
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
func StatusValidateName(name string) (verr *ValidationError) {
	if verr = ValidateNotEmpty("name", name); verr != nil {
		return verr
	}
	return ValidateLength("name", name, 200)
}
func StatusValidateDescription(description string) (verr *ValidationError) {
	return ValidateLength("description", description, 200)
}
func StatusValidate(name, description string) (verr *ValidationError) {
	verr = StatusValidateName(name)
	if verr != nil {
		return
	}
	verr = StatusValidateDescription(description)
	if verr != nil {
		return
	}
	return
}
