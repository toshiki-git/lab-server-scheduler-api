package repository

import (
	"context"
	"database/sql"

	"github.com/toshiki-git/lab-server-scheduler-api/model"
)

type ReservationRepository struct {
	db *sql.DB
}

func NewReservationRepository(db *sql.DB) *ReservationRepository {
	return &ReservationRepository{db: db}
}

func (r *ReservationRepository) Create(ctx context.Context, m *model.Reservation) error {
	query := "INSERT INTO reservations (user_id, title, is_all_day, start_time, end_time) VALUES ($1, $2, $3, $4, $5)"
	_, err := r.db.ExecContext(ctx, query, m.UserId, m.Title, m.IsAllDay, m.StartTime, m.EndTime)
	return err
}

func (repo *ReservationRepository) Read(ctx context.Context, id string) (*model.Reservation, error) {
	reservation := &model.Reservation{}
	query := `
		SELECT reservation_id, user_id, title, is_all_day, start_time, end_time, created_at, updated_at
		FROM reservations
		WHERE reservation_id = $1`
	err := repo.db.QueryRowContext(ctx, query, id).Scan(
		&reservation.ReservationId, &reservation.UserId, &reservation.Title, &reservation.IsAllDay,
		&reservation.StartTime, &reservation.EndTime, &reservation.CreatedAt, &reservation.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err // これにより、上位のハンドラで404 Not Foundを扱いやすくなる
		}
		return nil, err
	}
	return reservation, nil
}

/* // Update はユーザー情報を更新します。
func (r *ReservationRepository) Update(ctx context.Context, u *model.Reservation) error {
	_, err := r.db.ExecContext(ctx, "UPDATE users SET name = $1, email = $2 WHERE user_id = $3", u.Name, u.Email, u.UserId)
	return err
}

// Delete はユーザーを削除します。
func (r *ReservationRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM users WHERE user_id = $1", id)
	return err
}
*/
