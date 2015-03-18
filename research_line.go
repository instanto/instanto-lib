package instanto_lib_db

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type ResearchLine struct {
	Id                          int64  `json:"id"`
	Title                       string `json:"title"`
	Finished                    bool   `json:"finished"`
	Description                 string `json:"description"`
	Logo                        string `json:"logo"`
	CreatedBy                   string `json:"created_by"`
	UpdatedBy                   string `json:"updated_by"`
	CreatedAt                   int64  `json:"created_at"`
	UpdatedAt                   int64  `json:"updated_at"`
	PrimaryResearchArea         int64  `json:"primary_research_area"`
	RelResearchAreaCreatedBy    string `json:"research_area_created_by,omitempty"`
	RelResearchAreaCreatedAt    int64  `json:"research_area_created_at,omitempty"`
	RelFinancedProjectCreatedBy string `json:"financed_project_created_by,omitempty"`
	RelFinancedProjectCreatedAt string `json:"financed_project_created_at,omitempty"`
	RelPublicationCreatedBy     string `json:"publication_created_by,omitempty"`
	RelPublicationCreatedAt     string `json:"publication_created_at,omitempty"`
	RelStudentWorkCreatedBy     string `json:"student_work_created_by,omitempty"`
	RelStudentWorkCreatedAt     int64  `json:"student_work_created_at,omitempty"`
	RelPartnerCreatedBy         string `json:"partner_created_by,omitempty"`
	RelPartnerCreatedAt         int64  `json:"partner_created_at,omitempty"`
	RelMemberCreatedBy          string `json:"member_created_by,omitempty"`
	RelMemberCreatedAt          int64  `json:"member_created_at,omitempty"`
	RelArticleCreatedBy         string `json:"article_created_by,omitempty"`
	RelArticleCreatedAt         int64  `json:"article_created_at,omitempty"`
}

