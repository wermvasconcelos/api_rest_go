package routes

import (
	"api_rest_go/controllers"
	"log"
	"net/http"
)

func HandleResquest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
