package api

import (
	"log"

	"github.com/Ahmad-Mosha/inventory-system/internal/repository"
)

func main() {
	db , err := repository.InitDB("inventory.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	

}