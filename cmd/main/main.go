package main

import(
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/qussaikh/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	// Get port from environment (Render sets this)
	port := os.Getenv("PORT")
	if port == "" {
		port = "9010" // fallback for local dev
	}

	log.Printf("Server starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, r))
}
