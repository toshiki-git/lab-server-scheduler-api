package model

type User struct {
	UserId    string `json:"user_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	ImageURL  string `json:"image_url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
