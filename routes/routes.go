package routes

import (
	"net/http"

	"ritsrnjn/dummy-cc/handlers"
	"ritsrnjn/dummy-cc/middlewares"

	"github.com/go-chi/chi"
)

func GetRouter() *chi.Mux {

	router := chi.NewRouter()

	// Add routes here
	router.Use(middlewares.AuthTokenMiddleware)

	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Welcome to the chi"))
	})
	router.Get("/hello", handlers.HelloFunc)

	// POST method to create account
	router.Post("/account", handlers.CreateAccount)
	// GET method to get account details
	router.Get("/account/{accountID}", handlers.GetAccount)

	// POST method to create offer
	router.Post("/offer", handlers.CreateOffer)
	// PATCH method to update offer
	// router.Patch("/offer/{offerID}", handlers.UpdateOffer)
	// LIST method to list offers for a given accountID
	router.Get("/account/{accountID}/offer", handlers.ListOffers)

	return router
}
