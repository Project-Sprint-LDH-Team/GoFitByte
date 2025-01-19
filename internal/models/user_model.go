package models

import (
	"time"
)

type PreferenceType string
type WeightUnit string
type HeightUnit string

const (
	PreferenceCardio PreferenceType = "CARDIO"
	PreferenceWeight PreferenceType = "WEIGHT"
	WeightUnitKG     WeightUnit     = "KG"
	WeightUnitLBS    WeightUnit     = "LBS"
	HeightUnitCM     HeightUnit     = "CM"
	HeightUnitINCH   HeightUnit     = "INCH"
)

type User struct {
	ID         string         `json:"id"`
	Email      string         `json:"email" binding:"required,email,unique"`
	Password   string         `json:"-"`
	Name       string         `json:"name" binding:"omitempty,min=2,max=100"`
	ImageUri   string         `json:"image_uri"`
	Preference PreferenceType `json:"preference" binding:"required,oneof=CARDIO WEIGHT"`
	WeightUnit WeightUnit     `json:"weight_unit" binding:"required,oneof=KG LBS"`
	HeightUnit HeightUnit     `json:"height_unit" binding:"required,oneof=CM INCH"`
	Weight     int16          `json:"weight" binding:"required,min=10,max=1000"`
	Height     int16          `json:"height" binding:"required,min=10,max=1000"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	Activities []Activity     `json:"-"` // Relasi one-to-many ke activities

}

type AuthRequest struct {
	Email    string `json:"email" binding:"required,email,min=1,max=255"`
	Password string `json:"password" binding:"required,min=8,max=32"`
}

type AuthResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
