package instanto_lib_db

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type FinancedProject struct {
	Id                         int64  `json:"id"`
	Title                      string `json:"title"`
	Started                    int64  `json:"started"`
	Ended                      int64  `json:"ended"`
	Budget                     int64  `json:"budget"`
	Scope                      string `json:"scope"`
	CreatedBy                  string `json:"created_by"`
	UpdatedBy                  string `json:"updated_by"`
	CreatedAt                  int64  `json:"created_at"`
	UpdatedAt                  int64  `json:"updated_at"`
	PrimaryFundingBody         int64  `json:"primary_funding_body"`
	PrimaryRecord              string `json:"primary_record"`
	PrimaryLeader              int64  `json:"primary_leader"`
	RelFundingBodyRecord       string `json:"funding_body_record,omitempty"`
	RelFundingBodyCreatedBy    string `json:"funding_body_created_by,omitempty"`
	RelFundingBodyUpdatedBy    string `json:"funding_body_updated_by,omitempty"`
	RelFundingBodyCreatedAt    int64  `json:"funding_body_created_at,omitempty"`
	RelFundingBodyUpdatedAt    int64  `json:"funding_body_updated_at,omitempty"`
	RelMemberAsLeaderCreatedBy string `json:"member_as_leader_created_by,omitempty"`
	RelMemberAsLeaderCreatedAt int64  `json:"member_as_leader_created_at,omitempty"`
	RelMemberCreatedBy         string `json:"member_created_by,omitempty"`
	RelMemberCreatedAt         int64  `json:"member_created_at,omitempty"`
	RelResearchLineCreatedby   string `json:"research_line_created_by,omitempty"`
	RelResearchLineCreatedAt   int64  `json:"research_line_created_at,omitempty"`
}

