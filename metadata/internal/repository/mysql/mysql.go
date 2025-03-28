package mysql

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"movieexample.com/metadata/internal/repository"
	model "movieexample.com/metadata/pkg/model"
)

//Repository defines a MY-SQL based movie metadata repository

type Repositoiry struct {
	db *sql.DB
}

// New creates a neew MY-SQL based repository
func New() (*Repositoiry, error) {
	db, err := sql.Open("mysql", "root:password@/movieexample")
	if err != nil {
		return nil, err
	}
	return &Repositoiry{
		db: db,
	}, nil
}

func (r *Repositoiry) Get(ctx context.Context, id string) (*model.MetaData, error) {
	var title, description, director string
	row := r.db.QueryRowContext(ctx, "SELECT title, description, director FROM movies WHERE id = ?", id)
	if err := row.Scan(&title, &description, &director); err != nil {
		if err == sql.ErrNoRows {
			return nil, repository.ErrorNotFound
		}
		return nil, err
	}
	return &model.MetaData{
		ID:          id,
		Title:       title,
		Description: description,
		Director:    director,
	}, nil
}

// Put adds movie metadat for a guven movie id
func (r *Repositoiry) Put(ctx context.Context, id string, metadata *model.MetaData) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO movies (id, title, description, director) VALUES (?, ?, ?, ?)", id, metadata.Title, metadata.Description, metadata.Director)
	if err != nil {
		return err
	}
	return nil
}
