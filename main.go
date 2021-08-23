package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rodando a API")

	log.Fatal(http.ListenAndServe(":5000", nil))
}
