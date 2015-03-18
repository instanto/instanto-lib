package instanto_lib_db

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Newspaper struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Web       string `json:"web"`
	Logo      string `json:"logo"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func NewspaperCreate(name, web, createdBy string) (id int64, verr *ValidationError, err error) {
	verr = NewspaperValidate(name, web)
	if verr != nil {
		return
	}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO newspaper(name,web,created_by,updated_by,created_at,updated_at) VALUES(?,?,?,?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	result, err := stmt.Exec(name, web, createdBy, createdBy, ts, ts)
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
func NewspaperUpdate(id int64, name, web, updatedBy string) (numRows int64, verr *ValidationError, err error) {
	verr = NewspaperValidate(name, web)
	if verr != nil {
		return
	}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "UPDATE newspaper SET name=?,web=?,updated_by=?,updated_at=? WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	result, err := stmt.Exec(name, web, updatedBy, ts, id)
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

func NewspaperUpdateLogo(id int64, logo string, updatedBy string) (numRows int64, verr *ValidationError, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "UPDATE newspaper SET logo=?,updated_by=?,updated_at=? WHERE id=?"
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

func NewspaperDelete(id int64) (numRows int64, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM newspaper WHERE id=?"
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
func NewspaperGetAll() (newspapers []*Newspaper, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM newspaper"
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
		p := Newspaper{}
		err = rows.Scan(&p.Id, &p.Name, &p.Web, &p.Logo, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return
		}
		newspapers = append(newspapers, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func NewspaperGetById(id int64) (newspaper *Newspaper, err error) {
	newspaper = &Newspaper{}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM newspaper WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&newspaper.Id, &newspaper.Name, &newspaper.Web, &newspaper.Logo, &newspaper.CreatedBy, &newspaper.UpdatedBy, &newspaper.CreatedAt, &newspaper.UpdatedAt)
	if err != nil {
		return
	}
	return
}
func NewspaperCount() (count int64, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM newspaper"
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
func NewspaperExists(id int64) (exists bool, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM newspaper WHERE id=?"
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
func NewspaperGetArticles(id int64) (articles []*Article, err error) {
	articles, err = ArticleGetByNewspaper(id)
	return
}
func NewspaperGetColumns() []string {
	columns := []string{
		"id",
		"name",
		"web",
		"logo",
		"created_by",
		"updated_by",
		"created_at",
		"updated_at",
	}
	return columns
}
func NewspaperValidateName(name string) (verr *ValidationError) {
	if verr = ValidateNotEmpty("name", name); verr != nil {
		return verr
	}
	return ValidateLength("name", name, 200)
}
func NewspaperValidateWeb(web string) (verr *ValidationError) {
	return ValidateLength("web", web, 200)
}
func NewspaperValidateLogo(logo string) (err *ValidationError) {
	return ValidateLength("logo", logo, 200)
}
func NewspaperValidate(name, web string) (verr *ValidationError) {
	verr = NewspaperValidateName(name)
	if verr != nil {
		return
	}
	verr = NewspaperValidateWeb(web)
	if verr != nil {
		return
	}
	return
}
