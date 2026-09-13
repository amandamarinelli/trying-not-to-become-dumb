package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", helloWordHandler)

	handleCustomRoutes(mux)

	fmt.Println("Running server on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}

func handleCustomRoutes(mux *http.ServeMux) {
	routerMap := map[string]string{
		"/good-news": "https://www.google.com/search?q=google+boas+noticias+pra+essa+semana&oq=google+boas+noticias+pra+essa+semana",
		"/wrong":     "",
	}

	for r, url := range routerMap {
		if r == "" || url == "" {
			//ignore wrong routes
			continue
		}

		mux.HandleFunc(r, redirectTO(url))
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
