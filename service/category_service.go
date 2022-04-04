package service

import (
	"context"
	"silocorp/golang-restful-api/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryID int64)
	FindByID(ctx context.Context, categoryID int64) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
