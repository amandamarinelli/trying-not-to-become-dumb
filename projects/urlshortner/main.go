package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/amandamarinelli/trying-not-to-become-dumb/projects/urlshortner/internal"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	clientDB := initDB()
	mux := http.NewServeMux()

	mux.HandleFunc("/", helloWordHandler)

	handleCustomRoutes(mux, internal.GetRoutesFromDB(clientDB))

	fmt.Println("Running server on http://localhost:8080")
	http.ListenAndServe(":8080", mux)

}

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./my-database.db")
	if err != nil {
		log.Fatal(err, "error initializing DB")
	}

	internal.CreateRoutesTable(db)
	internal.PopulateRoutes(db)

	return db
}

func handleCustomRoutes(mux *http.ServeMux, routes []internal.Route) {
	for _, r := range routes {
		if r.Path == "" || r.URL == "" {
			//ignore wrong routes
			continue
		}

		mux.HandleFunc(r.Path, redirectTO(r.URL))
	}
}

func helloWordHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func redirectTO(url string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, url, http.StatusPermanentRedirect)
	}
}
