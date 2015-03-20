package instantolib

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	Id                       int64  `json:"id"`
	Title                    string `json:"title"`
	Web                      string `json:"web"`
	Date                     int64  `json:"date"`
	CreatedBy                string `json:"created_by"`
	UpdatedBy                string `json:"updated_by"`
	CreatedAt                int64  `json:"created_at"`
	UpdatedAt                int64  `json:"updated_at"`
	Newspaper                int64  `json:"newspaper"`
	RelResearchLineCreatedBy string `json:"research_line_created_by,omitempty"`
	RelResearchLineCreatedAt int64  `json:"research_line_created_at,omitempty"`
}

func (dbp *DBProvider) ArticleCreate(title, web string, date int64, createdBy string, newspaper int64) (id int64, verr *ValidationError, err error) {
	verr = articleValidate(title, web, date)
	if verr != nil {
		return
	}
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO article(title,web,date,created_by,updated_by,created_at,updated_at,newspaper) VALUES(?,?,?,?,?,?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	result, err := stmt.Exec(title, web, date, createdBy, createdBy, ts, ts, newspaper)
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
func (dbp *DBProvider) ArticleUpdate(id int64, title, web string, date int64, updatedBy string, newspaper int64) (numRows int64, verr *ValidationError, err error) {
	verr = articleValidate(title, web, date)
	if verr != nil {
		return
	}
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "UPDATE article SET title=?,web=?,date=?,updated_by=?,updated_at=?,newspaper=? WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	result, err := stmt.Exec(title, web, date, updatedBy, ts, newspaper, id)
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
func (dbp *DBProvider) ArticleDelete(id int64) (numRows int64, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM article WHERE id=?"
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
func (dbp *DBProvider) ArticleGetAll() (articles []*Article, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM article"
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
		p := Article{}
		err = rows.Scan(&p.Id, &p.Title, &p.Web, &p.Date, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.Newspaper)
		if err != nil {
			return
		}
		articles = append(articles, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func (dbp *DBProvider) ArticleGetById(id int64) (article *Article, err error) {
	article = &Article{}
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM article WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&article.Id, &article.Title, &article.Web, &article.Date, &article.CreatedBy, &article.UpdatedBy, &article.CreatedAt, &article.UpdatedAt, &article.Newspaper)
	if err != nil {
		return
	}
	return
}
func (dbp *DBProvider) ArticleGetByNewspaper(newspaperId int64) (articles []*Article, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM article WHERE newspaper=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(newspaperId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := Article{}
		err = rows.Scan(&p.Id, &p.Title, &p.Web, &p.Date, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.Newspaper)
		if err != nil {
			return
		}
		articles = append(articles, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func (dbp *DBProvider) ArticleGetByResearchLine(researchLineId int64) (articles []*Article, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT article.*,research_line_article.created_by,research_line_article.created_at FROM research_line_article INNER JOIN article ON research_line_article.article=article.id  WHERE research_line=?"
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
		p := Article{}
		err = rows.Scan(&p.Id, &p.Title, &p.Web, &p.Date, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.Newspaper, &p.RelResearchLineCreatedBy, &p.RelResearchLineCreatedAt)
		if err != nil {
			return
		}
		articles = append(articles, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func (dbp *DBProvider) ArticleCount() (count int64, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM article"
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
func (dbp *DBProvider) ArticleExists(id int64) (exists bool, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM article WHERE id=?"
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
func (dbp *DBProvider) ArticleAddResearchLine(id, researchLineId int64, createdBy string) (verr *ValidationError, err error) {
	db, err := dbp.getDB()
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
	_, err = stmt.Exec(researchLineId, id, createdBy, ts)
	if err != nil {
		if IsDbError1062(err) {
			verr = &ValidationError{"research_line", "this research_line has already been added"}
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
func (dbp *DBProvider) ArticleRemoveResearchLine(id, researchLineId int64) (removed bool, err error) {
	db, err := dbp.getDB()
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
		return
	}
	return
}

func (dbp *DBProvider) ArticleGetResearchLines(id int64) (researchLines []*ResearchLine, err error) {
	researchLines, err = dbp.ResearchLineGetByArticle(id)
	return
}
func (dbp *DBProvider) ArticleGetColumns() []string {
	columns := []string{
		"id",
		"title",
		"web",
		"date",
		"created_by",
		"updated_by",
		"created_at",
		"updated_at",
		"newspaper",
	}
	return columns
}
func articleValidateTitle(title string) (verr *ValidationError) {
	if verr = validateNotEmpty("title", title); verr != nil {
		return verr
	}
	return validateLength("title", title, 200)
}
func articleValidateWeb(web string) (verr *ValidationError) {
	return validateLength("web", web, 200)
}
func articleValidateDate(date int64) (verr *ValidationError) {
	return validateIsNumber("date", date)
}
func articleValidate(title, web string, date int64) (verr *ValidationError) {
	verr = articleValidateTitle(title)
	if verr != nil {
		return
	}
	verr = articleValidateWeb(web)
	if verr != nil {
		return
	}
	verr = articleValidateDate(date)
	if verr != nil {
		return
	}
	return
}
