package filltables

import (
	"NurseTasks/dbconnection"
	"NurseTasks/models"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func FillTables() {

	HelenPassword := "Helen123"
	hashedHelenPassword, _ := HashPassword(HelenPassword)

	PatriciaPassword := "Patricia123"
	hashedPatriciaPassword, _ := HashPassword(PatriciaPassword)

	NancyPassword := "Nancy123"
	hashedNancyPassword, _ := HashPassword(NancyPassword)

	BrookePassword := "Brooke123"
	hashedBrookePassword, _ := HashPassword(BrookePassword)

	User_Nurse := models.User{
		Name:         "Helen",
		Password:     hashedHelenPassword,
		IsNurse:      "True",
		IsSupervisor: "False",
		RegionID:     1,
		FacilityID:   2,
		IsLoggedIn:   "False",
	}

	User_Supervisor_Nurse := models.User{
		Name:         "Patricia",
		Password:     hashedPatriciaPassword,
		IsNurse:      "True",
		IsSupervisor: "True",
		RegionID:     1,
		FacilityID:   1,
		IsLoggedIn:   "False",
	}

	User_Supervisor1 := models.User{
		Name:         "Nancy",
		Password:     hashedNancyPassword,
		IsNurse:      "False",
		IsSupervisor: "True",
		RegionID:     2,
		FacilityID:   1,
		IsLoggedIn:   "False",
	}
	User_Supervisor2 := models.User{
		Name:         "Brooke",
		Password:     hashedBrookePassword,
		IsNurse:      "False",
		IsSupervisor: "True",
		RegionID:     2,
		FacilityID:   2,
		IsLoggedIn:   "False",
	}

	task1 := models.Task{
		Name:         "Dispensing medications",
		DepartmentID: 1,
	}
	task2 := models.Task{
		Name:         "Reviewing prescriptions for accuracy",
		DepartmentID: 1,
	}
	task3 := models.Task{
		Name:         "Checking for drug interactions",
		DepartmentID: 1,
	}
	task4 := models.Task{
		Name:         "Blood work",
		DepartmentID: 2,
	}
	task5 := models.Task{
		Name:         "Checking ldl cholesterol levels",
		DepartmentID: 2,
	}
	task6 := models.Task{
		Name:         "Sterilize the lab instruments",
		DepartmentID: 2,
	}
	task7 := models.Task{
		Name:         "Performing antigen test",
		DepartmentID: 3,
	}
	task8 := models.Task{
		Name:         "Performing PCR test",
		DepartmentID: 3,
	}
	task9 := models.Task{
		Name:         "Sending Covid test's results",
		DepartmentID: 3,
	}

	dep1 := models.Department{
		Name:       "Farmacy",
		FacilityID: 1,
	}
	dep2 := models.Department{
		Name:       "Lab",
		FacilityID: 1,
	}
	dep3 := models.Department{
		Name:       "CovidTest",
		FacilityID: 1,
	}
	dep4 := models.Department{
		Name:       "Farmacy",
		FacilityID: 2,
	}
	dep5 := models.Department{
		Name:       "Lab",
		FacilityID: 2,
	}
	dep6 := models.Department{
		Name:       "CovidTest",
		FacilityID: 2,
	}
	dep7 := models.Department{
		Name:       "Farmacy",
		FacilityID: 3,
	}
	dep8 := models.Department{
		Name:       "Lab",
		FacilityID: 3,
	}
	dep9 := models.Department{
		Name:       "CovidTest",
		FacilityID: 3,
	}

	dep10 := models.Department{
		Name:       "Farmacy",
		FacilityID: 4,
	}
	dep11 := models.Department{
		Name:       "Lab",
		FacilityID: 4,
	}
	dep12 := models.Department{
		Name:       "CovidTest",
		FacilityID: 4,
	}
	fac1 := models.Facility{
		Name:     "Florida Medical Center",
		RegionID: 1,
	}
	fac2 := models.Facility{
		Name:     "New York Medical Center",
		RegionID: 1,
	}
	fac3 := models.Facility{
		Name:     "Holywood Medical Center",
		RegionID: 2,
	}
	fac4 := models.Facility{
		Name:     "Napa Medical Center",
		RegionID: 2,
	}
	reg1 := models.Region{
		Name: "East",
	}

	reg2 := models.Region{
		Name: "West",
	}

	dbconnection.DB().Create(&task1)
	dbconnection.DB().Create(&task2)
	dbconnection.DB().Create(&task3)
	dbconnection.DB().Create(&task4)
	dbconnection.DB().Create(&task5)
	dbconnection.DB().Create(&task6)
	dbconnection.DB().Create(&task7)
	dbconnection.DB().Create(&task8)
	dbconnection.DB().Create(&task9)
	dbconnection.DB().Create(&User_Nurse)
	dbconnection.DB().Create(&User_Supervisor_Nurse)
	dbconnection.DB().Create(&User_Supervisor1)
	dbconnection.DB().Create(&User_Supervisor2)
	dbconnection.DB().Create(&dep1)
	dbconnection.DB().Create(&dep2)
	dbconnection.DB().Create(&dep3)
	dbconnection.DB().Create(&dep4)
	dbconnection.DB().Create(&dep5)
	dbconnection.DB().Create(&dep6)
	dbconnection.DB().Create(&dep7)
	dbconnection.DB().Create(&dep8)
	dbconnection.DB().Create(&dep9)
	dbconnection.DB().Create(&dep10)
	dbconnection.DB().Create(&dep11)
	dbconnection.DB().Create(&dep12)
	dbconnection.DB().Create(&fac1)
	dbconnection.DB().Create(&fac2)
	dbconnection.DB().Create(&fac3)
	dbconnection.DB().Create(&fac4)
	dbconnection.DB().Create(&reg1)
	dbconnection.DB().Create(&reg2)

}
