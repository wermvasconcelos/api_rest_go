package main

import (
	"api_rest_go/models"
	"api_rest_go/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	fmt.Printf("Iniciando o servidor Rest com Go")
	routes.HandleResquest()
}
