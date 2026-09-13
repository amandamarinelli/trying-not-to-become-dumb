package internal

import (
	"database/sql"
	"log"
)

type Route struct {
	Path string
	URL  string
}

func CreateRoutesTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS routes (
		path TEXT NOT NULL,
		url TEXT NOT NULL
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err, "error creating table")
	}
}

func PopulateRoutes(db *sql.DB) {
	var count int
	db.QueryRow("SELECT COUNT(*) FROM routes").Scan(&count)
	if count != 0 {
		return
	}

	query := `
		INSERT OR IGNORE INTO routes (path, url) 
		VALUES ("/good-news", "https://www.google.com/search?q=google+boas+noticias+pra+essa+semana&oq=google+boas+noticias+pra+essa+semana"), 
        ("/wrong","")
	`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err, "error populating table")
	}

}

func GetRoutesFromDB(db *sql.DB) []Route {
	var routes []Route
	query := `SELECT path, URL FROM routes`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err, "error getting routes")
	}

	defer rows.Close()

	for rows.Next() {
		var r Route
		err = rows.Scan(&r.Path, &r.URL)
		if err != nil {
			log.Fatal(err, "error scanning rows")
		}

		routes = append(routes, r)
	}

	return routes
}
