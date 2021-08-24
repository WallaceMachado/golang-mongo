package main

import (
	"fmt"
	"golang-mongo/models"
	"golang-mongo/services"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Rodando a API")

	user := models.User{

		Name:      "Jesus",
		Email:     "jesus.matiz@micorreo.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := services.Create(user)

	fmt.Println(err)

	if err != nil {
		fmt.Println("La prueba de persistencia de datos del usuario a fallado")

	} else {
		fmt.Println("La prueba finalizo con exito!")
	}

	log.Fatal(http.ListenAndServe(":5000", nil))
}
