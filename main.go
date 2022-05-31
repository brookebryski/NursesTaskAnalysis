package main

import (
	"NurseTasks/dbconnection"
	routePack "NurseTasks/router"
)

func main() {
	driver := dbconnection.Postgres
	dbconnection.New(driver)
	//dbconnection.DB().AutoMigrate(&models.User{}, &models.Facility{}, &models.Region{}, &models.Department{}, &models.Task{}, &models.TaskEntered{})
	routePack.Routing()
	//filltables.FillTables()

}
