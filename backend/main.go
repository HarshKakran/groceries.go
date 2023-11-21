package main

import (
	"fmt"
	"log"

	"github.com/HarshKakran/groceries.go/handlers"
	"github.com/HarshKakran/groceries.go/models"
	"github.com/HarshKakran/groceries.go/routers"
)

func main() {
	r := routers.Router()
	port := ":8080"

	// Start the server
	go func() {
		err := r.Run(port)
		if err != nil {
			fmt.Printf("Error starting server: %v\n", err)
		}
	}()

	fmt.Printf("Starting GO API service on port %s\n", port)
	fmt.Println(`
 ______     ______        ______     ______   __    
/\  ___\   /\  __ \      /\  __ \   /\  == \ /\ \   
\ \ \__ \  \ \ \/\ \     \ \  __ \  \ \  _-/ \ \ \  
 \ \_____\  \ \_____\     \ \_\ \_\  \ \_\    \ \_\ 
  \/_____/   \/_____/      \/_/\/_/   \/_/     \/_/ `)

	// This is to keep the main goroutine from exiting immediately

	err := models.Load(handlers.USERS_DATA_FILE, &models.Users)
	if err != nil {
		log.Fatalf("error while loading users data. %v", err)
	}

	err = models.Load(handlers.ORDERS_DATA_FILE, &models.ConsumerOrdersMapping)
	if err != nil {
		log.Fatalf("error while loading orders data. %v", err)
	}
	fmt.Println("data loaded")
	// fmt.Println(models.Users, "\n", models.ConsumerOrdersMapping)

	select {}
}
