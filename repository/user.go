package repository

import (
	"context"
	"database/sql"

	"github.com/toshiki-git/lab-server-scheduler-api/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create は新しいユーザーをデータベースに追加します。
func (r *UserRepository) Create(ctx context.Context, u *model.User) error {
	query := "INSERT INTO users (name, email, image_url) VALUES ($1, $2, $3)"
	_, err := r.db.ExecContext(ctx, query, u.Name, u.Email, u.ImageURL)
	return err
}

// Read はユーザー情報を取得します。
func (r *UserRepository) Read(ctx context.Context, id string) (*model.User, error) {
	u := &model.User{}
	query := "SELECT user_id, name, email, image_url, created_at, updated_at FROM users WHERE user_id = $1"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&u.UserId, &u.Name, &u.Email, &u.ImageURL, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Update はユーザー情報を更新します。
func (r *UserRepository) Update(ctx context.Context, u *model.User) error {
	_, err := r.db.ExecContext(ctx, "UPDATE users SET name = $1, email = $2 WHERE user_id = $3", u.Name, u.Email, u.UserId)
	return err
}

// Delete はユーザーを削除します。
func (r *UserRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM users WHERE user_id = $1", id)
	return err
}
