package psql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go-hex-sample/pkg/ink"
)

const (
	host     = "127.0.0.1"
	port     = "5432"
	dbName   = "inkdb"
	user     = "test"
	password = "test"
)

func Migrate() error {
	db, err := openDb()
	if err != nil {
		return err
	}
	db.AutoMigrate(&ink.Ink{})
	if err := db.Close(); err != nil {
		return err
	}
	return nil
}

func openDb() (*gorm.DB, error) {
	args := fmt.Sprintf("host=%v port=%v dbname=%v user=%v password=%v sslmode=disable",
		host, port, dbName, user, password)
	return gorm.Open("postgres", args)
}
