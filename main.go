package main

import (
	"fmt"

	"github.com/api_rest_go/models"
	"github.com/api_rest_go/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "Historia 1"},
		{Nome: "Nome 2", Historia: "Historia 2"},
	}

	fmt.Printf("Iniciando o servidor Rest com Go")
	routes.HandleResquest()
}
