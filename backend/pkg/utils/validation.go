package utils

import (
	"regexp"
	"strings"
)

// IsValidEmail validates email format
func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// IsValidName validates name format (2-100 characters, no special characters)
func IsValidName(name string) bool {
	if len(name) < 2 || len(name) > 100 {
		return false
	}
	
	// Check for only letters, spaces, and common punctuation
	nameRegex := regexp.MustCompile(`^[a-zA-Z\s\-'\.]+$`)
	return nameRegex.MatchString(name)
}

// SanitizeString removes leading/trailing whitespace and normalizes spaces
func SanitizeString(s string) string {
	return strings.TrimSpace(strings.Join(strings.Fields(s), " "))
}

// ValidateUserInput validates user input for create/update operations
func ValidateUserInput(name, email string) []string {
	var errors []string
	
	if name == "" {
		errors = append(errors, "Name is required")
	} else if !IsValidName(name) {
		errors = append(errors, "Name must be 2-100 characters and contain only letters, spaces, hyphens, apostrophes, and periods")
	}
	
	if email == "" {
		errors = append(errors, "Email is required")
	} else if !IsValidEmail(email) {
		errors = append(errors, "Email format is invalid")
	}
	
	return errors
}
