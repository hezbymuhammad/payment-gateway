package domain

import (
	"context"
)

type Transaction struct {
	ID                int64    `json:"id"`
	MerchantID        int64    `json:"merchantId"`
	ParentMerchantID  int64    `json:"parentMerchantId"`
	SettingID         int64    `json:"settingId"`
        Status            bool     `json:"status"`
}

type TransactionUsecase interface {
	GetByID(ctx context.Context, id int64) (Transaction, error)
        Store(ctx context.Context, t *Transaction) error
        Update(ctx context.Context, t *Transaction) error
}

type TransactionRepository interface {
	GetByID(ctx context.Context, id int64) (Transaction, error)
        Store(ctx context.Context, t *Transaction) error
        Update(ctx context.Context, t *Transaction) error
}
