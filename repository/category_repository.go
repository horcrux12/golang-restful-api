package repository

import (
	"context"
	"database/sql"
	"silocorp/golang-restful-api/model/domain"
)

type CategoryRepository interface {
	Create(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindByID(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
