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
	db   *gorm.DB // un puntero a gorm.DB
	once sync.Once
)

// ...

func newPostgresDB() {
	once.Do(func() {
		var err error
		db, err = gorm.Open("postgres", "postgres://edteam:edteam@localhost:7530/godb?sslmode=disable")
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		fmt.Println("conectado a postgres")
	})
}

// func newMySQLDB() {
// 	once.Do(func() {
// 		var err error
// 		db, err = gorm.Open("mysql", "edteam:edteam@tcp(localhost:7531)/godb?parseTime=true")
// 		if err != nil {
// 			log.Fatalf("can't open db: %v", err)
// 		}

// 		fmt.Println("conectado a mySQL")
// 	})
// }

// DB return a unique instance of db
func DB() *gorm.DB { // retorna un puntero a gorm.DB
	return db
}
