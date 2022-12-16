package service

import (
	"context"
	"database/sql"
	"restGo/exception"
	"restGo/helper"
	"restGo/model/domain"
	"restGo/model/web"
	"restGo/repository"

	"github.com/go-playground/validator/v10"
)

type categoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) categoryService {
	return &categoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *categoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommiteOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Create(ctx, tx, category)
	return helper.ToCategoryResponse(category)
}

func (service *categoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommiteOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)
	return helper.ToCategoryResponses(categories)
}

func (service *categoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommiteOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToCategoryResponse(category)
}

func (service *categoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommiteOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	category.Name = request.Name

	category = service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *categoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommiteOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CategoryRepository.Delete(ctx, tx, category)
}
