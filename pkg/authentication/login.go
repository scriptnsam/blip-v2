package authentication

import (
	"database/sql"

	"github.com/scriptnsam/blip-v2/pkg/database"
	"github.com/scriptnsam/blip-v2/pkg/security"
)

var (
	loggedIn bool = false
	userId   int
)

func Login(username string, password string) (string, error) {
	db, err := database.General()
	if err != nil {
		return "", err
	}
	defer db.Close()
	
	var (
		id int
		hashedPassword string
	)

	err = db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username).Scan(&id, &username, &hashedPassword)
	if err != nil {
		if err== sql.ErrNoRows {
			return "Username does not exist", nil
		}
		return "", err
	}



	r:=security.VerifyPassword(password, hashedPassword)
	if !r {
		return "Incorrect password", nil
	}

	userId = id
	loggedIn = true
	return "Logged in", nil
}

func Logout() {
	loggedIn = false
}

func IsLoggedIn() (bool, int) {
	if !loggedIn {
		return loggedIn, 0
	}
	loggedIn = true
	return loggedIn, userId
}