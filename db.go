package instantolib

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

func NewDBProvider(dsn string) (*DBProvider, error) {
	return &DBProvider{dsn}, nil
}

type DBProvider struct {
	dsn string
}

func (dbp *DBProvider) getDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", dbp.dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// IsDbError1062 checks if the error is a Error 1062: Duplicate entry
func IsDbError1062(err error) (is bool) {
	errString := err.Error()
	is = strings.Contains(errString, "Error 1062")
	return
}

// IsDbError1452 checks if the error is a Error 1452: Cannot add or update a child row
// We parse the error to return the wrong field
// Error 1452: Cannot add or update a child row: a foreign key constraint fails (`instanto`.`article`, CONSTRAINT `fk_article_3` FOREIGN KEY (`newspaper`) REFERENCES `newspaper` (`id`) ON DELETE CASCADE ON UPDATE CASCADE)
func IsDbError1452(err error) (is bool, field string) {
	errString := err.Error()
	is = strings.Contains(errString, "Error 1452")
	if !is {
		return
	}
	indexFK := strings.Index(errString, "FOREIGN KEY")
	indexRef := strings.Index(errString, "REFERENCES")
	indexFK += 14
	indexRef -= 3
	field = errString[indexFK:indexRef]
	return
}
