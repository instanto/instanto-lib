package instanto_lib_db

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Member struct {
	Id                                  int64  `json:"id"`
	FirstName                           string `json:"first_name"`
	LastName                            string `json:"last_name"`
	Degree                              string `json:"degree"`
	YearIn                              int64  `json:"year_in"`
	YearOut                             int64  `json:"year_out"`
	Email                               string `json:"email"`
	Cv                                  string `json:"cv"`
	Photo                               string `json:"photo"`
	CreatedBy                           string `json:"created_by"`
	UpdatedBy                           string `json:"updated_by"`
	CreatedAt                           int64  `json:"created_at"`
	UpdatedAt                           int64  `json:"updated_at"`
	PrimaryStatus                       int64  `json:"primary_status"`
	RelStatusCreatedBy                  string `json:"status_created_by,omitempty"`
	RelStatusCreatedAt                  int64  `json:"status_created_at,omitempty"`
	RelPartnerCreatedBy                 string `json:"partner_created_by,omitempty"`
	RelPartnerCreatedAt                 int64  `json:"partner_created_at,omitempty"`
	RelPublicationCreatedBy             string `json:"publication_created_by,omitempty"`
	RelPublicationCreatedAt             int64  `json:"publication_created_at,omitempty"`
	RelResearchLineCreatedBy            string `json:"research_line_created_by,omitempty"`
	RelResearchLineCreatedAt            int64  `json:"research_line_created_at,omitempty"`
	RelFinancedProjectAsLeaderCreatedBy string `json:"financed_project_as_leader_created_by,omitempty"`
	RelFinancedProjectAsLeaderCreatedAt int64  `json:"financed_project_as_leader_created_at,omitempty"`
	RelFinancedProjectCreatedBy         string `json:"financed_project_created_by,omitempty"`
	RelFinancedProjectCreatedAt         int64  `json:"financed_project_created_at,omitempty"`
}

