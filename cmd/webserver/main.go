package main

import (
	"log"
	"net/http"

	"github.com/suleymanoff/kursach/internal/config"
	"github.com/suleymanoff/kursach/internal/handlers"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer db.Close()

	// Обработчики API
	http.HandleFunc("/api/register", handlers.RegisterHandler(db))
	http.HandleFunc("/api/login", handlers.LoginHandler(db))

	// Статические файлы
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Главная страница
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
