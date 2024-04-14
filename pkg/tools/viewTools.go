package tools

import (
	"fmt"
	"runtime"

	"github.com/scriptnsam/blip-v2/pkg/authentication"
	"github.com/scriptnsam/blip-v2/pkg/database"
)

type Tools struct {
	Name         string
	PackageName string
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

	rows, err := db.Query("SELECT name,package_name,group_name,date_created FROM tools WHERE user_id=?",userId)
	if err != nil {
		return []Tools{}, err
	}
	defer rows.Close()

	var tools []Tools
	for rows.Next() {
		var tool Tools
		err := rows.Scan(&tool.Name, &tool.PackageName, &tool.Group, &tool.DateCreated)
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

func ViewToolsByGroup(group string) ([]Tools, error) {
	db,err:=database.General()
	if err!=nil{
		return []Tools{},err
	}

	// Get userId
	authStatus,userId:=authentication.IsLoggedIn()

	if !authStatus {
		return []Tools{},fmt.Errorf("\nNot Authenticated\nPlease run the login command `blip me login -u <username> -p <password>`")
	}

	rows, err := db.Query("SELECT name, package_name, group_name, date_created FROM tools WHERE user_id=? AND group_name=?", userId, group)
	if err != nil {
		return []Tools{}, err
	}
	defer rows.Close()

	var tools []Tools
	for rows.Next() {
		var tool Tools
		err := rows.Scan(&tool.Name, &tool.PackageName, &tool.Group, &tool.DateCreated)
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

type Group struct {
	Name         string
	DateCreated string
}

func ViewGroups()([]Group, error){
	db,err:=database.General()
	if err!=nil{
		return []Group{},err
	}

	// Get userId
	authStatus,userId:=authentication.IsLoggedIn()

	if !authStatus {
		return []Group{},fmt.Errorf("\nNot Authenticated\nPlease run the login command `blip me login -u <username> -p <password>`")
	}

	os:=runtime.GOOS;

	rows, err:=db.Query("SELECT name,date_created FROM tool_groups WHERE user_id=? AND os_type=?", userId,os)
	if err != nil {
		return []Group{}, err
	}
	defer rows.Close()

	var groups []Group
	for rows.Next() {
		var group Group
		err := rows.Scan(&group.Name, &group.DateCreated)
		if err != nil {
			return []Group{}, err
		}
		groups = append(groups, group)
	}
	if err = rows.Err(); err != nil {
		return []Group{}, err
	}

	return groups, nil
}