func FinancedProjectCreate(title string, started, ended, budget int64, scope string, createdBy string, primaryFundingBody int64, primaryRecord string, primaryLeader int64) (id int64, verr *ValidationError, err error) {
	verr = FinancedProjectValidate(title, started, ended, budget, scope)
	if verr != nil {
		return
	}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO financed_project(title,started,ended,budget,scope,created_by,updated_by,created_at,updated_at,primary_funding_body,primary_record,primary_leader) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	result, err := stmt.Exec(title, started, ended, budget, scope, createdBy, createdBy, ts, ts, primaryFundingBody, primaryRecord, primaryLeader)
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
func FinancedProjectUpdate(id int64, title string, started, ended, budget int64, scope string, updatedBy string, primaryFundingBody int64, primaryRecord string, primaryLeader int64) (numRows int64, verr *ValidationError, err error) {
	verr = FinancedProjectValidate(title, started, ended, budget, scope)
	if verr != nil {
		return
	}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "UPDATE financed_project SET title=?,started=?,ended=?,budget=?,scope=?,updated_by=?,updated_at=?,primary_funding_body=?,primary_record=?,primary_leader=? WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	result, err := stmt.Exec(title, started, ended, budget, scope, updatedBy, ts, primaryFundingBody, primaryRecord, primaryLeader, id)
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
func FinancedProjectDelete(id int64) (numRows int64, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM financed_project WHERE id=?"
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
func FinancedProjectGetAll() (financedProjects []*FinancedProject, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM financed_project"
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
		p := FinancedProject{}
		err = rows.Scan(&p.Id, &p.Title, &p.Started, &p.Ended, &p.Budget, &p.Scope, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.PrimaryFundingBody, &p.PrimaryRecord, &p.PrimaryLeader)
		if err != nil {
			return
		}
		financedProjects = append(financedProjects, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func FinancedProjectGetById(id int64) (financedProject *FinancedProject, err error) {
	financedProject = &FinancedProject{}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM financed_project WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&financedProject.Id, &financedProject.Title, &financedProject.Started, &financedProject.Ended, &financedProject.Budget, &financedProject.Scope, &financedProject.CreatedBy, &financedProject.UpdatedBy, &financedProject.CreatedAt, &financedProject.UpdatedAt, &financedProject.PrimaryFundingBody, &financedProject.PrimaryRecord, &financedProject.PrimaryLeader)
	if err != nil {
		return
	}
	return
}
func FinancedProjectGetByPrimaryFundingBody(fundingBodyId int64) (financedProjects []*FinancedProject, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM financed_project WHERE primary_funding_body=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(fundingBodyId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := FinancedProject{}
		err = rows.Scan(&p.Id, &p.Title, &p.Started, &p.Ended, &p.Budget, &p.Scope, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.PrimaryFundingBody, &p.PrimaryRecord, &p.PrimaryLeader)
		if err != nil {
			return
		}
		financedProjects = append(financedProjects, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func FinancedProjectGetByPrimaryLeader(leaderId int64) (financedProjects []*FinancedProject, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM financed_project WHERE primary_leader=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(leaderId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := FinancedProject{}
		err = rows.Scan(&p.Id, &p.Title, &p.Started, &p.Ended, &p.Budget, &p.Scope, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.PrimaryFundingBody, &p.PrimaryRecord, &p.PrimaryLeader)
		if err != nil {
			return
		}
		financedProjects = append(financedProjects, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func FinancedProjectGetByFundingBody(fundingBodyId int64) (financedProjects []*FinancedProject, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT financed_project.*,funding_body_financed_project.record,funding_body_financed_project.created_by,funding_body_financed_project.updated_by,funding_body_financed_project.created_at,funding_body_financed_project.updated_at  FROM funding_body_financed_project INNER JOIN financed_project ON funding_body_financed_project.financed_project=financed_project.id  WHERE funding_body=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(fundingBodyId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := FinancedProject{}
		err = rows.Scan(&p.Id, &p.Title, &p.Started, &p.Ended, &p.Budget, &p.Scope, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.PrimaryFundingBody, &p.PrimaryRecord, &p.PrimaryLeader, &p.RelFundingBodyRecord, &p.RelFundingBodyCreatedBy, &p.RelFundingBodyUpdatedBy, &p.RelFundingBodyCreatedAt, &p.RelFundingBodyUpdatedAt)
		if err != nil {
			return
		}
		financedProjects = append(financedProjects, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func FinancedProjectGetByLeader(leaderId int64) (financedProjects []*FinancedProject, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT financed_project.*,financed_project_leader.created_by,financed_project_leader.created_at FROM financed_project_leader INNER JOIN financed_project ON financed_project_leader.financed_project=financed_project.id  WHERE member=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(leaderId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := FinancedProject{}
		err = rows.Scan(&p.Id, &p.Title, &p.Started, &p.Ended, &p.Budget, &p.Scope, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.PrimaryFundingBody, &p.PrimaryRecord, &p.PrimaryLeader, &p.RelMemberAsLeaderCreatedBy, &p.RelMemberAsLeaderCreatedAt)
		if err != nil {
			return
		}
		financedProjects = append(financedProjects, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func FinancedProjectGetByMember(memberId int64) (financedProjects []*FinancedProject, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT financed_project.*,financed_project_member.created_by,financed_project_member.created_at FROM financed_project_member INNER JOIN financed_project ON financed_project_member.financed_project=financed_project.id  WHERE member=?"
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
		p := FinancedProject{}
		err = rows.Scan(&p.Id, &p.Title, &p.Started, &p.Ended, &p.Budget, &p.Scope, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.PrimaryFundingBody, &p.PrimaryRecord, &p.PrimaryLeader, &p.RelMemberCreatedBy, &p.RelMemberCreatedAt)
		if err != nil {
			return
		}
		financedProjects = append(financedProjects, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func FinancedProjectGetByResearchLine(researchLineId int64) (financedProjects []*FinancedProject, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT financed_project.*,research_line_financed_project.created_by,research_line_financed_project.created_at FROM research_line_financed_project INNER JOIN financed_project ON research_line_financed_project.financed_project=financed_project.id  WHERE research_line=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(researchLineId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := FinancedProject{}
		err = rows.Scan(&p.Id, &p.Title, &p.Started, &p.Ended, &p.Budget, &p.Scope, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.PrimaryFundingBody, &p.PrimaryRecord, &p.PrimaryLeader, &p.RelResearchLineCreatedby, &p.RelResearchLineCreatedAt)
		if err != nil {
			return
		}
		financedProjects = append(financedProjects, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func FinancedProjectCount() (count int64, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM financed_project"
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
func FinancedProjectExists(id int64) (exists bool, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM financed_project WHERE id=?"
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

func FinancedProjectAddFundingBody(id, fundingBodyId int64, record, createdBy string) (verr *ValidationError, err error) {
	verr = FinancedProjectValidateRecord(record)
	if verr != nil {
		return
	}
	financedProject, err := FinancedProjectGetById(id)
	if err != nil {
		return
	}
	if financedProject.PrimaryFundingBody == fundingBodyId {
		verr = &ValidationError{"funding_body", "this funding body is already the primary"}
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
	_, err = stmt.Exec(fundingBodyId, id, record, createdBy, createdBy, ts, ts)
	if err != nil {
		if IsDbError1062(err) {
			verr = &ValidationError{"funding_body", "this funding body has already been added"}
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
func FinancedProjectRemoveFundingBody(id, fundingBodyId int64) (removed bool, err error) {
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
	result, err := stmt.Exec(fundingBodyId, id)
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

func FinancedProjectGetFundingBodies(id int64) (fundingBodies []*FundingBody, err error) {
	fundingBodies, err = FundingBodyGetByFinancedProject(id)
	return
}
func FinancedProjectAddLeader(id, leaderId int64, createdBy string) (verr *ValidationError, err error) {
	financedProject, err := FinancedProjectGetById(id)
	if err != nil {
		return
	}
	if financedProject.PrimaryLeader == leaderId {
		verr = &ValidationError{"leader", "this leader is already the primary"}
		return
	}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO financed_project_leader(financed_project,member,created_by,created_at) VALUES(?,?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	_, err = stmt.Exec(id, leaderId, createdBy, ts)
	if err != nil {
		if IsDbError1062(err) {
			verr = &ValidationError{"leader", "this leader has already been added"}
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
func FinancedProjectRemoveLeader(id, leaderId int64) (removed bool, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM financed_project_leader WHERE member=? AND financed_project=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(leaderId, id)
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

func FinancedProjectGetLeaders(id int64) (leaders []*Member, err error) {
	leaders, err = MemberGetByFinancedProjectAsLeader(id)
	return
}

func FinancedProjectAddMember(id, memberId int64, createdBy string) (verr *ValidationError, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO financed_project_member(financed_project,member,created_by,created_at) VALUES(?,?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	_, err = stmt.Exec(id, memberId, createdBy, ts)
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
func FinancedProjectRemoveMember(id, memberId int64) (removed bool, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM financed_project_member WHERE member=? AND financed_project=?"
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
	}
	return
}

func FinancedProjectGetMembers(id int64) (members []*Member, err error) {
	members, err = MemberGetByFinancedProject(id)
	return
}
func FinancedProjectAddResearchLine(id, researchLineId int64, createdBy string) (verr *ValidationError, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO research_line_financed_project(research_line,financed_project,created_by,created_at) VALUES(?,?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	_, err = stmt.Exec(researchLineId, id, createdBy, ts)
	if err != nil {
		if IsDbError1062(err) {
			verr = &ValidationError{"research_line", "this research line has already been added"}
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
func FinancedProjectRemoveResearchLine(id, researchLineId int64) (removed bool, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM research_line_financed_project WHERE research_line=? AND financed_project=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(researchLineId, id)
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

func FinancedProjectGetResearchLines(id int64) (researchLines []*ResearchLine, err error) {
	researchLines, err = ResearchLineGetByFinancedProject(id)
	return
}
func FinancedProjectGetColumns() []string {
	columns := []string{
		"id",
		"title",
		"started",
		"ended",
		"budget",
		"scope",
		"created_by",
		"updated_by",
		"created_at",
		"updated_at",
		"primary_funding_body",
		"primary_record",
		"primary_leader",
	}
	return columns
}
func FinancedProjectValidateTitle(title string) (verr *ValidationError) {
	if verr = ValidateNotEmpty("title", title); verr != nil {
		return verr
	}
	return ValidateLength("title", title, 200)
}
func FinancedProjectValidateStartedAndEnded(started, ended int64) (verr *ValidationError) {
	if verr = ValidateIsNumber("started", started); verr != nil {
		return verr
	}
	return nil
}
func FinancedProjectValidateBudget(budget int64) (verr *ValidationError) {
	return ValidateIsNumber("budget", budget)
}
func FinancedProjectValidateScope(scope string) (verr *ValidationError) {
	return ValidateScope("scope", scope)
}
func FinancedProjectValidateRecord(record string) (err *ValidationError) {
	return ValidateLength("record", record, 200)
}
func FinancedProjectValidate(title string, started, ended, budget int64, scope string) (verr *ValidationError) {
	verr = FinancedProjectValidateTitle(title)
	if verr != nil {
		return
	}
	verr = FinancedProjectValidateStartedAndEnded(started, ended)
	if verr != nil {
		return
	}
	verr = FinancedProjectValidateBudget(budget)
	if verr != nil {
		return
	}
	verr = FinancedProjectValidateScope(scope)
	if verr != nil {
		return
	}
	return
}
