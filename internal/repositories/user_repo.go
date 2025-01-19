package repositories

import (
	"database/sql"
	"errors"
	"fit-byte-go/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// update user data
func (r *UserRepository) UpdateUser(userID string, user *models.User) error {
	query := `UPDATE users set preference=$1, weight_unit=$2, height_unit=$3, weight=$4, height=$5, name=$6,image_uri=$7, created_at=Now(), updated_at=Now() where id=$8`

	_, err := r.db.Exec(
		query,
		user.Preference,
		user.WeightUnit,
		user.HeightUnit,
		user.Weight,
		user.Height,
		user.Name,
		user.ImageUri,
		userID,
	)
	if err != nil {
		return err
	}
	return nil
}

// get user data
func (r *UserRepository) GetUserByID(userID string) (*models.User, error) {
	var user models.User
	query := `SELECT COALESCE(preference, '') as preference, COALESCE(weight_unit, '') as weight_unit, COALESCE(height_unit, '') as height_unit, COALESCE(weight, 0) as weight, COALESCE(height, 0) as height, email, COALESCE(name, '') as name, COALESCE(image_uri, '') as image_uri FROM users WHERE id=$1`

	err := r.db.QueryRow(query, userID).Scan(
		&user.Preference,
		&user.WeightUnit,
		&user.HeightUnit,
		&user.Weight,
		&user.Height,
		&user.Email,
		&user.Name,
		&user.ImageUri,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &models.User{}, nil
		}
		return &models.User{}, err
	}
	return &user, nil
}
