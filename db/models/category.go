package models

type Category struct {
	Id       uint   `json:"id" gorm:"primary_key;autoIncrement"`
	Men      string `json:"men"`
	Women    string `json:"women"`
	Children string `json:"children"`
}
