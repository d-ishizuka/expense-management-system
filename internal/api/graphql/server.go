package graphql

import (
	"github.com/graphql-go/graphql"
	"net/http"
	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

func NewServer() *Server {
	router := mux.NewRouter()
	return &Server{router: router}
}

func (s *Server) SetupRoutes() {
	s.router.HandleFunc("/graphql", s.handleGraphQL()).Methods("POST")
}

func (s *Server) handleGraphQL() http.HandlerFunc {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
	if err != nil {
		panic(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// GraphQLのリクエストを処理するロジックをここに追加
	}
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.router)
}