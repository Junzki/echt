package models

type Content struct {
	ID    int     	`gorm:"primary_key"`
	Title string	`gorm:"type:varchar(128),default:''"`
	Intro string	`gorm:"default:''"`
	Link  string	`gorm:"not null,varchar(255)"`
}
