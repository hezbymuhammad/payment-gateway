package domain

import (
	"context"
)

type Merchant struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
}

type MerchantGroup struct {
	ParentMerchantID         int64      `json:"parentMerchantId"`
	ChildMerchantID          int64      `json:"childMerchantId"`
}

type MerchantUsecase interface {
        Store(ctx context.Context, m *Merchant) error
        SetChild(ctx context.Context, mg *MerchantGroup) error
}

type MerchantRepository interface {
        Store(ctx context.Context, m *Merchant) error
        InitSetting(ctx context.Context, m *Merchant) error
        SetChild(ctx context.Context, mg *MerchantGroup) error
        IsAuthorizedParent(ctx context.Context, mg *MerchantGroup) (bool, error)
}
