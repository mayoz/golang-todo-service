package internal

import (
	"database/sql"
	"errors"
)

type repo struct {
	db *sql.DB
}

// NewRepo creates a new repository instance
func NewRepository(db *sql.DB) *repo {
	return &repo{db}
}

// Get all of the items from the storage
func (r *repo) Get() ([]*Todo, error) {
	todos := make([]*Todo, 0)

	rows, err := r.db.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		todo := new(Todo)

		err = rows.Scan(&todo.ID, &todo.Text, &todo.Completed, &todo.CreatedAt)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

// Store a new item into the storage
func (r *repo) Store(text string) (int64, error) {
	stmt, err := r.db.Prepare("INSERT INTO todos (text) VALUES (?)")
	if err != nil {
		return -1, err
	}

	res, err := stmt.Exec(text)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

// Store a new item into the storage
func (r *repo) Find(id int64) (Todo, error) {
	var todo Todo

	row := r.db.QueryRow("SELECT * FROM todos WHERE id = ?", id)

	err := row.Scan(&todo.ID, &todo.Text, &todo.Completed, &todo.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return todo, errors.New("resource not found")
		}

		return todo, err
	}

	return todo, nil
}

// Complete mark an item in the storage
func (r *repo) Complete(id int64) error {
	_, err := r.db.Exec("UPDATE todos SET completed = 1 WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

// Uncomplete mark an item in the storage
func (r *repo) Uncomplete(id int64) error {
	_, err := r.db.Exec("UPDATE todos SET completed = 0 WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

// Destroy an item from the storage
func (r *repo) Destroy(id int64) error {
	_, err := r.db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
