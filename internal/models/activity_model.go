package models

import "time"

type Activity struct {
	ID                string    `json:"id"`
	UserID            string    `json:"user_id"`
	ActivityType      string    `json:"activity_type"`
	DoneAt            time.Time `json:"done_at"`
	DurationInMinutes int       `json:"duration_in_minutes"`
	CaloriesBurned    int       `json:"calories_burned"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
