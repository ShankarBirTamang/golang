package main

import (
	"database/sql"
	"fmt"
)

// STEP 1: Create a table
// This function creates a users table if it doesn't exist
func createTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		age INTEGER,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}
	fmt.Println("✅ Table 'users' created successfully!")
	return nil
}

// STEP 2: CRUD Operations

// CREATE - Insert a new user
func createUser(db *sql.DB, name, email string, age int) error {
	query := `INSERT INTO users (name, email, age) VALUES ($1, $2, $3) RETURNING id`

	var id int
	err := db.QueryRow(query, name, email, age).Scan(&id)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}
	fmt.Printf("✅ User created with ID: %d\n", id)
	return nil
}

// READ - Get all users
func getAllUsers(db *sql.DB) error {
	query := `SELECT id, name, email, age, created_at FROM users`

	rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	fmt.Println("\n📋 All Users:")
	fmt.Println("ID | Name | Email | Age | Created At")
	fmt.Println("---|------|-------|-----|------------")

	for rows.Next() {
		var id, age int
		var name, email string
		var createdAt string

		err := rows.Scan(&id, &name, &email, &age, &createdAt)
		if err != nil {
			return fmt.Errorf("failed to scan row: %w", err)
		}
		fmt.Printf("%d | %s | %s | %d | %s\n", id, name, email, age, createdAt)
	}

	return rows.Err()
}

// READ - Get a single user by ID
func getUserByID(db *sql.DB, id int) error {
	query := `SELECT id, name, email, age FROM users WHERE id = $1`

	var name, email string
	var age int
	err := db.QueryRow(query, id).Scan(&id, &name, &email, &age)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user with ID %d not found", id)
		}
		return fmt.Errorf("failed to get user: %w", err)
	}

	fmt.Printf("\n👤 User ID %d: %s (%s), Age: %d\n", id, name, email, age)
	return nil
}

// UPDATE - Update user information
func updateUser(db *sql.DB, id int, name string, age int) error {
	query := `UPDATE users SET name = $1, age = $2 WHERE id = $3`

	result, err := db.Exec(query, name, age, id)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with ID %d not found", id)
	}

	fmt.Printf("✅ User ID %d updated successfully!\n", id)
	return nil
}

// DELETE - Delete a user
func deleteUser(db *sql.DB, id int) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := db.Exec(query, id)
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

	fmt.Printf("✅ User ID %d deleted successfully!\n", id)
	return nil
}
