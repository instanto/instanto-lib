package instanto_lib_db

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type FundingBody struct {
	Id                          int64  `json:"id"`
	Name                        string `json:"name"`
	Web                         string `json:"web"`
	Scope                       string `json:"scope"`
	CreatedBy                   string `json:"created_by"`
	UpdatedBy                   string `json:"updated_by"`
	CreatedAt                   int64  `json:"created_at"`
	UpdatedAt                   int64  `json:"updated_at"`
	RelFinancedProjectRecord    string `json:"financed_project_record,omitempty"`
	RelFinancedProjectCreatedBy string `json:"financed_project_created_by,omitempty"`
	RelFinancedProjectUpdatedBy string `json:"financed_project_updated_by,omitempty"`
	RelFinancedProjectCreatedAt int64  `json:"financed_project_created_at,omitempty"`
	RelFinancedProjectUpdatedAt int64  `json:"financed_project_updated_at,omitempty"`
}

func FundingBodyCreate(name, web, scope string, createdBy string) (id int64, verr *ValidationError, err error) {
	verr = FundingBodyValidate(name, web, scope)
	if verr != nil {
		return
	}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO funding_body(name,web,scope,created_by,updated_by,created_at,updated_at) VALUES(?,?,?,?,?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	result, err := stmt.Exec(name, web, scope, createdBy, createdBy, ts, ts)
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
func FundingBodyUpdate(id int64, name, web, scope string, updatedBy string) (numRows int64, verr *ValidationError, err error) {
	verr = FundingBodyValidate(name, web, scope)
	if verr != nil {
		return
	}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "UPDATE funding_body SET name=?,web=?,scope=?,updated_by=?,updated_at=? WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	result, err := stmt.Exec(name, web, scope, updatedBy, ts, id)
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
func FundingBodyDelete(id int64) (numRows int64, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM funding_body WHERE id=?"
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
func FundingBodyGetAll() (fundingBodys []*FundingBody, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM funding_body"
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
		p := FundingBody{}
		err = rows.Scan(&p.Id, &p.Name, &p.Web, &p.Scope, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return
		}
		fundingBodys = append(fundingBodys, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func FundingBodyGetById(id int64) (fundingBody *FundingBody, err error) {
	fundingBody = &FundingBody{}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM funding_body WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&fundingBody.Id, &fundingBody.Name, &fundingBody.Web, &fundingBody.Scope, &fundingBody.CreatedBy, &fundingBody.UpdatedBy, &fundingBody.CreatedAt, &fundingBody.UpdatedAt)
	if err != nil {
		return
	}
	return
}
func FundingBodyGetByFinancedProject(financedProjectId int64) (fundingBodies []*FundingBody, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT funding_body.*,funding_body_financed_project.record,funding_body_financed_project.created_by,funding_body_financed_project.updated_by,funding_body_financed_project.created_at,funding_body_financed_project.updated_at  FROM funding_body_financed_project INNER JOIN funding_body ON funding_body_financed_project.funding_body=funding_body.id  WHERE financed_project=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(financedProjectId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := FundingBody{}
		err = rows.Scan(&p.Id, &p.Name, &p.Web, &p.Scope, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.RelFinancedProjectRecord, &p.RelFinancedProjectCreatedBy, &p.RelFinancedProjectUpdatedBy, &p.RelFinancedProjectCreatedAt, &p.RelFinancedProjectUpdatedAt)
		if err != nil {
			return
		}
		fundingBodies = append(fundingBodies, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func FundingBodyCount() (count int64, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM funding_body"
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
func FundingBodyExists(id int64) (exists bool, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM funding_body WHERE id=?"
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

func FundingBodyAddFinancedProject(id, financedProjectId int64, record, createdBy string) (verr *ValidationError, err error) {
	verr = FundingBodyValidateRecord(record)
	if verr != nil {
		return
	}
	financedProject, err := FinancedProjectGetById(financedProjectId)
	if err != nil {
		return
	}
	if financedProject.PrimaryFundingBody == id {
		verr = &ValidationError{"financed_project", "this financed project has this funding body as primary"}
		return
	}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO funding_body_financed_project(funding_body,financed_project,record,created_by,updated_by,created_at,updated_at) VALUES(?,?,?,?,?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	_, err = stmt.Exec(id, financedProjectId, record, createdBy, createdBy, ts, ts)
	if err != nil {
		if IsDbError1062(err) {
			verr = &ValidationError{"financed_project", "this financed project has already been added"}
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
func FundingBodyRemoveFinancedProject(id, financedProjectId int64) (removed bool, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM funding_body_financed_project WHERE funding_body=? AND financed_project=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(id, financedProjectId)
	if err != nil {
		return
	}
	numRows, err := result.RowsAffected()
	if err != nil {
		return
	}
	if numRows != 0 {
		removed = true
	}
	return
}

func FundingBodyGetFinancedProjects(id int64) (financedProjects []*FinancedProject, err error) {
	financedProjects, err = FinancedProjectGetByFundingBody(id)
	return
}

func FundingBodyGetColumns() []string {
	columns := []string{
		"id",
		"name",
		"web",
		"scope",
		"created_by",
		"updated_by",
		"created_at",
		"updated_at",
	}
	return columns
}
func FundingBodyValidateName(name string) (verr *ValidationError) {
	if verr = ValidateNotEmpty("name", name); verr != nil {
		return verr
	}
	return ValidateLength("name", name, 200)
}
func FundingBodyValidateWeb(web string) (verr *ValidationError) {
	return ValidateLength("web", web, 200)
}
func FundingBodyValidateScope(scope string) (verr *ValidationError) {
	return ValidateScope("scope", scope)
}
func FundingBodyValidateRecord(record string) (err *ValidationError) {
	return ValidateLength("record", record, 200)
}
func FundingBodyValidate(name, web, scope string) (verr *ValidationError) {
	verr = FundingBodyValidateName(name)
	if verr != nil {
		return
	}
	verr = FundingBodyValidateWeb(web)
	if verr != nil {
		return
	}
	verr = FundingBodyValidateScope(scope)
	if verr != nil {
		return
	}
	return
}
