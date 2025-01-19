package repositories

import (
	"database/sql"
	"fit-byte-go/internal/models"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

// Register user to store on db
func (r *AuthRepository) Register(user *models.AuthRequest, userID string) error {
	query := `INSERT INTO users (id, email, password) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, userID, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := `SELECT id,email,password FROM users WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
