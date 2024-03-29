package authentication

import (
	"database/sql"

	"github.com/scriptnsam/blip-v2/pkg/database"
	"github.com/scriptnsam/blip-v2/pkg/security"
	"github.com/scriptnsam/blip-v2/pkg/utils"
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
    
		dbId int

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
	
		sqliteStmt:=[]string{
			`CREATE TABLE IF NOT EXISTS authentication (
			id INTEGER PRIMARY KEY,
			user_id INTEGER NOT NULL,
			token TEXT NOT NULL,
			time DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TRIGGER IF NOT EXISTS update_time_trigger 
		AFTER UPDATE ON authentication 
		FOR EACH ROW 
		BEGIN
			UPDATE authentication SET time = CURRENT_TIMESTAMP WHERE id = OLD.id;
		END;`,
	}

		for _, stmt:=range sqliteStmt {
			_, err = sqlDb.Exec(stmt)
			if err != nil {
				return "", err
			}
		}

	// GENERATE TOKEN
	token, err:=utils.GenerateToken(username,password)
	if err != nil {
		return "", err
	}

	// Check if user's details is already stored in the db. If it's already stored, Just update the token
	sqlDb.QueryRow("SELECT id FROM authentication WHERE user_id = ?", id).Scan(&dbId)
	if dbId != 0 {
		_, err = sqlDb.Exec("UPDATE authentication SET token = ? WHERE user_id = ?", token, id)
		if err != nil {
			return "", err
		}
		return "Logged in (from update)", nil
	}

	// Insert token into authentication table
	_, err = sqlDb.Exec(
		"INSERT INTO authentication (user_id, token) VALUES (?, ?)",id,token)
		if err != nil {
			return "", err
		}

	return "Logged in", nil
}

func IsLoggedIn() (bool, int) {
	sqlDb,err:=database.SqLite();
	if err != nil {
		return false, 0
	}

	defer sqlDb.Close()

	var (
		token string
		userId int
	)

	// get logged in user id
	if err:=sqlDb.QueryRow("SELECT user_id, token FROM authentication").Scan(&userId,&token);err!=nil{
		return false, 0
	}

	// check if token is valid
	_,err=utils.ParseToken(token)
	if err != nil {
		return false, 0
	}
	
	return true, userId
}