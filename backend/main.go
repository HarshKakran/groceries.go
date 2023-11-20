package main

import (
	"fmt"

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

	fmt.Printf("Server is now running on port %s\n", port)
	fmt.Println("Starting GO API service...")
	fmt.Println(`
 ______     ______        ______     ______   __    
/\  ___\   /\  __ \      /\  __ \   /\  == \ /\ \   
\ \ \__ \  \ \ \/\ \     \ \  __ \  \ \  _-/ \ \ \  
 \ \_____\  \ \_____\     \ \_\ \_\  \ \_\    \ \_\ 
  \/_____/   \/_____/      \/_/\/_/   \/_/     \/_/ `)

	// This is to keep the main goroutine from exiting immediately
	select {}
}
