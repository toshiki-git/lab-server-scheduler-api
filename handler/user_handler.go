package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/toshiki-git/lab-server-scheduler-api/middleware"
	"github.com/toshiki-git/lab-server-scheduler-api/model"
	"github.com/toshiki-git/lab-server-scheduler-api/repository"
)

func UserHandler(repo *repository.UserRepository, w http.ResponseWriter, r *http.Request) {
	middleware.EnableCORS(&w)
	if r.Method == "OPTIONS" {
		return
	}

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

	// ユーザーが既に存在するか確認
	exists, err := repo.Exists(ctx, user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if exists {
		// 既に存在する場合は何もせずに成功ステータスを返す
		w.WriteHeader(http.StatusOK)
		return
	}

	// ユーザーを作成
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
