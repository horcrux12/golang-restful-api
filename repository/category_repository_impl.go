package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/horcrux12/golang-restful-api/helper"
	"github.com/horcrux12/golang-restful-api/model/domain"
)

type CategoryRepositoryImpl struct {
	Table string
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{
		Table: "category",
	}
}

func (input *CategoryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := "INSERT into " + input.Table + " (name) ?"

	result, err := tx.ExecContext(ctx, query, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.ID = id
	return category
}

func (input *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := "UPDATE " + input.Table + " SET name = ? WHERE id = ?"

	_, err := tx.ExecContext(ctx, query, category.Name, category.ID)
	helper.PanicIfError(err)

	return category
}

func (input *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	query := "DELETE FROM " + input.Table + " WHERE id = ?"

	_, err := tx.ExecContext(ctx, query, category.Name, category.ID)
	helper.PanicIfError(err)
}

func (input *CategoryRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, category domain.Category) (result domain.Category, err error) {
	query := "SELECT id, name FROM " + input.Table + "WHERE id = ?"

	rows, err := tx.QueryContext(ctx, query, category.ID)
	helper.PanicIfError(err)

	if rows.Next() {
		err = rows.Scan(&result.ID, &result.Name)
		helper.PanicIfError(err)
		return
	} else {
		err = errors.New("category is not found")
		return
	}
	return
}

func (input *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) (result []domain.Category) {
	query := "SELECT id, name FROM " + input.Table
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)

	for rows.Next() {
		var tempCategory domain.Category
		err = rows.Scan(&tempCategory.ID, &tempCategory.Name)
		helper.PanicIfError(err)
		result = append(result, tempCategory)
	}

	return
}
