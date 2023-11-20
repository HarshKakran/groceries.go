package routers

import (
	"github.com/HarshKakran/groceries.go/handlers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// common routes
	router.HandleFunc("/api/signin", handlers.Signin).Methods("POST")
	router.HandleFunc("/api/signup", handlers.Signup).Methods("POST")
	router.HandleFunc("/api/signout", handlers.Signout).Methods("GET")

	// owner routes
	router.HandleFunc("/api/store/", handlers.StoreInfo).Methods("GET")
	router.HandleFunc("/api/store/item", handlers.CreateItem).Methods("POST")
	router.HandleFunc("api/store/item/{id}", handlers.UpdateItem).Methods("PUT")
	router.HandleFunc("api/store/item/{id}", handlers.DeleteItem).Methods("DELETE")

	// consumer routes
	router.HandleFunc("api/consumer", handlers.ConsumerPage).Methods("GET")
	router.HandleFunc("api/consumer/checkout", handlers.CreateOrder).Methods("POST")
}
