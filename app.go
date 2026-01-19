package main

import (
	"context"
	"database/sql"
	"errors"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx  context.Context
	repo *repository
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
	runtime.LogInfo(ctx, "DB initialize")
	initDB(ctx)
	runtime.LogInfo(ctx, "DB connect")
	a.repo = NewRepo(ctx)
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

type Todo struct {
	Id          string
	Title       string
	Description string
	Completed   bool   `json:"Completed,omitempty"`
	CreatedAt   string `json:"CreatedAt,omitempty"`
	CompletedAt string `json:"CompletedAt,omitempty"`
}

type repository struct {
	db  *sql.DB
	ctx context.Context
}

func NewRepo(ctx context.Context) *repository {
	return &repository{db: DB, ctx: ctx}
}

func (a *App) AddTodo(todo Todo) error {

	query := `INSERT INTO todos (id, title, description) VALUES (?, ?, ?)`
	_, err := a.repo.db.Exec(query, todo.Id, todo.Title, todo.Description)
	if err != nil {
		runtime.LogInfof(a.ctx, "error: %s", err.Error())
		return err
	}

	runtime.LogInfof(a.ctx, "todo: %+v", todo)

	return nil
}

func (a *App) ListTodo() []Todo {
	query := `SELECT id, title, description, createdAt, completedAt FROM todos`
	rows, err := a.repo.db.Query(query)
	if err != nil {
		return []Todo{}
	}
	defer rows.Close()
	var todos []Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.CompletedAt)
		if err != nil {
			return []Todo{}
		}
		todos = append(todos, todo)
	}
	runtime.LogInfof(a.ctx, "list todos: %+v", todos)
	return todos
}

func (a *App) DeleteTodo(id string) error {
	query := `DELETE FROM todos WHERE id = ?`
	result, err := a.repo.db.Exec(query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("todo not found")
		}
		return err
	}
	affectedRow, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRow == 0 {
		return errors.New("todo ayammmm")
	}
	return nil
}
