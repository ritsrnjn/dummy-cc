package main

import (
	"net/http"

	"ritsrnjn/dummy-cc/routes"
)

func main() {
	router := routes.GetRouter()
	http.ListenAndServe(":3000", router)
}
