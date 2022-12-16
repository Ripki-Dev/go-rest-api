package repository

import (
	"context"
	"database/sql"
	"errors"
	"restGo/helper"
	"restGo/model/domain"
)

type CategoryRepositoryImpl struct{}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO categories(name) values(?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicError(err)

	id, err := result.LastInsertId()
	helper.PanicError(err)

	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "SELECT id, name FROM categories"
	result, err := tx.QueryContext(ctx, SQL)
	helper.PanicError(err)

	defer result.Close()

	var categories []domain.Category
	for result.Next() {
		category := domain.Category{}
		err := result.Scan(&category.Id, &category.Name)
		helper.PanicError(err)
		categories = append(categories, category)
	}

	return categories
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "SELECT id, name FROM categories WHERE id = ?"
	result, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicError(err)

	defer result.Close()

	category := domain.Category{}

	if result.Next() {
		err := result.Scan(&category.Id, &category.Name)
		helper.PanicError(err)

		return category, nil
	} else {
		return category, errors.New("category not found")
	}
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "UPDATE categories SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "DELETE FROM categories WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicError(err)
}
