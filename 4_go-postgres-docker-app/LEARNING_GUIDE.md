# Go + PostgreSQL Learning Guide

## ✅ What You've Completed

- Database connection setup
- Docker Compose configuration
- Environment variable management

## 📚 Next Steps to Learn

### **Step 1: Creating Tables** ✅ (Implemented)

**What you learned:**

- Using `CREATE TABLE` SQL statements
- Executing DDL (Data Definition Language) from Go
- Using `db.Exec()` for non-query operations
- Understanding primary keys, data types, and constraints

**Key Concepts:**

- `SERIAL` - Auto-incrementing integer
- `PRIMARY KEY` - Unique identifier
- `UNIQUE` - Ensures no duplicate values
- `NOT NULL` - Required field
- `TIMESTAMP DEFAULT CURRENT_TIMESTAMP` - Auto-set creation time

---

### **Step 2: CRUD Operations** ✅ (Implemented)

**What you learned:**

- **CREATE**: Inserting data with `INSERT INTO ... VALUES`
- **READ**: Querying data with `SELECT`
- **UPDATE**: Modifying existing records
- **DELETE**: Removing records

**Key Functions:**

- `db.Exec()` - For INSERT, UPDATE, DELETE
- `db.Query()` - For SELECT returning multiple rows
- `db.QueryRow()` - For SELECT returning single row
- `rows.Scan()` - Reading data from query results
- `result.RowsAffected()` - Checking how many rows were modified

---

### **Step 3: Prepared Statements** (Next to Learn)

**Why it's important:**

- Prevents SQL injection attacks
- Improves performance (query is parsed once)
- Cleaner code with parameterized queries

**What to learn:**

```go
// Instead of string concatenation (BAD):
query := fmt.Sprintf("SELECT * FROM users WHERE id = %d", id)

// Use prepared statements (GOOD):
query := "SELECT * FROM users WHERE id = $1"
db.QueryRow(query, id)
```

**Practice:**

- Convert all queries to use `$1, $2, $3` placeholders
- Understand why `$1` is safer than string formatting

---

### **Step 4: Struct Mapping** (Next to Learn)

**What you'll learn:**

- Mapping database rows to Go structs
- Handling nullable fields with `sql.NullString`, `sql.NullInt64`
- Creating reusable data models

**Example:**

```go
type User struct {
    ID        int       `db:"id"`
    Name      string    `db:"name"`
    Email     string    `db:"email"`
    Age       int       `db:"age"`
    CreatedAt time.Time `db:"created_at"`
}

func scanUser(rows *sql.Rows) (*User, error) {
    var u User
    err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age, &u.CreatedAt)
    return &u, err
}
```

**Benefits:**

- Type safety
- Reusable code
- Easier to work with data

---

### **Step 5: Transactions** (Advanced)

**What you'll learn:**

- Grouping multiple operations
- Atomic operations (all succeed or all fail)
- Rollback on errors

**Example:**

```go
tx, err := db.Begin()
if err != nil {
    return err
}
defer tx.Rollback() // Will rollback if Commit() isn't called

// Multiple operations
tx.Exec("INSERT INTO users ...")
tx.Exec("UPDATE accounts ...")

err = tx.Commit() // Make changes permanent
```

**Use cases:**

- Transferring money between accounts
- Creating orders with multiple items
- Any operation that must be all-or-nothing

---

### **Step 6: Error Handling** (Best Practices)

**What to learn:**

- Checking for `sql.ErrNoRows`
- Proper error wrapping with `fmt.Errorf` and `%w`
- Logging vs returning errors
- Graceful error handling

**Example:**

```go
err := db.QueryRow(query, id).Scan(&user)
if err != nil {
    if err == sql.ErrNoRows {
        return nil, fmt.Errorf("user not found")
    }
    return nil, fmt.Errorf("database error: %w", err)
}
```

---

### **Step 7: HTTP REST API** (Optional - Full Stack)

**What you'll learn:**

- Creating HTTP handlers
- RESTful endpoints (GET, POST, PUT, DELETE)
- JSON encoding/decoding
- Request/response handling

**Example endpoints:**

- `GET /users` - List all users
- `GET /users/:id` - Get single user
- `POST /users` - Create user
- `PUT /users/:id` - Update user
- `DELETE /users/:id` - Delete user

**Technologies:**

- `net/http` - Standard library
- `gorilla/mux` - Router (optional)
- `encoding/json` - JSON handling

---

### **Step 8: Connection Pooling** (Performance)

**What you'll learn:**

- Configuring database connection pool
- Setting max open connections
- Setting max idle connections
- Connection lifetime

**Example:**

```go
db.SetMaxOpenConns(25)
db.SetMaxIdleConns(5)
db.SetConnMaxLifetime(5 * time.Minute)
```

---

### **Step 9: Migrations** (Production Ready)

**What you'll learn:**

- Version control for database schema
- Tools like `golang-migrate` or `sql-migrate`
- Up and down migrations
- Managing schema changes over time

---

### **Step 10: Testing** (Quality Assurance)

**What you'll learn:**

- Unit testing database operations
- Integration testing with test database
- Mocking database connections
- Test fixtures and setup/teardown

---

## 🎯 Recommended Learning Order

1. ✅ **Step 1 & 2** - Tables and CRUD (You're here!)
2. **Step 3** - Prepared Statements (Critical for security)
3. **Step 4** - Struct Mapping (Makes code cleaner)
4. **Step 6** - Error Handling (Best practices)
5. **Step 5** - Transactions (When you need atomicity)
6. **Step 7** - HTTP API (If building a web service)
7. **Step 8** - Connection Pooling (Performance optimization)
8. **Step 9** - Migrations (Production deployment)
9. **Step 10** - Testing (Quality assurance)

## 🚀 Quick Practice Ideas

1. **Add validation**: Check email format, age > 0, etc.
2. **Add search**: Find users by name or email
3. **Add pagination**: Limit results to 10 per page
4. **Add sorting**: Order by name, age, created_at
5. **Add relationships**: Create a `posts` table with foreign key to `users`

## 📖 Resources

- [Go database/sql tutorial](https://go.dev/doc/database/)
- [PostgreSQL Go driver (lib/pq)](https://github.com/lib/pq)
- [SQL Tutorial](https://www.w3schools.com/sql/)

---

**Happy Learning! 🎉**
