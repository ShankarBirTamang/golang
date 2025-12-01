package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// Database connection string from environment variables
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Open database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()

	// Test connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}
	fmt.Println("✅ Connected to Postgres successfully!")

	// STEP 1: Create table
	if err := createTable(db); err != nil {
		log.Fatal(err)
	}

	// STEP 2: CRUD Operations Demo
	fmt.Println("\n" + "==================================================")
	fmt.Println("CRUD OPERATIONS DEMO")
	fmt.Println("==================================================")

	// CREATE - Add some users
	fmt.Println("\n1️⃣ CREATE Operations:")
	if err := createUser(db, "Alice", "alice@example.com", 25); err != nil {
		log.Println("Error:", err)
	}
	if err := createUser(db, "Bob", "bob@example.com", 30); err != nil {
		log.Println("Error:", err)
	}
	if err := createUser(db, "Charlie", "charlie@example.com", 28); err != nil {
		log.Println("Error:", err)
	}

	// READ - Get all users
	fmt.Println("\n2️⃣ READ Operations:")
	if err := getAllUsers(db); err != nil {
		log.Println("Error:", err)
	}

	// READ - Get single user
	if err := getUserByID(db, 1); err != nil {
		log.Println("Error:", err)
	}

	// UPDATE - Update a user
	fmt.Println("\n3️⃣ UPDATE Operations:")
	if err := updateUser(db, 1, "Alice Updated", 26); err != nil {
		log.Println("Error:", err)
	}

	// Verify update
	if err := getUserByID(db, 1); err != nil {
		log.Println("Error:", err)
	}

	// DELETE - Delete a user
	fmt.Println("\n4️⃣ DELETE Operations:")
	if err := deleteUser(db, 3); err != nil {
		log.Println("Error:", err)
	}

	// Show final state
	fmt.Println("\n📊 Final State:")
	if err := getAllUsers(db); err != nil {
		log.Println("Error:", err)
	}
}
