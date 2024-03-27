package authentication

import (
	"database/sql"

	"github.com/scriptnsam/blip-v2/pkg/database"
	"github.com/scriptnsam/blip-v2/pkg/security"
	"github.com/scriptnsam/blip-v2/pkg/utils"
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

	// connect to sqlLite
	sqlDb,err:=database.SqLite();
	if err != nil {
		return "", err
	}
	
	defer sqlDb.Close()

	// Crete authentication table if not exists
	_, err = sqlDb.Exec(
		`CREATE TABLE IF NOT EXISTS authentication (
			id INTEGER PRIMARY KEY,
			user_id INTEGER NOT NULL,
			token TEXT NOT NULL
		)`,
	)
	if err != nil {
		return "", err
	}

	// GENERATE TOKEN
	token, err:=utils.GenerateToken(username,password)
	if err != nil {
		return "", err
	}

	// Insert token into authentication table
	_, err = sqlDb.Exec(
		"INSERT INTO authentication (user_id, token) VALUES (?, ?)",userId,token)

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