package main

import (
	"expense-management-system/internal/api/rest"
)

func main() {
    server := rest.NewServer()
    server.InitializeRoutes()

    server.Start(":8080")
	// // Initialize the router
	// r := mux.NewRouter()

	// // Set up REST API
	// rest.SetupRoutes(r)

	// // Set up GraphQL server
	// graphql.SetupServer(r)

	// // Set up gRPC server
	// go grpc.StartServer()

	// // Middleware for logging and authentication
	// r.Use(middleware.LoggingMiddleware)
	// r.Use(middleware.AuthMiddleware)

	// // Start the HTTP server
	// log.Println("Starting server on :8080")
	// if err := http.ListenAndServe(":8080", r); err != nil {
	// 	log.Fatalf("Could not start server: %s\n", err)
	// }
}
