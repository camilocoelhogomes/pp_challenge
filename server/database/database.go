package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseFacotory struct {
	db *gorm.DB
}

func (d *DatabaseFacotory) GetDataBase() gorm.DB {

	if d.db == nil {
		dsn := "host=localhost port=5432 user=postgres password=admin dbname=postgres"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatalln(err)
		}

		d.db = db
	}

	return *d.db
}
