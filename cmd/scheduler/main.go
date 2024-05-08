package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/toshiki-git/lab-server-scheduler-api/db"
	"github.com/toshiki-git/lab-server-scheduler-api/handler"
	"github.com/toshiki-git/lab-server-scheduler-api/repository"
)

func main() {
	db := db.InitDB()
	repo := repository.NewUserRepository(db)

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		handler.UserHandler(repo, w, r)
	})

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
