package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/toshiki-git/lab-server-scheduler-api/db"
	"github.com/toshiki-git/lab-server-scheduler-api/model"
	"github.com/toshiki-git/lab-server-scheduler-api/repository"
)

func createUser(repo *repository.UserRepository, w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		return
	}

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	err = repo.Create(ctx, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func main() {
	db := db.InitDB()
	repo := repository.NewUserRepository(db)
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		createUser(repo, w, r)
	})
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
