package main

import (
	"fmt"
	"golang-mongo/services"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rodando a API")
	userId := "6124e0ca7526bc5b240d15fb"
	err := services.Delete(userId)

	if err != nil {
		fmt.Println("Error al tratar de eliminar el usuario")

	} else {
		fmt.Println("La prueba de eliminación finalizo con exito!")
	}

	log.Fatal(http.ListenAndServe(":5000", nil))
}
