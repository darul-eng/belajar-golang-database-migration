package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-restful-api-2/helper"
	"golang-restful-api-2/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO category(name) VALUES ($1) RETURNING id"
	var lastInsertId int
	result := tx.QueryRowContext(ctx, SQL, category.Name)

	err := result.Scan(&lastInsertId)
	helper.PanicIferror(err)

	category.Id = lastInsertId

	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "UPDATE category SET name = $1 WHERE id = $2"
	//tx.QueryRowContext(ctx, SQL, category.Name, category.Id)
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIferror(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "delete from category where id = $1"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIferror(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "select id, name from category where id = $1"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIferror(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIferror(err)

		return category, nil
	} else {
		return category, errors.New("category is not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "select id, name from category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIferror(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIferror(err)
		categories = append(categories, category)
	}

	return categories
}
