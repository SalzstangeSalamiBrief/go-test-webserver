package database

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"go-http/models"
)

var Db *pg.DB

func ConnectToDb(username string, password string, database string) {
	options := pg.Options{
		Addr:     ":5432",
		User:     username,
		Password: password,
		Database: database,
	}

	Db = pg.Connect(&options)
}

func CreateSchema() error {
	err := Db.Model((*models.User)(nil)).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
	return err
}

func SeedData() error {
	var users []models.User
	Db.Model(&users).Select()
	if len(users) > 0 { // check if the data should be seeded
		return nil
	}

	user := &models.User{FirstName: "Udo", LastName: "Peters", Age: 82}
	_, insertError := Db.Model(user).Insert()
	return insertError
}
