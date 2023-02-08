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
	// user authentication middleware

	return router
}
