package models

import (
	"database/sql"
	"fmt"
)

// User represents a user in the system.
type User struct {
	Username string
	Password string
	Email    string
}

var (
	// ErrDuplicateUsername is returned when trying to insert a user with an existing username.
	ErrDuplicateUsername = fmt.Errorf("duplicate username")

	// ErrDuplicateEmail is returned when trying to insert a user with an existing email.
	ErrDuplicateEmail = fmt.Errorf("duplicate email")
)

func InsertUser(db *sql.DB, username, email, password string) error {
	// Check if username already exists
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", username).Scan(&count)
	if err != nil {
		return fmt.Errorf("error checking duplicate username: %w", err)
	}
	if count > 0 {
		return ErrDuplicateUsername
	}

	// Check if email already exists
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", email).Scan(&count)
	if err != nil {
		return fmt.Errorf("error checking duplicate email: %w", err)
	}
	if count > 0 {
		return ErrDuplicateEmail
	}

	// Insert the user into the database
	stmt, err := db.Prepare(`INSERT INTO users (username, email, password) VALUES (?, ?, ?)`)
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, email, password)
	if err != nil {
		return fmt.Errorf("error executing statement: %w", err)
	}

	return nil
}

// ScanUser retrieves a user from the database by ID.
func ScanUser(db *sql.DB, id int) ([]User, error) {
	query := fmt.Sprintf("SELECT username,password,email FROM users WHERE id=%d", id)
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Username, &user.Password, &user.Email); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return users, nil
}
