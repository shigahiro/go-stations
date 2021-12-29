package service

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/TechBowl-japan/go-stations/model"
	_ "github.com/mattn/go-sqlite3"
)

// A TODOService implements CRUD of TODO entities.
type TODOService struct {
	db *sql.DB
}

// NewTODOService returns new TODOService.
func NewTODOService(db *sql.DB) *TODOService {
	return &TODOService{
		db: db,
	}
}

// CreateTODO creates a TODO on DB.
func (s *TODOService) CreateTODO(ctx context.Context, subject, description string) (*model.TODO, error) {
	const (
		insert  = `INSERT INTO todos(subject, description) VALUES(?, ?)`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)
	todo := &model.TODO{}

	db, err := sql.Open("sqlite3", ".sqlite3/todo.db")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db.Close()

	result, err := db.ExecContext(ctx, insert, subject, description)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	todo.ID = int(id)
	err = db.QueryRowContext(ctx, confirm, id).Scan(&todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return todo, nil
}

// ReadTODO reads TODOs on DB.
func (s *TODOService) ReadTODO(ctx context.Context, prevID, size int64) ([]*model.TODO, error) {
	const (
		read       = `SELECT id, subject, description, created_at, updated_at FROM todos ORDER BY id DESC LIMIT ?`
		readWithID = `SELECT id, subject, description, created_at, updated_at FROM todos WHERE id < ? ORDER BY id DESC LIMIT ?`
	)

	return nil, nil
}

// UpdateTODO updates the TODO on DB.
func (s *TODOService) UpdateTODO(ctx context.Context, id int64, subject, description string) (*model.TODO, error) {
	const (
		update  = `UPDATE todos SET subject = ?, description = ? WHERE id = ?`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)
	if id == 0 || subject == "" {
		return nil, errors.New("idもしくはsubjectがありません")
	}
	db := s.db

	result, err := db.ExecContext(ctx, update, subject, description, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	row, _ := result.RowsAffected()
	if row == 0 {
		return nil, &model.ErrNotFound{}
	}

	todo := model.TODO{}
	// この一行間違ってる気がする
	todo.ID = int(row)
	err = db.QueryRowContext(ctx, confirm, id).Scan(&todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &todo, nil
}

// DeleteTODO deletes TODOs on DB by ids.
func (s *TODOService) DeleteTODO(ctx context.Context, ids []int64) error {
	const deleteFmt = `DELETE FROM todos WHERE id IN (?%s)`

	return nil
}
