package rest

import (
	"log"
	"net/http"

	"expense-management-system/internal/api/rest/handlers"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

func NewServer() *Server {
	router := mux.NewRouter()
	return &Server{router: router}
}

func (s *Server) InitializeRoutes() {
	s.router.HandleFunc("/api/expenses", handlers.GetExpenses).Methods(http.MethodGet)
	s.router.HandleFunc("/api/expenses", handlers.CreateExpense).Methods(http.MethodPost)
	s.router.HandleFunc("/api/expenses/{id}", handlers.GetExpense).Methods(http.MethodGet)
	s.router.HandleFunc("/api/expenses/{id}", handlers.UpdateExpense).Methods(http.MethodPut)
	s.router.HandleFunc("/api/expenses/{id}", handlers.DeleteExpense).Methods(http.MethodDelete)
}

func (s *Server) Start(addr string) {
	log.Printf("Starting REST server on %s", addr)
	if err := http.ListenAndServe(addr, s.router); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