func MemberCreate(firstName, lastName, degree string, yearIn, yearOut int64, email, createdBy string, primaryStatus int64) (id int64, verr *ValidationError, err error) {
	verr = MemberValidate(firstName, lastName, degree, yearIn, yearOut, email)
	if verr != nil {
		return
	}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO member(first_name,last_name,degree,year_in,year_out,email,created_by,updated_by,created_at,updated_at,primary_status) VALUES(?,?,?,?,?,?,?,?,?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	result, err := stmt.Exec(firstName, lastName, degree, yearIn, yearOut, email, createdBy, createdBy, ts, ts, primaryStatus)
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
func MemberUpdate(id int64, firstName, lastName, degree string, yearIn, yearOut int64, email, updatedBy string, primaryStatus int64) (numRows int64, verr *ValidationError, err error) {
	verr = MemberValidate(firstName, lastName, degree, yearIn, yearOut, email)
	if verr != nil {
		return
	}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "UPDATE member SET first_name=?,last_name=?,degree=?,year_in=?,year_out=?,email=?,updated_by=?,updated_at=?,primary_status=? WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	result, err := stmt.Exec(firstName, lastName, degree, yearIn, yearOut, email, updatedBy, ts, primaryStatus, id)
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
func MemberUpdateCv(id int64, cv string, updatedBy string) (numRows int64, verr *ValidationError, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "UPDATE member SET cv=?,updated_by=?,updated_at=? WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	result, err := stmt.Exec(cv, updatedBy, ts, id)
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
func MemberUpdatePhoto(id int64, photo string, updatedBy string) (numRows int64, verr *ValidationError, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "UPDATE member SET photo=?,updated_by=?,updated_at=? WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	result, err := stmt.Exec(photo, updatedBy, ts, id)
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
func MemberDelete(id int64) (numRows int64, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM member WHERE id=?"
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
func MemberGetAll() (members []*Member, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM member"
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
		p := Member{}
		err = rows.Scan(&p.Id, &p.FirstName, &p.LastName, &p.Degree, &p.YearIn, &p.YearOut, &p.Email, &p.Cv, &p.Photo, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.PrimaryStatus)
		if err != nil {
			return
		}
		members = append(members, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func MemberGetById(id int64) (member *Member, err error) {
	member = &Member{}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM member WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&member.Id, &member.FirstName, &member.LastName, &member.Degree, &member.YearIn, &member.YearOut, &member.Email, &member.Cv, &member.Photo, &member.CreatedBy, &member.UpdatedBy, &member.CreatedAt, &member.UpdatedAt, &member.PrimaryStatus)
	if err != nil {
		return
	}
	return
}
func MemberGetByPrimaryStatus(statusId int64) (members []*Member, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM member WHERE primary_status=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(statusId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := Member{}
		err = rows.Scan(&p.Id, &p.FirstName, &p.LastName, &p.Degree, &p.YearIn, &p.YearOut, &p.Email, &p.Cv, &p.Photo, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.PrimaryStatus)
		if err != nil {
			return
		}
		members = append(members, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func MemberGetByStatus(statusId int64) (members []*Member, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT member.*,member_status.created_by,member_status.created_at FROM member_status INNER JOIN member ON member_status.member=member.id  WHERE status=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(statusId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := Member{}
		err = rows.Scan(&p.Id, &p.FirstName, &p.LastName, &p.Degree, &p.YearIn, &p.YearOut, &p.Email, &p.Cv, &p.Photo, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.PrimaryStatus, &p.RelStatusCreatedBy, &p.RelStatusCreatedAt)
		if err != nil {
			return
		}
		members = append(members, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func MemberGetByPartner(partnerId int64) (members []*Member, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT member.*,partner_member.created_by,partner_member.created_at FROM partner_member INNER JOIN member ON partner_member.member=member.id  WHERE partner=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(partnerId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := Member{}
		err = rows.Scan(&p.Id, &p.FirstName, &p.LastName, &p.Degree, &p.YearIn, &p.YearOut, &p.Email, &p.Cv, &p.Photo, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.PrimaryStatus, &p.RelPartnerCreatedBy, &p.RelPartnerCreatedAt)
		if err != nil {
			return
		}
		members = append(members, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func MemberGetByPublication(publicationId int64) (members []*Member, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT member.*,member_publication.created_by,member_publication.created_at FROM member_publication INNER JOIN member ON member_publication.member=member.id  WHERE publication=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(publicationId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := Member{}
		err = rows.Scan(&p.Id, &p.FirstName, &p.LastName, &p.Degree, &p.YearIn, &p.YearOut, &p.Email, &p.Cv, &p.Photo, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.PrimaryStatus, &p.RelPublicationCreatedBy, &p.RelPublicationCreatedAt)
		if err != nil {
			return
		}
		members = append(members, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}

func MemberGetByResearchLine(researchLineId int64) (members []*Member, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT member.*,research_line_member.created_by,research_line_member.created_at FROM research_line_member INNER JOIN member ON research_line_member.member=member.id  WHERE research_line=?"
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
		p := Member{}
		err = rows.Scan(&p.Id, &p.FirstName, &p.LastName, &p.Degree, &p.YearIn, &p.YearOut, &p.Email, &p.Cv, &p.Photo, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.PrimaryStatus, &p.RelResearchLineCreatedBy, &p.RelResearchLineCreatedAt)
		if err != nil {
			return
		}
		members = append(members, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func MemberGetByFinancedProjectAsLeader(financedProjectId int64) (members []*Member, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT member.*,financed_project_leader.created_by,financed_project_leader.created_at FROM financed_project_leader INNER JOIN member ON financed_project_leader.member=member.id  WHERE financed_project=?"
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
		p := Member{}
		err = rows.Scan(&p.Id, &p.FirstName, &p.LastName, &p.Degree, &p.YearIn, &p.YearOut, &p.Email, &p.Cv, &p.Photo, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.PrimaryStatus, &p.RelFinancedProjectAsLeaderCreatedBy, &p.RelFinancedProjectAsLeaderCreatedAt)
		if err != nil {
			return
		}
		members = append(members, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func MemberGetByFinancedProject(financedProjectId int64) (members []*Member, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT member.*,financed_project_member.created_by,financed_project_member.created_at FROM financed_project_member INNER JOIN member ON financed_project_member.member=member.id  WHERE financed_project=?"
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
		p := Member{}
		err = rows.Scan(&p.Id, &p.FirstName, &p.LastName, &p.Degree, &p.YearIn, &p.YearOut, &p.Email, &p.Cv, &p.Photo, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.PrimaryStatus, &p.RelFinancedProjectCreatedBy, &p.RelFinancedProjectCreatedAt)
		if err != nil {
			return
		}
		members = append(members, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func MemberCount() (count int64, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM member"
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
func MemberExists(id int64) (exists bool, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM member WHERE id=?"
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

func MemberAddStatus(id, statusId int64, createdBy string) (verr *ValidationError, err error) {
	member, err := MemberGetById(id)
	if err != nil {
		return
	}
	if member.PrimaryStatus == statusId {
		verr = &ValidationError{"status", "this status is already the primary"}
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
	_, err = stmt.Exec(id, statusId, createdBy, ts)
	if err != nil {
		if IsDbError1062(err) {
			verr = &ValidationError{"member", "this status has already been added"}
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
func MemberRemoveStatus(id, statusId int64) (removed bool, err error) {
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
	result, err := stmt.Exec(id, statusId)
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

func MemberGetStatuses(id int64) (statuses []*Status, err error) {
	statuses, err = StatusGetByMember(id)
	return
}
func MemberAddPartner(id, partnerId int64, createdBy string) (verr *ValidationError, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO partner_member(partner,member,created_by,created_at) VALUES(?,?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	_, err = stmt.Exec(partnerId, id, createdBy, ts)
	if err != nil {
		if IsDbError1062(err) {
			verr = &ValidationError{"partner", "this partner has already been added"}
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
func MemberRemovePartner(id, partnerId int64) (removed bool, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM partner_member WHERE partner=? AND member=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(partnerId, id)
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

func MemberGetPartners(id int64) (partners []*Partner, err error) {
	partners, err = PartnerGetByMember(id)
	return
}
func MemberAddPublication(id, publicationId int64, createdBy string) (verr *ValidationError, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO member_publication(member,publication,created_by,created_at) VALUES(?,?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	_, err = stmt.Exec(id, publicationId, createdBy, ts)
	if err != nil {
		if IsDbError1062(err) {
			verr = &ValidationError{"publication", "this publication has already been added"}
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
func MemberRemovePublication(id, publicationId int64) (removed bool, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM member_publication WHERE member=? AND publication=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(id, publicationId)
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

func MemberGetPublications(id int64) (publications []*Publication, err error) {
	publications, err = PublicationGetByMember(id)
	return
}
func MemberAddResearchLine(id, researchLineId int64, createdBy string) (verr *ValidationError, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO research_line_member(research_line,member,created_by,created_at) VALUES(?,?,?,?)"
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
func MemberRemoveResearchLine(id, researchLineId int64) (removed bool, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM research_line_member WHERE research_line=? AND member=?"
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

func MemberGetResearchLines(id int64) (researchLines []*ResearchLine, err error) {
	researchLines, err = ResearchLineGetByMember(id)
	return
}
func MemberAddFinancedProjectAsLeader(id, financedProjectId int64, createdBy string) (verr *ValidationError, err error) {
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
	_, err = stmt.Exec(financedProjectId, id, createdBy, ts)
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
func MemberRemoveFinancedProjectAsLeader(id, financedProjectId int64) (removed bool, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM financed_project_leader WHERE financed_project=? AND member=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(financedProjectId, id)
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

func MemberGetFinancedProjectsAsLeader(id int64) (financedProjects []*FinancedProject, err error) {
	financedProjects, err = FinancedProjectGetByLeader(id)
	return
}
func MemberAddFinancedProject(id, financedProjectId int64, createdBy string) (verr *ValidationError, err error) {
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
	_, err = stmt.Exec(financedProjectId, id, createdBy, ts)
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
func MemberRemoveFinancedProject(id, financedProjectId int64) (removed bool, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM financed_project_member WHERE financed_project=? AND member=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(financedProjectId, id)
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

func MemberGetFinancedProjects(id int64) (financedProjects []*FinancedProject, err error) {
	financedProjects, err = FinancedProjectGetByMember(id)
	return
}
func MemberGetStudentWorks(id int64) (studentWorks []*StudentWork, err error) {
	studentWorks, err = StudentWorkGetByAuthor(id)
	return
}

func MemberGetColumns() []string {
	columns := []string{
		"id",
		"first_name",
		"last_name",
		"degree",
		"year_in",
		"year_out",
		"email",
		"cv",
		"photo",
		"updated_by",
		"created_at",
		"updated_at",
		"primary_status",
	}
	return columns
}
func MemberValidateFirstName(firstName string) (verr *ValidationError) {
	if verr = ValidateNotEmpty("first_name", firstName); verr != nil {
		return verr
	}
	return ValidateLength("first_name", firstName, 200)
}
func MemberValidateLastName(lastName string) (verr *ValidationError) {
	if verr = ValidateNotEmpty("last_name", lastName); verr != nil {
		return verr
	}
	return ValidateLength("last_name", lastName, 200)
}
func MemberValidateDegree(degree string) (verr *ValidationError) {
	return ValidateDegree("degree", degree)
}
func MemberValidateYearIn(yearIn int64) (verr *ValidationError) {
	return ValidateIsNumber("year_in", yearIn)
}
func MemberValidateYearOut(yearOut int64) (verr *ValidationError) {
	return ValidateIsNumber("year_out", yearOut)
}
func MemberValidateEmail(email string) (verr *ValidationError) {
	return ValidateLength("email", email, 200)
}
func MemberValidate(firstName, lastName, degree string, yearIn, yearOut int64, email string) (verr *ValidationError) {
	verr = MemberValidateFirstName(firstName)
	if verr != nil {
		return
	}
	verr = MemberValidateLastName(lastName)
	if verr != nil {
		return
	}
	verr = MemberValidateDegree(degree)
	if verr != nil {
		return
	}
	verr = MemberValidateYearIn(yearIn)
	if verr != nil {
		return
	}
	verr = MemberValidateYearOut(yearOut)
	if verr != nil {
		return
	}
	verr = MemberValidateEmail(email)
	if verr != nil {
		return
	}
	return
}
