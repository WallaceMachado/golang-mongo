package services_test

import (
	"golang-mongo/models"
	"golang-mongo/services"
	"testing"
	"time"
)

var userId string

func TestCreate(t *testing.T) {

	user := models.User{

		Name:      "Jesus",
		Email:     "jesus.matiz@micorreo.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := services.Create(user)

	if err != nil {
		t.Error("La prueba de persistencia de datos del usuario a fallado")
		t.Fail()
	} else {
		t.Log("La prueba finalizo con exito!")
	}
}
