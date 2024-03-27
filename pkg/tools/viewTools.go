package tools

import (
	"fmt"

	"github.com/scriptnsam/blip-v2/pkg/authentication"
	"github.com/scriptnsam/blip-v2/pkg/database"
)

type Tools struct {
	Name         string
	DownloadLink string
	Group        string
	DateCreated string
}

func ViewTools() ([]Tools,error) {
	db,err:=database.General()
	if err!=nil{
		return []Tools{},err
	}

	// Get userId
	authStatus,userId:=authentication.IsLoggedIn()

	if !authStatus {
		return []Tools{},fmt.Errorf("\nNot Authenticated\nPlease run the login command `blip me login -u <username> -p <password>`")
	}

	rows, err := db.Query("SELECT name,download_link,group_name,date_created FROM tools WHERE user_id=?",userId)
	if err != nil {
		return []Tools{}, err
	}
	defer rows.Close()

	var tools []Tools
	for rows.Next() {
		var tool Tools
		err := rows.Scan(&tool.Name, &tool.DownloadLink, &tool.Group, &tool.DateCreated)
		if err != nil {
			return []Tools{}, err
		}
		tools = append(tools, tool)
	}

	if err = rows.Err(); err != nil {
		return []Tools{}, err
	}
	return tools, nil
}