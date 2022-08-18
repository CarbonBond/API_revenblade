package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/CarbonBond/API_revenblade/graph"
	"github.com/CarbonBond/API_revenblade/graph/generated"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"github.com/gorilla/websocket"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
  router := chi.NewRouter()

  router.Use(cors.New(cors.Options{
    AllowedOrigins: []string{"http://localhost:8080", "http://localhost:5173"},
    AllowCredentials: true,
    Debug: true,
  }).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

  srv.AddTransport(&transport.Websocket{
    Upgrader: websocket.Upgrader{
      CheckOrigin: func(r *http.Request) bool {
        return r.Host == "localhost"
      },
      ReadBufferSize: 1024,
      WriteBufferSize: 1024,
    },
  })

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

