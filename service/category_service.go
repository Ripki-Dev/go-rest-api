package service

import (
	"context"
	"restGo/model/web"
)

type categoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
}
