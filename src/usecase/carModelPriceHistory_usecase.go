package usecase

import (
	"CarBuyingAndSelling/src/config"
	"CarBuyingAndSelling/src/domain/filter"
	"CarBuyingAndSelling/src/domain/model"
	"CarBuyingAndSelling/src/domain/repository"
	"CarBuyingAndSelling/src/usecase/dto"
	"context"
)

type CarModelPriceHistoryUsecase struct {
	base *BaseUsecase[model.CarModelPriceHistory, dto.CreateCarModelPriceHistory, dto.UpdateCarModelPriceHistory, dto.CarModelPriceHistory]
}

func NewCarModelPriceHistoryUsecase(cfg *config.Config, repository repository.CarModelPriceHistoryRepository) *CarModelPriceHistoryUsecase {
	return &CarModelPriceHistoryUsecase{
		base: NewBaseUsecase[model.CarModelPriceHistory, dto.CreateCarModelPriceHistory, dto.UpdateCarModelPriceHistory, dto.CarModelPriceHistory](cfg, repository),
	}
}

// Create
func (u *CarModelPriceHistoryUsecase) Create(ctx context.Context, req dto.CreateCarModelPriceHistory) (dto.CarModelPriceHistory, error) {
	return u.base.Create(ctx, req)
}

// Update
func (s *CarModelPriceHistoryUsecase) Update(ctx context.Context, id int, req dto.UpdateCarModelPriceHistory) (dto.CarModelPriceHistory, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarModelPriceHistoryUsecase) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CarModelPriceHistoryUsecase) GetById(ctx context.Context, id int) (dto.CarModelPriceHistory, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CarModelPriceHistoryUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.CarModelPriceHistory], error) {
	return s.base.GetByFilter(ctx, req)
}
