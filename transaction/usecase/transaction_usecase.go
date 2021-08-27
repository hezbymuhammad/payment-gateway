package usecase

import (
        "context"
        "errors"

	"github.com/hezbymuhammad/payment-gateway/domain"
)

type transactionUsecase struct {
        merchantRepo domain.MerchantRepository
        transactionRepo domain.TransactionUsecase
}

func NewTransactionUsecase(mr domain.MerchantRepository, tr domain.TransactionRepository) domain.TransactionUsecase {
        return &transactionUsecase{
                merchantRepo: mr,
                transactionRepo: tr,
        }
}

func (tu *transactionUsecase) GetByID(ctx context.Context, id int64) (domain.Transaction, error) {
        return tu.transactionRepo.GetByID(ctx, id)
}
func (tu *transactionUsecase) Store(ctx context.Context, t *domain.Transaction) error {
        if t.MerchantID != t.ParentMerchantID {
                return tu.storeForChild(ctx, t)
        } else {
                return tu.store(ctx, t)
        }
}
func (tu *transactionUsecase) Update(ctx context.Context, t *domain.Transaction) error {
        return tu.transactionRepo.Update(ctx, t)
}

func (tu *transactionUsecase) store(ctx context.Context, t *domain.Transaction) error {
        return tu.transactionRepo.Store(ctx, t)
}
func (tu *transactionUsecase) storeForChild(ctx context.Context, t *domain.Transaction) error {
        authorized, err := tu.merchantRepo.IsAuthorizedParent(
                ctx,
                &domain.MerchantGroup{
                        ParentMerchantID: t.ParentMerchantID,
                        ChildMerchantID: t.MerchantID,
                },
        )
        if err != nil {
                return err
        }
        if authorized == false {
                return errors.New("Unauthorized")
        }

        return tu.transactionRepo.Store(ctx, t)
}
