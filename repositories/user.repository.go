package repositories

import (
	"context"
	"golang-mongo/database"
	"golang-mongo/models"
)

var collection = database.GetCollection("users")
var ctx = context.Background()

func Create(user models.User) error {

	var err error

	_, err = collection.InsertOne(ctx, user)

	if err != nil {
		return err
	}

	return nil

}

func Read() (models.Users, error) {

	return nil, nil
}

func Update(user models.User, userId string) error {

	return nil
}

func Delete(userId string) error {

	return nil
}
