package main

import (
	_ "github.com/lib/pq"
	"golang-restful-api-2/helper"
	"golang-restful-api-2/middleware"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}
func main() {
	server := InitializedServer()
	err := server.ListenAndServe()
	helper.PanicIferror(err)
}