func ResearchLineCreate(title string, finished bool, description, createdBy string, primaryResearchArea int64) (id int64, verr *ValidationError, err error) {
	verr = ResearchLineValidate(title, description)
	if verr != nil {
		return
	}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO research_line(title,finished,description,created_by,updated_by,created_at,updated_at,primary_research_area) VALUES(?,?,?,?,?,?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	result, err := stmt.Exec(title, finished, description, createdBy, createdBy, ts, ts, primaryResearchArea)
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
func ResearchLineUpdate(id int64, title string, finished bool, description string, updatedBy string, primaryResearchArea int64) (numRows int64, verr *ValidationError, err error) {
	verr = ResearchLineValidate(title, description)
	if verr != nil {
		return
	}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "UPDATE research_line SET title=?,finished=?,description=?,updated_by=?,updated_at=?,primary_research_area=? WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	result, err := stmt.Exec(title, finished, description, updatedBy, ts, primaryResearchArea, id)
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
func ResearchLineUpdateLogo(id int64, logo string, updatedBy string) (numRows int64, verr *ValidationError, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "UPDATE research_line SET logo=?,updated_by=?,updated_at=? WHERE id=?"
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
func ResearchLineDelete(id int64) (numRows int64, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM research_line WHERE id=?"
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
func ResearchLineGetAll() (researchLines []*ResearchLine, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM research_line"
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
		p := ResearchLine{}
		err = rows.Scan(&p.Id, &p.Title, &p.Finished, &p.Description, &p.Logo, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.PrimaryResearchArea)
		if err != nil {
			return
		}
		researchLines = append(researchLines, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func ResearchLineGetById(id int64) (researchLine *ResearchLine, err error) {
	researchLine = &ResearchLine{}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM research_line WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&researchLine.Id, &researchLine.Title, &researchLine.Finished, &researchLine.Description, &researchLine.Logo, &researchLine.CreatedBy, &researchLine.UpdatedBy, &researchLine.CreatedAt, &researchLine.UpdatedAt, &researchLine.PrimaryResearchArea)
	if err != nil {
		return
	}
	return
}
func ResearchLineGetByPrimaryResearchArea(researchAreaId int64) (researchLines []*ResearchLine, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM research_line WHERE primary_research_area=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(researchAreaId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := ResearchLine{}
		err = rows.Scan(&p.Id, &p.Title, &p.Finished, &p.Description, &p.Logo, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.PrimaryResearchArea)
		if err != nil {
			return
		}
		researchLines = append(researchLines, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}

func ResearchLineGetByResearchArea(researchAreaId int64) (researchLines []*ResearchLine, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT research_line.*,research_area_research_line.created_by,research_area_research_line.created_at FROM research_area_research_line INNER JOIN research_line ON research_area_research_line.research_line=research_line.id  WHERE research_area=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(researchAreaId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := ResearchLine{}
		err = rows.Scan(&p.Id, &p.Title, &p.Finished, &p.Description, &p.Logo, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.PrimaryResearchArea, &p.RelResearchAreaCreatedBy, &p.RelResearchAreaCreatedAt)
		if err != nil {
			return
		}
		researchLines = append(researchLines, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func ResearchLineGetByFinancedProject(financedProjectId int64) (researchLines []*ResearchLine, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT research_line.*,research_line_financed_project.created_by,research_line_financed_project.created_at FROM research_line_financed_project INNER JOIN research_line ON research_line_financed_project.research_line=research_line.id  WHERE financed_project=?"
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
		p := ResearchLine{}
		err = rows.Scan(&p.Id, &p.Title, &p.Finished, &p.Description, &p.Logo, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.PrimaryResearchArea, &p.RelFinancedProjectCreatedBy, &p.RelFinancedProjectCreatedAt)
		if err != nil {
			return
		}
		researchLines = append(researchLines, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}

func ResearchLineGetByPublication(publicationId int64) (researchLines []*ResearchLine, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT research_line.*,research_line_publication.created_by,research_line_publication.created_at FROM research_line_publication INNER JOIN research_line ON research_line_publication.research_line=research_line.id  WHERE publication=?"
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
		p := ResearchLine{}
		err = rows.Scan(&p.Id, &p.Title, &p.Finished, &p.Description, &p.Logo, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.PrimaryResearchArea, &p.RelPublicationCreatedBy, &p.RelPublicationCreatedAt)
		if err != nil {
			return
		}
		researchLines = append(researchLines, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func ResearchLineGetByStudentWork(studentWorkId int64) (researchLines []*ResearchLine, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT research_line.*,research_line_student_work.created_by,research_line_student_work.created_at FROM research_line_student_work INNER JOIN research_line ON research_line_student_work.research_line=research_line.id  WHERE student_work=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(studentWorkId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := ResearchLine{}
		err = rows.Scan(&p.Id, &p.Title, &p.Finished, &p.Description, &p.Logo, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.PrimaryResearchArea, &p.RelStudentWorkCreatedBy, &p.RelStudentWorkCreatedAt)
		if err != nil {
			return
		}
		researchLines = append(researchLines, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func ResearchLineGetByPartner(partnerId int64) (researchLines []*ResearchLine, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT research_line.*,research_line_partner.created_by,research_line_partner.created_at FROM research_line_partner INNER JOIN research_line ON research_line_partner.research_line=research_line.id  WHERE partner=?"
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
		p := ResearchLine{}
		err = rows.Scan(&p.Id, &p.Title, &p.Finished, &p.Description, &p.Logo, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.PrimaryResearchArea, &p.RelPartnerCreatedBy, &p.RelPartnerCreatedAt)
		if err != nil {
			return
		}
		researchLines = append(researchLines, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func ResearchLineGetByMember(memberId int64) (researchLines []*ResearchLine, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT research_line.*,research_line_member.created_by,research_line_member.created_at FROM research_line_member INNER JOIN research_line ON research_line_member.research_line=research_line.id  WHERE member=?"
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
		p := ResearchLine{}
		err = rows.Scan(&p.Id, &p.Title, &p.Finished, &p.Description, &p.Logo, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.PrimaryResearchArea, &p.RelMemberCreatedBy, &p.RelMemberCreatedAt)
		if err != nil {
			return
		}
		researchLines = append(researchLines, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func ResearchLineGetByArticle(articleId int64) (researchLines []*ResearchLine, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT research_line.*,research_line_article.created_by,research_line_article.created_at FROM research_line_article INNER JOIN research_line ON research_line_article.research_line=research_line.id  WHERE article=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(articleId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := ResearchLine{}
		err = rows.Scan(&p.Id, &p.Title, &p.Finished, &p.Description, &p.Logo, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.PrimaryResearchArea, &p.RelArticleCreatedBy, &p.RelArticleCreatedAt)
		if err != nil {
			return
		}
		researchLines = append(researchLines, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}

func ResearchLineCount() (count int64, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM research_line"
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
func ResearchLineExists(id int64) (exists bool, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM research_line WHERE id=?"
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

func ResearchLineAddResearchArea(id, researchAreaId int64, createdBy string) (verr *ValidationError, err error) {
	researchLine, err := ResearchLineGetById(id)
	if err != nil {
		return
	}
	if researchLine.PrimaryResearchArea == researchAreaId {
		verr = &ValidationError{"research_area", "this research area is already the primary"}
		return
	}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO research_area_research_line(research_area,research_line,created_by,created_at) VALUES(?,?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	_, err = stmt.Exec(researchAreaId, id, createdBy, ts)
	if err != nil {
		if IsDbError1062(err) {
			verr = &ValidationError{"research_area", "this research area has already been added"}
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
func ResearchLineRemoveResearchArea(id, researchAreaId int64) (removed bool, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM research_area_research_line WHERE research_area=? AND research_line=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(researchAreaId, id)
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

func ResearchLineGetResearchAreas(id int64) (researchAreas []*ResearchArea, err error) {
	researchAreas, err = ResearchAreaGetByResearchLine(id)
	return
}

func ResearchLineAddFinancedProject(id, financedProjectId int64, createdBy string) (verr *ValidationError, err error) {
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
	_, err = stmt.Exec(id, financedProjectId, createdBy, ts)
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
func ResearchLineRemoveFinancedProject(id, financedProjectId int64) (removed bool, err error) {
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

func ResearchLineGetFinancedProjects(id int64) (financedProjects []*FinancedProject, err error) {
	financedProjects, err = FinancedProjectGetByResearchLine(id)
	return
}
func ResearchLineAddArticle(id, articleId int64, createdBy string) (verr *ValidationError, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO research_line_article(research_line,article,created_by,created_at) VALUES(?,?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	_, err = stmt.Exec(id, articleId, createdBy, ts)
	if err != nil {
		if IsDbError1062(err) {
			verr = &ValidationError{"article", "this article has already been added"}
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
func ResearchLineRemoveArticle(id, articleId int64) (removed bool, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM research_line_article WHERE research_line=? AND article=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(id, articleId)
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

func ResearchLineGetArticles(id int64) (articles []*Article, err error) {
	articles, err = ArticleGetByResearchLine(id)
	return
}
func ResearchLineAddPartner(id, partnerId int64, createdBy string) (verr *ValidationError, err error) {
	db, err := DBGet()
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
	_, err = stmt.Exec(id, partnerId, createdBy, ts)
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
func ResearchLineRemovePartner(id, partnerId int64) (removed bool, err error) {
	db, err := DBGet()
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
	result, err := stmt.Exec(id, partnerId)
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

func ResearchLineGetPartners(id int64) (partners []*Partner, err error) {
	partners, err = PartnerGetByResearchLine(id)
	return
}

func ResearchLineAddMember(id, memberId int64, createdBy string) (verr *ValidationError, err error) {
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
func ResearchLineRemoveMember(id, memberId int64) (removed bool, err error) {
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

func ResearchLineGetMembers(id int64) (members []*Member, err error) {
	members, err = MemberGetByResearchLine(id)
	return
}

func ResearchLineAddPublication(id, publicationId int64, createdBy string) (verr *ValidationError, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO research_line_publication(research_line,publication,created_by,created_at) VALUES(?,?,?,?)"
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
func ResearchLineRemovePublication(id, publicationId int64) (removed bool, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM research_line_publication WHERE research_line=? AND publication=?"
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

func ResearchLineGetPublications(id int64) (publications []*Publication, err error) {
	publications, err = PublicationGetByResearchLine(id)
	return
}
func ResearchLineAddStudentWork(id, studentWorkId int64, createdBy string) (verr *ValidationError, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO research_line_student_work(research_line,student_work,created_by,created_at) VALUES(?,?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	_, err = stmt.Exec(id, studentWorkId, createdBy, ts)
	if err != nil {
		if IsDbError1062(err) {
			verr = &ValidationError{"student_work", "this student work has already been added"}
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
func ResearchLineRemoveStudentWork(id, studentWorkId int64) (removed bool, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM research_line_student_work WHERE research_line=? AND student_work=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(id, studentWorkId)
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

func ResearchLineGetStudentWorks(id int64) (studentWorks []*StudentWork, err error) {
	studentWorks, err = StudentWorkGetByResearchLine(id)
	return
}
func ResearchLineGetColumns() []string {
	columns := []string{
		"id",
		"title",
		"finished",
		"description",
		"logo",
		"created_by",
		"updated_by",
		"created_at",
		"updated_at",
		"primary_research_area",
	}
	return columns
}
func ResearchLineValidateTitle(title string) (verr *ValidationError) {
	if verr = ValidateNotEmpty("title", title); verr != nil {
		return verr
	}
	return ValidateLength("title", title, 200)
}
func ResearchLineValidateDescription(description string) (verr *ValidationError) {
	return ValidateLength("description", description, 200)
}
func ResearchLineValidate(title, description string) (verr *ValidationError) {
	verr = ResearchLineValidateTitle(title)
	if verr != nil {
		return
	}
	verr = ResearchLineValidateDescription(description)
	if verr != nil {
		return
	}
	return
}
