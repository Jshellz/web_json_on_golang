package models

type Product struct {
	ID            uint       `json:"id" gorm:"primary_key;autoIncrement"`
	Title         string     `json:"title"`
	Price         float32    `json:"price"`
	Photo         string     `json:"photo"`
	CategoryValue []Category `gorm:"foreignKey:Id;references:ID"`
}
