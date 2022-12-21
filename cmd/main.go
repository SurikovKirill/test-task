package main

import (
	"net/http"
	_ "test-task/docs"
	"test-task/internal/routes"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @version 1.0
// @in header
// @host localhost:8080
// @Test task
// @BasePath /

func main() {
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	http.HandleFunc("exchange/", routes.Get)
	http.ListenAndServe(":8080", nil)
}
