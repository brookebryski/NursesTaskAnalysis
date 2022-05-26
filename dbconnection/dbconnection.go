package dbconnection

import (
	// ...
	"fmt"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db   *gorm.DB // upointer to gorm.DB
	once sync.Once
)

// Driver for dbconnection
type Driver string

// Drivers

const (
	MySQL    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

//New creates de connection with the database
func New(d Driver) {
	switch d {
	case MySQL:
		newMySQLDB()
	case Postgres:
		newPostgresDB()
	}
}

func newPostgresDB() {
	once.Do(func() {
		var err error
		db, err = gorm.Open("postgres", "postgres://gylgtzme:iC0YHBEmHCSC3JskSiKRhUEcqoLclOGA@fanny.db.elephantsql.com/gylgtzme")
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		fmt.Println("Connected to postgres")
	})
}

func newMySQLDB() {
	once.Do(func() {
		var err error
		db, err = gorm.Open("mysql", "")
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		fmt.Println("Connected to mySQL")
	})
}

// DB return a unique instance of db
func DB() *gorm.DB { // returns a pointer to gorm.DB
	return db
}
