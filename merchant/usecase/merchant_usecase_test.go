package usecase_test

import (
	"context"
        "errors"
        "testing"

        "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/hezbymuhammad/payment-gateway/domain"
	"github.com/hezbymuhammad/payment-gateway/domain/mocks"
	merchantUsecase "github.com/hezbymuhammad/payment-gateway/merchant/usecase"
)

func TestStore(t *testing.T) {
        mockRepo := new(mocks.MerchantRepository)
        data := domain.Merchant{Name: "lorem"}

        mockRepo.On("Store", mock.Anything, mock.Anything).Return(nil).Once()
        mockRepo.On("InitSetting", mock.Anything, mock.Anything).Return(nil).Once()
        u := merchantUsecase.NewMerchantUsecase(mockRepo)

        err := u.Store(context.TODO(), &data)

        assert.NoError(t, err)
}


func TestFailedStore(t *testing.T) {
        mockRepo := new(mocks.MerchantRepository)
        data := domain.Merchant{Name: "lorem"}
        dummyErr := errors.New("some err")

        mockRepo.On("Store", mock.Anything, mock.Anything).Return(dummyErr).Once()
        mockRepo.On("InitSetting", mock.Anything, mock.Anything).Return(nil).Once()
        u := merchantUsecase.NewMerchantUsecase(mockRepo)

        err := u.Store(context.TODO(), &data)
        assert.Equal(t, err, dummyErr)
}

func TestSetChild(t *testing.T) {
        mockRepo := new(mocks.MerchantRepository)
        mockRepo.On("SetChild", mock.Anything, mock.Anything).Return(nil).Once()

        u := merchantUsecase.NewMerchantUsecase(mockRepo)
        data := &domain.MerchantGroup{
                ParentMerchantID: 1,
                ChildMerchantID: 2,
        }

        err := u.SetChild(context.TODO(), data)
        assert.NoError(t, err)
}

func TestFailedSetChild(t *testing.T) {
        mockRepo := new(mocks.MerchantRepository)
        dummyErr := errors.New("some err")
        mockRepo.On("SetChild", mock.Anything, mock.Anything).Return(dummyErr).Once()

        u := merchantUsecase.NewMerchantUsecase(mockRepo)
        data := &domain.MerchantGroup{
                ParentMerchantID: 1,
                ChildMerchantID: 2,
        }

        err := u.SetChild(context.TODO(), data)
        assert.Equal(t, err, dummyErr)
}
