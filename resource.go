package instantolib

import (
	"time"
)

type Resource struct {
	Id                       int64  `json:"id"`
	Filename                 string `json:"filename"`
	MimeType                 int64  `json:"mime_type"`
	Size                     int64  `json:"size"`
	Private                  bool   `json:"private"`
	ResourceType             int64  `json:"resource_type"`
	CreatedBy                string `json:"created_by"`
	UpdatedBy                string `json:"updated_by"`
	CreatedAt                int64  `json:"created_at"`
	UpdatedAt                int64  `json:"updated_at"`
	RelResearchLineCreatedBy string `json:"research_line_created_by,omitempty"`
	RelResearchLineCreatedAt int64  `json:"research_line_created_at,omitempty"`
}

func (dbp *DBProvider) ResourceCreate(filename, mimeType string, size int64, private bool, createdBy string, resourceType int64) (id int64, verr *ValidationError, err error) {
	verr = resourceValidate(filename, mimeType, size)
	if verr != nil {
		return
	}
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO resource(filename,mime_type,size,private,created_by,updated_by,created_at,updated_at,resource_type) VALUES(?,?,?,?,?,?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	result, err := stmt.Exec(filename, mimeType, size, private, createdBy, createdBy, ts, ts, resourceType)
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
func (dbp *DBProvider) ResourceUpdate(id int64, filename, mimeType string, size int64, private bool, updatedBy string, resourceType int64) (numRows int64, verr *ValidationError, err error) {
	verr = resourceValidate(filename, mimeType, size)
	if verr != nil {
		return
	}
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "UPDATE resource SET filename=?,mime_type=?,size=?,private=?,updated_by=?,updated_at=?,resource_type=? WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	ts := time.Now().Unix()
	result, err := stmt.Exec(filename, mimeType, size, private, updatedBy, ts, resourceType, id)
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
func (dbp *DBProvider) ResourceDelete(id int64) (numRows int64, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM resource WHERE id=?"
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
func (dbp *DBProvider) ResourceGetAll() (resources []*Resource, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM resource ORDER BY filename ASC"
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
		p := Resource{}
		err = rows.Scan(&p.Id, &p.Filename, &p.MimeType, &p.Size, &p.Private, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.ResourceType)
		if err != nil {
			return
		}
		resources = append(resources, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func (dbp *DBProvider) ResourceGetById(id int64) (resource *Resource, err error) {
	resource = &Resource{}
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM resource WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&resource.Id, &resource.Filename, &resource.MimeType, &resource.Size, &resource.Private, &resource.CreatedBy, &resource.UpdatedBy, &resource.CreatedAt, &resource.UpdatedAt, &resource.ResourceType)
	if err != nil {
		return
	}
	return
}
func (dbp *DBProvider) ResourceGetByResourceType(resourceTypeId int64) (resources []*Resource, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM resource WHERE resource_type=? ORDER BY filename ASC"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(resourceTypeId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := Resource{}
		err = rows.Scan(&p.Id, &p.Filename, &p.MimeType, &p.Size, &p.Private, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.ResourceType)
		if err != nil {
			return
		}
		resources = append(resources, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func (dbp *DBProvider) ResourceGetByResearchLine(researchLineId int64) (resources []*Resource, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT resource.*,research_line_resource.created_by,research_line_resource.created_at FROM research_line_resource INNER JOIN resource ON research_line_resource.resource=resource.id  WHERE research_line=? ORDER BY filename ASC"
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
		p := Resource{}
		err = rows.Scan(&p.Id, &p.Filename, &p.MimeType, &p.Size, &p.Private, &p.CreatedBy, &p.UpdatedBy, &p.CreatedAt, &p.UpdatedAt, &p.ResourceType, &p.RelResearchLineCreatedBy, &p.RelResearchLineCreatedAt)
		if err != nil {
			return
		}
		resources = append(resources, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func (dbp *DBProvider) ResourceCount() (count int64, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM resource"
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
func (dbp *DBProvider) ResourceExists(id int64) (exists bool, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM resource WHERE id=?"
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
func (dbp *DBProvider) ResourceAddResearchLine(id, researchLineId int64, createdBy string) (verr *ValidationError, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO research_line_resource(research_line,resource,created_by,created_at) VALUES(?,?,?,?)"
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
func (dbp *DBProvider) ResourceRemoveResearchLine(id, researchLineId int64) (removed bool, err error) {
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "DELETE FROM research_line_resource WHERE research_line=? AND resource=?"
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

func (dbp *DBProvider) ResourceGetResearchLines(id int64) (researchLines []*ResearchLine, err error) {
	researchLines, err = dbp.ResearchLineGetByResource(id)
	return
}
func (dbp *DBProvider) ResourceGetColumns() []string {
	columns := []string{
		"id",
		"filename",
		"mimeType",
		"size",
		"private",
		"resource_type",
		"created_by",
		"updated_by",
		"created_at",
		"updated_at",
		"resourceType",
	}
	return columns
}
func resourceValidateFilename(filename string) (verr *ValidationError) {
	if verr = validateNotEmpty("filename", filename); verr != nil {
		return verr
	}
	return validateLength("filename", filename, 200)
}
func resourceValidateMimeType(mimeType string) (verr *ValidationError) {
	return validateLength("mimeType", mimeType, 200)
}
func resourceValidateDate(size int64) (verr *ValidationError) {
	return validateIsNumber("size", size)
}
func resourceValidate(filename, mimeType string, size int64) (verr *ValidationError) {
	verr = resourceValidateFilename(filename)
	if verr != nil {
		return
	}
	verr = resourceValidateMimeType(mimeType)
	if verr != nil {
		return
	}
	verr = resourceValidateDate(size)
	if verr != nil {
		return
	}
	return
}
