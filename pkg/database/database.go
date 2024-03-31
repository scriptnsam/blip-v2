package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	_ "modernc.org/sqlite" // import sqlite3 driver for mysql and sqlite
)
type Database struct {
	*sql.DB
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database     string
}

// connect to the databse
func Connect() (*Database, error) {
	config :=DBConfig{
		Host: viper.GetString("MYSQL_HOST"),
		Port: viper.GetString("MYSQL_PORT"),
		User: viper.GetString("MYSQL_USERNAME"),
		Password: viper.GetString("MYSQL_PASSWORD"),
		Database: viper.GetString("DB_NAME"),
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.User, config.Password, config.Host, config.Port, config.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("\nerror occured connecting to the database.\nPlease try again")
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("\nerror occured connecting to the database.\nPlease try again")
	}

	// create necessary tables
	if err := CreateTables(&Database{db}); err != nil {
		return nil, err
	}
	return &Database{db}, nil
}

func CreateTables(db *Database) error {
	statements := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			username VARCHAR(50) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL,
			email VARCHAR(100) UNIQUE,
			full_name VARCHAR(100),
			role ENUM('admin', 'user') DEFAULT 'user',
			date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			last_login TIMESTAMP,
			active BOOLEAN DEFAULT true
		)`,
		`CREATE TABLE IF NOT EXISTS tools(
			id INT AUTO_INCREMENT PRIMARY KEY,
			user_id INT NOT NULL,
			name VARCHAR(50) NOT NULL UNIQUE,
			group_name VARCHAR(50) NOT NULL,
			os_type VARCHAR(50) NOT NULL,
			download_link VARCHAR(100) NOT NULL,
			file_extension VARCHAR(10) NOT NULL,
			date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		)`,
		`CREATE TABLE IF NOT EXISTS tool_groups(
			id INT AUTO_INCREMENT PRIMARY KEY,
			user_id INT NOT NULL,
			name VARCHAR(50) NOT NULL,
			os_type VARCHAR(50) NOT NULL,
			date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		)`,
		`CREATE TABLE IF NOT EXISTS os_types(
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(50) NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS logs(
			id INT AUTO_INCREMENT PRIMARY KEY,
			user_id INT NOT NULL,
			tool_id INT NOT NULL,
			message VARCHAR(255) NOT NULL,
			date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
			FOREIGN KEY (tool_id) REFERENCES tools(id) ON DELETE CASCADE
		)`,

		// `INSERT INTO users (username, password, email, full_name, role) VALUES ('scriptnsam', '$2a$10$LIj00tP8pX3v9GOKqh07HuQpGrHqKR.BSSTH.DZwDPJbgt4jk9IVW', 'oluwafemisam40@gmail.com', 'Samuel Oluwafemi', 'user')`,
		// `INSERT INTO tool_groups (user_id,name,os_type) VALUES (1,"Plan B","windows")`,
		// `INSERT INTO tools (user_id, name, group_name, os_type, download_link,file_extension) VALUES (1,"cpu-z","New","windows","https://download.cpuid.com/cpu-z/cpu-z_2.09-en.exe","exe")`,
	}

	for _,stmt:=range statements{
		_, err := db.Exec(stmt)
		if err != nil {
			return fmt.Errorf("\nerror occured creating the table.\nPlease try again")
		}

	}

	return nil
}

func General()(*Database,error){
	db,err:=Connect()
	if err != nil {
		return db,err
	}
	// defer db.Close()
	return db,nil
}

func SqLite()(*sql.DB,error){
	db,err:=sql.Open("sqlite","./blip.db")
	if err!=nil{
		return nil,fmt.Errorf("\nerror occured creating the database file.\nPlease try again")
	}

	return db,nil
}