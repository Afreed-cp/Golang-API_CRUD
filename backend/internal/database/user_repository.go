package database

import (
	"fmt"

	"goapi/internal/models"
)

// UserRepository handles user-related database operations
type UserRepository struct {
	db *DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *DB) *UserRepository {
	return &UserRepository{db: db}
}

// GetAll retrieves all users from the database
func (r *UserRepository) GetAll() ([]models.User, error) {
	query := `
		SELECT id, name, email, created_at, updated_at 
		FROM users 
		ORDER BY created_at DESC
	`
	
	rows, err := r.db.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating users: %w", err)
	}

	return users, nil
}

// GetByID retrieves a user by ID
func (r *UserRepository) GetByID(id int) (*models.User, error) {
	query := `
		SELECT id, name, email, created_at, updated_at 
		FROM users 
		WHERE id = $1
	`
	
	var user models.User
	err := r.db.DB.QueryRow(query, id).Scan(
		&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt,
	)
	
	if err != nil {
		if IsNoRowsError(err) {
			return nil, fmt.Errorf("user with ID %d not found", id)
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}


// Create creates a new user
func (r *UserRepository) Create(req models.CreateUserRequest) (*models.User, error) {
	query := `
		INSERT INTO users (name, email) 
		VALUES ($1, $2) 
		RETURNING id, name, email, created_at, updated_at
	`
	
	var user models.User
	err := r.db.DB.QueryRow(query, req.Name, req.Email).Scan(
		&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt,
	)
	
	if err != nil {
		if IsUniqueConstraintError(err) {
			return nil, fmt.Errorf("email already exists")
		}
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &user, nil
}


// Update updates an existing user
func (r *UserRepository) Update(id int, req models.UpdateUserRequest) (*models.User, error) {
	query := `
		UPDATE users 
		SET name = $1, email = $2 
		WHERE id = $3 
		RETURNING id, name, email, created_at, updated_at
	`
	
	var user models.User
	err := r.db.DB.QueryRow(query, req.Name, req.Email, id).Scan(
		&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt,
	)
	
	if err != nil {
		if IsNoRowsError(err) {
			return nil, fmt.Errorf("user with ID %d not found", id)
		}
		if IsUniqueConstraintError(err) {
			return nil, fmt.Errorf("email already exists")
		}
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return &user, nil
}


// Delete deletes a user by ID
func (r *UserRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	
	result, err := r.db.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with ID %d not found", id)
	}

	return nil
}

// GetByEmail retrieves a user by email
func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	query := `
		SELECT id, name, email, created_at, updated_at 
		FROM users 
		WHERE email = $1
	`
	
	var user models.User
	err := r.db.DB.QueryRow(query, email).Scan(
		&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt,
	)
	
	if err != nil {
		if IsNoRowsError(err) {
			return nil, fmt.Errorf("user with email %s not found", email)
		}
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	return &user, nil
}


