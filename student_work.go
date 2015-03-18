package instanto_lib_db

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type StudentWork struct {
	Id                       int64  `json:"id"`
	Title                    string `json:"title"`
	Year                     int64  `json:"year"`
	School                   string `json:"school"`
	Volume                   string `json:"volume"`
	CreatedBy                string `json:"created_by"`
	UpdatedBy                string `json:"updated_by"`
	CreatedAt                int64  `json:"created_at"`
	UpdatedAt                int64  `json:"updated_at"`
	StudentWorkType          int64  `json:"student_work_type"`
	Author                   int64  `json:"author"`
	RelResearchLineCreatedBy string `json:"research_line_created_by,omitempty"`
	RelResearchLineCreatedAt int64  `json:"research_line_created_at,omitempty"`
}

func StudentWorkCreate(title string, year int64, school, volume, createdBy string, studentWorkType, author int64) (id int64, verr *ValidationError, err error) {
	verr = StudentWorkValidate(title, year, school, volume)
	if verr != nil {
		return
	}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO student_work(title,year,school,volume,created_by,updated_by,created_at,updated_at,student_work_type, author) VALUES(?,?,?,?,?,?,?,?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	result, err := stmt.Exec(title, year, school, volume, createdBy, createdBy, ts, ts, studentWorkType, author)
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
func StudentWorkUpdate(id int64, title string, year int64, school, volume, updatedBy string, studentWorkType, author int64) (numRows int64, verr *ValidationError, err error) {
	verr = StudentWorkValidate(title, year, school, volume)
	if verr != nil {
		return
	}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "UPDATE student_work SET title=?,year=?,school=?,volume=?,updated_by=?,updated_at=?,student_work_type=?,author=? WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	result, err := stmt.Exec(title, year, school, volume, updatedBy, ts, studentWorkType, author, id)
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
func StudentWorkDelete(id int64) (numRows int64, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM student_work WHERE id=?"
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
func StudentWorkGetAll() (studentWorks []*StudentWork, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM student_work"
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
		p := StudentWork{}
		err = rows.Scan(&p.Id, &p.Title, &p.Year, &p.School, &p.Volume, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.StudentWorkType, &p.Author)
		if err != nil {
			return
		}
		studentWorks = append(studentWorks, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func StudentWorkGetById(id int64) (studentWork *StudentWork, err error) {
	studentWork = &StudentWork{}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM student_work WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&studentWork.Id, &studentWork.Title, &studentWork.Year, &studentWork.School, &studentWork.Volume, &studentWork.CreatedBy, &studentWork.UpdatedBy, &studentWork.CreatedAt, &studentWork.UpdatedAt, &studentWork.StudentWorkType, &studentWork.Author)
	if err != nil {
		return
	}
	return
}
func StudentWorkGetByStudentWorkType(studentWorkTypeId int64) (studentWorks []*StudentWork, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM student_work WHERE student_work_type=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(studentWorkTypeId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := StudentWork{}
		err = rows.Scan(&p.Id, &p.Title, &p.Year, &p.School, &p.Volume, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.StudentWorkType, &p.Author)
		if err != nil {
			return
		}
		studentWorks = append(studentWorks, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func StudentWorkGetByAuthor(authorId int64) (studentWorks []*StudentWork, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM student_work WHERE author=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(authorId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := StudentWork{}
		err = rows.Scan(&p.Id, &p.Title, &p.Year, &p.School, &p.Volume, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.StudentWorkType, &p.Author)
		if err != nil {
			return
		}
		studentWorks = append(studentWorks, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func StudentWorkGetByResearchLine(researchLineId int64) (studentWorks []*StudentWork, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT student_work.*,research_line_student_work.created_by,research_line_student_work.created_at FROM research_line_student_work INNER JOIN student_work ON research_line_student_work.student_work=student_work.id  WHERE research_line=?"
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
		p := StudentWork{}
		err = rows.Scan(&p.Id, &p.Title, &p.Year, &p.School, &p.Volume, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.StudentWorkType, &p.Author, &p.RelResearchLineCreatedBy, &p.RelResearchLineCreatedAt)
		if err != nil {
			return
		}
		studentWorks = append(studentWorks, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func StudentWorkCount() (count int64, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM student_work"
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
func StudentWorkExists(id int64) (exists bool, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM student_work WHERE id=?"
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
func StudentWorkAddResearchLine(id, researchLineId int64, createdBy string) (verr *ValidationError, err error) {
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
func StudentWorkRemoveResearchLine(id, researchLineId int64) (removed bool, err error) {
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

func StudentWorkGetResearchLines(id int64) (researchLines []*ResearchLine, err error) {
	researchLines, err = ResearchLineGetByStudentWork(id)
	return
}
func StudentWorkGetColumns() []string {
	columns := []string{
		"id",
		"title",
		"year",
		"school",
		"volume",
		"created_by",
		"updated_by",
		"created_at",
		"updated_at",
		"student_work_type",
		"author",
	}
	return columns
}
func StudentWorkValidateTitle(title string) (verr *ValidationError) {
	if verr = ValidateNotEmpty("title", title); verr != nil {
		return verr
	}
	return ValidateLength("title", title, 200)
}
func StudentWorkValidateYear(year int64) (verr *ValidationError) {
	return ValidateIsNumber("year", year)
}
func StudentWorkValidateSchool(school string) (verr *ValidationError) {
	return ValidateLength("school", school, 200)
}
func StudentWorkValidateVolume(volume string) (verr *ValidationError) {
	return ValidateLength("volume", volume, 200)
}
func StudentWorkValidate(title string, year int64, school, volume string) (verr *ValidationError) {
	verr = StudentWorkValidateTitle(title)
	if verr != nil {
		return
	}
	verr = StudentWorkValidateYear(year)
	if verr != nil {
		return
	}
	verr = StudentWorkValidateSchool(school)
	if verr != nil {
		return
	}
	verr = StudentWorkValidateVolume(volume)
	if verr != nil {
		return
	}
	return
}
