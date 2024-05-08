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
	userrRepo := repository.NewUserRepository(db)
	reservationRepo := repository.NewReservationRepository(db)

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		handler.UserHandler(userrRepo, w, r)
	})
	http.HandleFunc("/reservations", func(w http.ResponseWriter, r *http.Request) {
		handler.ReservationHandler(reservationRepo, w, r)
	})

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
