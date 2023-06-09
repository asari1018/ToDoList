package db

// schema.go provides data models in DB
import (
	"time"
)

// Task corresponds to a row in `tasks` table
type Task struct {
	ID        uint64    `db:"id"`
	Name      string    `db:"name"`
	Title     string    `db:"title"`
	CreatedAt time.Time `db:"created_at"`
	IsDone    bool      `db:"is_done"`
}

// Task corresponds to a row in `tasks` table
type User struct {
	ID        uint64    `db:"id"`
	Name     string    `db:"name"`
	Password string	   `db:"password"`
}
