package usecase

import (
        "context"

	"github.com/hezbymuhammad/payment-gateway/domain"
)

type merchantUsecase struct {
        merchantRepo domain.MerchantRepository
}

func NewMerchantUsecase(mr domain.MerchantRepository) domain.MerchantUsecase {
        return &merchantUsecase{
                merchantRepo: mr,
        }
}

func (mu *merchantUsecase) Store(ctx context.Context, m *domain.Merchant) error {
        return mu.merchantRepo.Store(ctx, m)
}

func (mu *merchantUsecase) SetChild(ctx context.Context, mg *domain.MerchantGroup) error {
        return mu.merchantRepo.SetChild(ctx, mg)
}
