// internal/repository/todo.go
package repository

import (
	"database/sql"
	"errors"
	"log"
	"time"
)

type Todo struct {
	Id          string       `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Completed   bool         `json:"completed"`
	CreatedAt   string       `json:"createdAt"`
	CompletedAt sql.NullTime `json:"completedAt"`
}

type UpdateTodoInput struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TodoRepo struct {
	db *sql.DB
}

func NewTodoRepo(db *sql.DB) *TodoRepo {
	return &TodoRepo{db: db}
}

func (r *TodoRepo) Create(todo Todo) error {
	query := `INSERT INTO todos(id, title, description) VALUES(?, ?, ?)`
	args := []any{todo.Id, todo.Title, todo.Description}

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *TodoRepo) List() ([]Todo, error) {
	rows, err := r.db.Query("SELECT id, title, description, completed, createdAt, completedAt FROM todos ORDER BY createdAt DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var t Todo
		err := rows.Scan(&t.Id, &t.Title, &t.Description, &t.Completed, &t.CreatedAt, &t.CompletedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}

func (r *TodoRepo) Delete(id string) error {
	query := `DELETE FROM todos WHERE id = ?`
	args := []any{id}

	result, err := r.db.Exec(query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("todo is not found")
		}
		return err
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		return errors.New("delete todo is failed")
	}

	return nil
}

func (r *TodoRepo) Update(todo UpdateTodoInput) error {
	query := `SELECT id, title, description, completed, createdAt, completedAt FROM todos where id = ?`
	args := []any{todo.Id}

	var t Todo
	err := r.db.QueryRow(query, args...).Scan(&t.Id, &t.Title, &t.Description, &t.Completed, &t.CreatedAt, &t.CompletedAt)
	if err != nil {
		return err
	}

	if t.Title != todo.Title {
		t.Title = todo.Title
	}

	if t.Description != todo.Description {
		t.Description = todo.Description
	}

	if t.Completed != todo.Completed {
		t.Completed = todo.Completed
	}

	completedAt := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	updateQuery := `UPDATE todos SET title = ?, description = ?, completed = ?, completedAt = ? WHERE id = ?`
	updateArgs := []any{t.Title, t.Description, t.Completed, completedAt, t.Id}

	result, err := r.db.Exec(updateQuery, updateArgs...)
	if err != nil {
		log.Printf("error execute update: %s", err.Error())
		return err
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		log.Printf("error affected rows: %s", err.Error())
		return err
	}
	if affectedRows == 0 {
		return errors.New("updated failed")
	}

	return nil
}
