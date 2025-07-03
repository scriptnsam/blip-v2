package models

type App struct {
	Name    string `json:"name"`
	Source  string `json:"source"`
	Command string `json:"command"`
}
