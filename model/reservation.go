package model

type Reservation struct {
	ReservationId string `json:"reservation_id"`
	UserId        string `json:"user_id"`
	Title         string `json:"title"`
	IsAllDay      bool   `json:"is_all_day"`
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
