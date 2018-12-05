package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

var glob string

func init() {
	argsWithoutProg := os.Args[1:]
	glob = argsWithoutProg[0]
}

func main() {

	s := `
				scalar Time
                schema {
                        query: Query
                }
                type Query {
                        photos: [Photo!]!
				}
				type Photo {
					id: ID!
					name: String!
					dateTime: Time!
					model: String
					make: String
					lensModel: String
					lensMake: String
					iso: Int
					focalLength: Float
					focalLengthIn35mmFilm: Int
					exposureTime: String
					fNumber: Float
					thumbnail: String
					file: String
				}
        `
	schema := graphql.MustParseSchema(s, &query{})
	r := chi.NewRouter()

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	r.Use(cors.Handler)

	r.Handle("/graphql", &relay.Handler{Schema: schema})
	r.HandleFunc("/img/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		http.ServeFile(w, r, photoMap[id])
	})
	log.Fatal(http.ListenAndServe(":8080", r))
}
