package authentication

import (
	"database/sql"

	"github.com/scriptnsam/blip-v2/pkg/database"
)

func Logout() (string, error) {
	sqlDb, err := database.SqLite()
	if err != nil {
		return "Could not open the database", err
	}
	defer sqlDb.Close()

	// 🔍 Check if the table exists
	var tableName string
	err = sqlDb.QueryRow(`
		SELECT name FROM sqlite_master
		WHERE type='table' AND name='authentication';
	`).Scan(&tableName)

	if err == sql.ErrNoRows {
		return "No login session found to log out from", nil
	} else if err != nil {
		return "Failed to verify authentication table", err
	}

	// ✅ Delete from the table
	_, err = sqlDb.Exec("DELETE FROM authentication")
	if err != nil {
		return "Failed to log out due to a database error", err
	}

	return "User logged out successfully", nil
}
