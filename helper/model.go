package helper

import (
	"restGo/model/domain"
	"restGo/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var CategoryResponse []web.CategoryResponse
	for _, category := range categories {
		CategoryResponse = append(CategoryResponse, web.CategoryResponse(category))
	}
	return CategoryResponse
}
