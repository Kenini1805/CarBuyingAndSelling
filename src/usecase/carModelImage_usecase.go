package usecase

import (
	"CarBuyingAndSelling/src/config"
	"CarBuyingAndSelling/src/domain/filter"
	"CarBuyingAndSelling/src/domain/model"
	"CarBuyingAndSelling/src/domain/repository"
	"CarBuyingAndSelling/src/usecase/dto"
	"context"
)

type CarModelImageUsecase struct {
	base *BaseUsecase[model.CarModelImage, dto.CreateCarModelImage, dto.UpdateCarModelImage, dto.CarModelImage]
}

func NewCarModelImageUsecase(cfg *config.Config, repository repository.CarModelImageRepository) *CarModelImageUsecase {
	return &CarModelImageUsecase{
		base: NewBaseUsecase[model.CarModelImage, dto.CreateCarModelImage, dto.UpdateCarModelImage, dto.CarModelImage](cfg, repository),
	}
}

// Create
func (u *CarModelImageUsecase) Create(ctx context.Context, req dto.CreateCarModelImage) (dto.CarModelImage, error) {
	return u.base.Create(ctx, req)
}

// Update
func (s *CarModelImageUsecase) Update(ctx context.Context, id int, req dto.UpdateCarModelImage) (dto.CarModelImage, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarModelImageUsecase) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CarModelImageUsecase) GetById(ctx context.Context, id int) (dto.CarModelImage, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CarModelImageUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.CarModelImage], error) {
	return s.base.GetByFilter(ctx, req)
}
