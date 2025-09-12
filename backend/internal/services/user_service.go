package services

import (
	"fmt"

	"goapi/internal/database"
	"goapi/internal/models"
)

// UserService handles user business logic
type UserService struct {
	userRepo *database.UserRepository
}

// NewUserService creates a new user service
func NewUserService(userRepo *database.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// GetAllUsers retrieves all users
func (s *UserService) GetAllUsers() ([]models.UserResponse, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	responses := make([]models.UserResponse, len(users))
	for i, user := range users {
		responses[i] = user.ToResponse()
	}

	return responses, nil
}

// GetUserByID retrieves a user by ID
func (s *UserService) GetUserByID(id int) (*models.UserResponse, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	response := user.ToResponse()
	return &response, nil
}

// CreateUser creates a new user
func (s *UserService) CreateUser(req models.CreateUserRequest) (*models.UserResponse, error) {
	// Validate email uniqueness
	_, err := s.userRepo.GetByEmail(req.Email)
	if err == nil {
		return nil, fmt.Errorf("email already exists")
	}

	user, err := s.userRepo.Create(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	response := user.ToResponse()
	return &response, nil
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(id int, req models.UpdateUserRequest) (*models.UserResponse, error) {
	// Check if user exists
	_, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Check if email is being changed and if new email already exists
	existingUser, err := s.userRepo.GetByEmail(req.Email)
	if err == nil && existingUser.ID != id {
		return nil, fmt.Errorf("email already exists")
	}

	user, err := s.userRepo.Update(id, req)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	response := user.ToResponse()
	return &response, nil
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(id int) error {
	err := s.userRepo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}
