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

func ReservationHandler(repo *repository.ReservationRepository, w http.ResponseWriter, r *http.Request) {
	middleware.EnableCORS(&w)
	if r.Method == "OPTIONS" {
		return
	}

	switch r.Method {
	case "POST":
		createReservation(repo, w, r)
	case "GET":
		readReservation(repo, w, r)
	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

func createReservation(repo *repository.ReservationRepository, w http.ResponseWriter, r *http.Request) {
	var reservation model.Reservation
	err := json.NewDecoder(r.Body).Decode(&reservation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	err = repo.Create(ctx, &reservation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(reservation)
}

func readReservation(repo *repository.ReservationRepository, w http.ResponseWriter, r *http.Request) {
	reservationID := r.URL.Query().Get("reservation_id")
	if reservationID == "" {
		http.Error(w, "Reservation ID is required", http.StatusBadRequest)
		return
	}

	reservation, err := repo.Read(r.Context(), reservationID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Reservation not found", http.StatusNotFound)
		} else {
			http.Error(w, "Server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reservation)
}
