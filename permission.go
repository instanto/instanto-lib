package instanto_lib_db

import (
	_ "github.com/go-sql-driver/mysql"
)

type Permission struct {
	Id          string `json:"id"`
	DisplayName string `json:"display_name"`
}

func PermissionGetAll() (permissions []*Permission, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM permission"
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
		p := Permission{}
		err = rows.Scan(&p.Id, &p.DisplayName)
		if err != nil {
			return
		}
		permissions = append(permissions, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}
func PermissionGetById(id string) (permission *Permission, err error) {
	permission = &Permission{}
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM permission WHERE id=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&permission.Id, &permission.DisplayName)
	if err != nil {
		return
	}
	return
}

func PermissionGetByRol(rolId string) (permissions []*Permission, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT permission.*, FROM rol_permission INNER JOIN permission ON rol_permission.permission=permission.id  WHERE rol=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(rolId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := Permission{}
		err = rows.Scan(&p.Id, &p.DisplayName)
		if err != nil {
			return
		}
		permissions = append(permissions, &p)
	}
	err = rows.Err()
	if err != nil {
		return
	}
	return
}

/*

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
*/
func PermissionCount() (count int64, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM permission"
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
func PermissionExists(id string) (exists bool, err error) {
	db, err := DBGet()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT COUNT(id) as count FROM permission WHERE id=?"
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
func PermissionGetColumns() []string {
	columns := []string{
		"id",
		"display_name",
	}
	return columns
}
