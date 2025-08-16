package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dom-m17/lms/backend/internal/competition"
	"github.com/dom-m17/lms/backend/internal/db"
	"github.com/dom-m17/lms/backend/internal/entry"
	"github.com/dom-m17/lms/backend/internal/match"
	"github.com/dom-m17/lms/backend/internal/selection"
	"github.com/dom-m17/lms/backend/internal/subgraph"
	graphresolvers "github.com/dom-m17/lms/backend/internal/subgraph/resolvers"
	"github.com/dom-m17/lms/backend/internal/team"
	"github.com/dom-m17/lms/backend/internal/user"
	"github.com/rs/cors"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func main() {
	_ = godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	conn, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_SOURCE"))
	if err != nil {
		log.Fatalf("could not connect to the database: %v", err)
	}

	querier := db.New(conn)

	srv := handler.New(subgraph.NewExecutableSchema(subgraph.Config{
		Resolvers: &graphresolvers.Resolver{
			TeamService:        team.NewService(querier),
			CompetitionService: competition.NewService(querier),
			UserService:        user.NewService(querier),
			MatchService:       match.NewService(querier),
			SelectionService:   selection.NewService(querier),
			EntryService:       entry.NewService(querier),
		},
	}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	handlerWithCORS := cors.AllowAll().Handler(mux)

	//! Temporary code to populate the DB
	//! This will eventually be a cron job or called from the FE
	// ctx := context.Background()
	// footballDataSvc := footballdata.New(querier)
	// err = footballDataSvc.PopulateTeams(ctx)
	// if err != nil {
	// 	log.Printf("error populating teams: %v", err)
	// }
	// footballDataSvc.PopulateMatches(ctx)
	// if err != nil {
	// 	log.Printf("error populating matches: %v", err)
	// }
	//!

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, handlerWithCORS))
}
