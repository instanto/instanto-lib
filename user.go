package instantolib

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"-"`
	Enabled     bool   `json:"enabled"`
	DisplayName string `json:"display_name"`
	UGroup      string `json:"ugroup"`
}

func (dbp *DBProvider) UserCreate(username, email, password string, enabled bool, displayName, ugroup string) (ok bool, verr *ValidationError, err error) {
	verr = userValidate(username, email, password, displayName)
	if verr != nil {
		return
	}
	_, err = bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return
	}
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "INSERT INTO user(username,email,password,enabled,display_name,ugroup) VALUES(?,?,?,?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		if IsDbError1062(err) {
			verr = &ValidationError{"username", "this username is taken, use another"}
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
	defer stmt.Close()
	_, err = stmt.Exec(username, email /*string(hashedPassword)*/, password, enabled, displayName, ugroup)
	if err != nil {
		return
	}
	ok = true
	return
}

func (dbp *DBProvider) UserGetByUsername(username string) (user *User, err error) {
	user = &User{}
	db, err := dbp.getDB()
	if err != nil {
		return
	}
	defer db.Close()
	query := "SELECT * FROM user WHERE username=?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(username).Scan(&user.Username, &user.Email, &user.Password, &user.Enabled, &user.DisplayName, &user.UGroup)
	if err != nil {
		return
	}
	return
}
func (dbp *DBProvider) UserCheckLogin(username, password string) (user *User, verr *ValidationError, err error) {
	user, err = dbp.UserGetByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			verr = &ValidationError{"username/passsword", "not match"}
			return
		}
		return
	}
	//err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if password != user.Password {
		verr = &ValidationError{"username/passsword", "not match"}
		return
	}
	return
}

func userValidateUsername(username string) (verr *ValidationError) {
	// check is unique
	return
}
func userValidateEmail(email string) (verr *ValidationError) {
	// check is unique
	return
}
func userValidatePassword(password string) (verr *ValidationError) {
	// check string lenght limit
	return
}
func userValidateDisplayName(dislpayName string) (verr *ValidationError) {
	// check string lenght limit
	return
}
func userValidate(username, email, password, displayName string) (verr *ValidationError) {
	verr = userValidateUsername(username)
	if verr != nil {
		return
	}
	verr = userValidateEmail(email)
	if verr != nil {
		return
	}
	verr = userValidatePassword(password)
	if verr != nil {
		return
	}
	verr = userValidateDisplayName(displayName)
	if verr != nil {
		return
	}
	return
}
