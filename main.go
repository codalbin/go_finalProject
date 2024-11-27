package main

import (
	"go-final-project/config"
	"go-final-project/controllers/homecontroller"
	"go-final-project/controllers/categorycontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	http.HandleFunc("/", homecontroller.Welcome)
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	log.Println("Server runing on port 8080")
	http.ListenAndServe(":8080", nil)
}
