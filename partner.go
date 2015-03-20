package instantolib

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Partner struct {
	Id                       int64  `json:"id"`
	Name                     string `json:"name"`
	Web                      string `json:"web"`
	Logo                     string `json:"logo"`
	SameDepartment           bool   `json:"same_department"`
	Scope                    string `json:"scope"`
	CreatedBy                string `json:"created_by"`
	UpdatedBy                string `json:"updated_by"`
	CreatedAt                int64  `json:"created_at"`
	UpdatedAt                int64  `json:"updated_at"`
	RelMemberCreatedBy       string `json:"member_created_by,omitempty"`
	RelMemberCreatedAt       int64  `json:"member_created_at,omitempty"`
	RelResearchLineCreatedBy string `json:"research_line_created_by,omitempty"`
	RelResearchLineCreatedAt int64  `json:"research_line_created_at,omitempty"`
}

func (dbp *DBProvider) PartnerCreate(name, web string, sameDepartment bool, scope string, createdBy string) (id int64, verr *ValidationError, err error) {
	verr = partnerValidate(name, web, sameDepartment, scope)
	if verr != nil {
		return
	}
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO partner(name,web,same_department,scope,created_by,updated_by,created_at,updated_at) VALUES(?,?,?,?,?,?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	result, err := stmt.Exec(name, web, sameDepartment, scope, createdBy, createdBy, ts, ts)
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
func (dbp *DBProvider) PartnerUpdate(id int64, name, web string, sameDepartment bool, scope string, updatedBy string) (numRows int64, verr *ValidationError, err error) {
	verr = partnerValidate(name, web, sameDepartment, scope)
	if verr != nil {
		return
	}
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "UPDATE partner SET name=?,web=?,same_department=?,scope=?,updated_by=?,updated_at=? WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	result, err := stmt.Exec(name, web, sameDepartment, scope, updatedBy, ts, id)
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
func (dbp *DBProvider) PartnerUpdateLogo(id int64, logo string, updatedBy string) (numRows int64, verr *ValidationError, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "UPDATE partner SET logo=?,updated_by=?,updated_at=? WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	result, err := stmt.Exec(logo, updatedBy, ts, id)
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
func (dbp *DBProvider) PartnerDelete(id int64) (numRows int64, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM partner WHERE id=?"
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
func (dbp *DBProvider) PartnerGetAll() (partners []*Partner, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM partner"
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
		p := Partner{}
		err = rows.Scan(&p.Id, &p.Name, &p.Web, &p.Logo, &p.SameDepartment, &p.Scope, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return
		}
		partners = append(partners, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func (dbp *DBProvider) PartnerGetById(id int64) (partner *Partner, err error) {
	partner = &Partner{}
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM partner WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&partner.Id, &partner.Name, &partner.Web, &partner.Logo, &partner.SameDepartment, &partner.Scope, &partner.CreatedBy, &partner.UpdatedBy, &partner.CreatedAt, &partner.UpdatedAt)
	if err != nil {
		return
	}
	return
}
func (dbp *DBProvider) PartnerGetByMember(memberId int64) (partners []*Partner, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT partner.*,partner_member.created_by,partner_member.created_at FROM partner_member INNER JOIN partner ON partner_member.partner=partner.id  WHERE member=?"
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
		p := Partner{}
		err = rows.Scan(&p.Id, &p.Name, &p.Web, &p.Logo, &p.SameDepartment, &p.Scope, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.RelMemberCreatedBy, &p.RelMemberCreatedAt)
		if err != nil {
			return
		}
		partners = append(partners, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func (dbp *DBProvider) PartnerGetByResearchLine(researchLineId int64) (partners []*Partner, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT partner.*,research_line_partner.created_by,research_line_partner.created_at FROM research_line_partner INNER JOIN partner ON research_line_partner.partner=partner.id  WHERE research_line=?"
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
		p := Partner{}
		err = rows.Scan(&p.Id, &p.Name, &p.Web, &p.Logo, &p.SameDepartment, &p.Scope, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.RelResearchLineCreatedBy, &p.RelResearchLineCreatedAt)
		if err != nil {
			return
		}
		partners = append(partners, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func (dbp *DBProvider) PartnerCount() (count int64, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM partner"
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
func (dbp *DBProvider) PartnerExists(id int64) (exists bool, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM partner WHERE id=?"
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
func (dbp *DBProvider) PartnerAddMember(id, memberId int64, createdBy string) (verr *ValidationError, err error) {
	db, err := dbp.getDB()
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
func (dbp *DBProvider) PartnerRemoveMember(id, memberId int64) (removed bool, err error) {
	db, err := dbp.getDB()
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
	result, err := stmt.Exec(id, memberId)
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

func (dbp *DBProvider) PartnerGetMembers(id int64) (members []*Member, err error) {
	members, err = dbp.MemberGetByPartner(id)
	return
}

func (dbp *DBProvider) PartnerAddResearchLine(id, researchLineId int64, createdBy string) (verr *ValidationError, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO research_line_partner(research_line,partner,created_by,created_at) VALUES(?,?,?,?)"
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
func (dbp *DBProvider) PartnerRemoveResearchLine(id, researchLineId int64) (removed bool, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM research_line_partner WHERE research_line=? AND partner=?"
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

func (dbp *DBProvider) PartnerGetResearchLines(id int64) (researchLines []*ResearchLine, err error) {
	researchLines, err = dbp.ResearchLineGetByPartner(id)
	return
}
func (dbp *DBProvider) PartnerGetColumns() []string {
	columns := []string{
		"id",
		"name",
		"web",
		"logo",
		"same_department",
		"scope",
		"created_by",
		"updated_by",
		"created_at",
		"updated_at",
	}
	return columns
}
func partnerValidateName(name string) (verr *ValidationError) {
	if verr = validateNotEmpty("name", name); verr != nil {
		return verr
	}
	return validateLength("name", name, 200)
}
func partnerValidateWeb(web string) (verr *ValidationError) {
	return validateLength("web", web, 200)
}
func partnerValidateLogo(logo string) (verr *ValidationError) {
	return validateLength("logo", logo, 200)
}
func partnerValidateScope(scope string) (verr *ValidationError) {
	return validateScope("scope", scope)
}
func partnerValidate(name, web string, sameDepartment bool, scope string) (verr *ValidationError) {
	verr = partnerValidateName(name)
	if verr != nil {
		return
	}
	verr = partnerValidateWeb(web)
	if verr != nil {
		return
	}
	verr = partnerValidateScope(scope)
	if verr != nil {
		return
	}
	return
}
