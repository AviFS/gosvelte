package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	base := filepath.Join("templates", "base.html")
	index := filepath.Join("templates", "index.html")

	tmpl := template.Must(template.ParseFiles(base, index))
	tmpl.ExecuteTemplate(w, "base", nil)
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	base := filepath.Join("templates", "base.html")
	index := filepath.Join("templates", "blog.html")

	tmpl := template.Must(template.ParseFiles(base, index))
	tmpl.ExecuteTemplate(w, "base", nil)
}

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Println("Error loading .env file")
	// }

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8080"
	// }
	port := "8080"

	// apiKey := os.Getenv("NEWS_API_KEY")
	// if apiKey == "" {
	// 	log.Fatal("NEWS_API_KEY in .env must be set")
	// }

	// fs := http.FileServer(http.Dir("assets"))

	mux := http.NewServeMux()
	mux.HandleFunc("/blog", blogHandler)
	mux.HandleFunc("/", indexHandler)
	// mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.ListenAndServe(":"+port, mux)

}
