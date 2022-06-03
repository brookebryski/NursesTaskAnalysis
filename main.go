package main

import (
	"NursesTaskAnalysis/dbconnection"
	"NursesTaskAnalysis/filltables"
	"NursesTaskAnalysis/router"
)

func main() {
	driver := dbconnection.Postgres
	dbconnection.New(driver)
	//dbconnection.DB().AutoMigrate(&models.User{}, &models.Facility{}, &models.Region{}, &models.Department{}, &models.Task{}, &models.TaskEntered{})
	router.Routing()
	filltables.FillTables()

}
