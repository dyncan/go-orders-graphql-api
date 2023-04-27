package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dyncan/go-orders-graphql-api/bootstrap"
	"github.com/dyncan/go-orders-graphql-api/config"
	"github.com/dyncan/go-orders-graphql-api/graph"
	"github.com/dyncan/go-orders-graphql-api/graph/generated"
	c "github.com/dyncan/go-orders-graphql-api/pkg/config"
	"log"
	"net/http"
)

func init() {
	// 初始化配置信息
	config.Initialize()
}

func main() {
	// 初始化 SQL
	db := bootstrap.SetupDB()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB: db,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", c.GetString("app.port"))
	log.Fatal(http.ListenAndServe(":"+c.GetString("app.port"), nil))
}