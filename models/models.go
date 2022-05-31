package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"type:varchar(20); not null"`
	Password     string `gorm:"type:varchar(100); not null"`
	IsNurse      string `gorm:"not null"`
	IsSupervisor string `gorm:"not null"`
	RegionID     uint   `gorm:"not null"`
	FacilityID   uint   `gorm:"not null"`
	IsLoggedIn   string `gorm:"not null"`
}

type Region struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(20); not null"`
	//Facilities []Facility `gorm:"not null"`
}

type Facility struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(30); not null"`
	RegionID uint   `gorm:"not null"`
	//Departments []Department `gorm:"not null"`
}

type Department struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"type:varchar(40); not null"`
	FacilityID uint   `gorm:"not null"`
	//Tasks      []Task
}

type Task struct {
	ID           uint   `json:"id"`   //`gorm:"primaryKey"`
	Name         string `json:"name"` //`gorm:"type:varchar(50); not null"`
	DepartmentID uint   `json:"department_id"`
}

type TaskEntered struct {
	gorm.Model
	TaskID uint `gorm:"not null"`
	UserID uint `gorm:"not null"`
}


