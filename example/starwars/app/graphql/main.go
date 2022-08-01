package main

import (
	"log"
	"net/http"

	"github.com/alehechka/grpc-graphql-gateway/example/starwars/spec/starwars"
	"github.com/alehechka/grpc-graphql-gateway/runtime"
	"github.com/friendsofgo/graphiql"
	"google.golang.org/grpc"
)

func main() {
	mux := runtime.NewServeMux(runtime.Cors())
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := starwars.RegisterStartwarsServiceGraphqlHandler(mux, nil, "localhost:50051", opts...); err != nil {
		log.Fatalln(err)
	}

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		log.Fatalln(err)
	}

	http.Handle("/graphql", mux)
	http.Handle("/graphiql", graphiqlHandler)
	log.Fatalln(http.ListenAndServe(":8888", nil))
}
