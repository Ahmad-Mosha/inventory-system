package main

import (
	"log"

	"github.com/Ahmad-Mosha/inventory-system/internal/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := repository.InitDB("inventory.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default()
	router.Run()
}
