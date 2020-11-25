package models

type Issue struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	ToolsID string `json:"toolsid"`
	IssueID string `json:"issueid"`
}
