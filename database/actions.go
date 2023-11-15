package database

import "go-http/models"

func GetUserById(id int) (models.User, error) {
	var data models.User
	data.Id = id
	err := Db.Model(&data).WherePK().Select()
	return data, err
}

func GetEntitiesFromDb[T any]() (T, error) {
	var data T
	err := Db.Model(&data).Select()
	return data, err
}
