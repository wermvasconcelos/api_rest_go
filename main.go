package main

import (
	"fmt"

	"github.com/api_rest_go/routes"
)

func main() {
	fmt.Printf("Iniciando o servidor Rest com Go")
	routes.HandleResquest()
}
