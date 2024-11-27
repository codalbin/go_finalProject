package main

import (
	"go-final-project/config"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()
	
	log.Println("Server runing on port 8080")
	http.ListenAndServe(":8080", nil)
}
