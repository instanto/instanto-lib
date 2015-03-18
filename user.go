package instanto_lib_db

import (
	"database/sql"

	"code.google.com/p/go.crypto/bcrypt"
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

func UserCreate(username, email, password string, enabled bool, displayName, ugroup string) (ok bool, verr *ValidationError, err error) {
	verr = UserValidate(username, email, password, displayName)
	if verr != nil {
		return
	}
	_, err = bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return
	}
	db, err := DBGet()
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

func UserGetByUsername(username string) (user *User, err error) {
	user = &User{}
	db, err := DBGet()
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
func UserCheckLogin(username, password string) (verr *ValidationError, err error) {
	user, err := UserGetByUsername(username)
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

func UserValidateUsername(username string) (verr *ValidationError) {
	// check is unique
	return
}
func UserValidateEmail(email string) (verr *ValidationError) {
	// check is unique
	return
}
func UserValidatePassword(password string) (verr *ValidationError) {
	// check string lenght limit
	return
}
func UserValidateDisplayName(dislpayName string) (verr *ValidationError) {
	// check string lenght limit
	return
}
func UserValidate(username, email, password, displayName string) (verr *ValidationError) {
	verr = UserValidateUsername(username)
	if verr != nil {
		return
	}
	verr = UserValidateEmail(email)
	if verr != nil {
		return
	}
	verr = UserValidatePassword(password)
	if verr != nil {
		return
	}
	verr = UserValidateDisplayName(displayName)
	if verr != nil {
		return
	}
	return
}
