package usecase_test

import (
	"context"
        "testing"

        "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/hezbymuhammad/payment-gateway/domain"
	"github.com/hezbymuhammad/payment-gateway/domain/mocks"
	transactionUsecase "github.com/hezbymuhammad/payment-gateway/transaction/usecase"
)

func TestStore(t *testing.T) {
        mockMerchantRepo := new(mocks.MerchantRepository)
        mockTransactionRepo := new(mocks.TransactionRepository)
        data := domain.Transaction{
                MerchantID: 1,
                ParentMerchantID: 1,
                SettingID: 1,
                Status: false,
        }

        mockMerchantRepo.On("IsAuthorizedParent", mock.Anything, mock.Anything).Return(true, nil).Once()
        mockTransactionRepo.On("Store", mock.Anything, mock.Anything).Return(nil).Once()
        u := transactionUsecase.NewTransactionUsecase(mockMerchantRepo, mockTransactionRepo)

        err := u.Store(context.TODO(), &data)

        assert.NoError(t, err)
}

func TestStoreForChild(t *testing.T) {
        mockMerchantRepo := new(mocks.MerchantRepository)
        mockTransactionRepo := new(mocks.TransactionRepository)
        data := domain.Transaction{
                MerchantID: 1,
                ParentMerchantID: 2,
                SettingID: 1,
                Status: false,
        }

        mockMerchantRepo.On("IsAuthorizedParent", mock.Anything, mock.Anything).Return(true, nil).Once()
        mockTransactionRepo.On("Store", mock.Anything, mock.Anything).Return(nil).Once()
        u := transactionUsecase.NewTransactionUsecase(mockMerchantRepo, mockTransactionRepo)

        err := u.Store(context.TODO(), &data)

        assert.NoError(t, err)
}

func TestStoreForChildUnauthorized(t *testing.T) {
        mockMerchantRepo := new(mocks.MerchantRepository)
        mockTransactionRepo := new(mocks.TransactionRepository)
        data := domain.Transaction{
                MerchantID: 1,
                ParentMerchantID: 2,
                SettingID: 1,
                Status: false,
        }

        mockMerchantRepo.On("IsAuthorizedParent", mock.Anything, mock.Anything).Return(false, nil).Once()
        mockTransactionRepo.On("Store", mock.Anything, mock.Anything).Return(nil).Once()
        u := transactionUsecase.NewTransactionUsecase(mockMerchantRepo, mockTransactionRepo)

        err := u.Store(context.TODO(), &data)

        assert.Error(t, err)
}

func TestGetByID(t *testing.T) {
        mockMerchantRepo := new(mocks.MerchantRepository)
        mockTransactionRepo := new(mocks.TransactionRepository)
        data := domain.Transaction{
                ID: 1,
                MerchantID: 1,
                ParentMerchantID: 2,
                SettingID: 1,
                Status: false,
        }
        mockTransactionRepo.On("GetByID", mock.Anything, int64(1)).Return(data, nil).Once()
        u := transactionUsecase.NewTransactionUsecase(mockMerchantRepo, mockTransactionRepo)

        res, err := u.GetByID(context.TODO(), int64(1))

        assert.NoError(t, err)
        assert.Equal(t, res, data)
}

func TestUpdate(t *testing.T) {
        mockMerchantRepo := new(mocks.MerchantRepository)
        mockTransactionRepo := new(mocks.TransactionRepository)
        data := &domain.Transaction{
                ID: 1,
                MerchantID: 1,
                ParentMerchantID: 2,
                SettingID: 1,
                Status: false,
        }
        mockTransactionRepo.On("Update", mock.Anything, mock.Anything).Return(nil).Once()
        u := transactionUsecase.NewTransactionUsecase(mockMerchantRepo, mockTransactionRepo)

        err := u.Update(context.TODO(), data)

        assert.NoError(t, err)
}
