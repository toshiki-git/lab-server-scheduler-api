package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/toshiki-git/lab-server-scheduler-api/db"
	"github.com/toshiki-git/lab-server-scheduler-api/model"
	"github.com/toshiki-git/lab-server-scheduler-api/repository"
)

func userHandler(repo *repository.UserRepository, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		createUser(repo, w, r)
	case "GET":
		readUser(repo, w, r)
	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

func createUser(repo *repository.UserRepository, w http.ResponseWriter, r *http.Request) {
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

func readUser(repo *repository.UserRepository, w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	user, err := repo.Read(r.Context(), userID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func main() {
	db := db.InitDB()
	repo := repository.NewUserRepository(db)

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		userHandler(repo, w, r)
	})

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
