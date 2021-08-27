package http_test

import (
	"encoding/json"
        "errors"
        "testing"
	"net/http"
	"net/http/httptest"
        "strings"

        "github.com/labstack/echo"
        "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/hezbymuhammad/payment-gateway/domain"
	"github.com/hezbymuhammad/payment-gateway/domain/mocks"
	merchantHttp "github.com/hezbymuhammad/payment-gateway/merchant/delivery/http"
)

func TestStore(t *testing.T) {
        mockUsecase := new(mocks.MerchantUsecase)
        mockUsecase.On("Store", mock.Anything, mock.Anything).Return(nil).Once()

        data := &domain.Merchant{
                Name: "lorem",
        }
        j, err := json.Marshal(data)
        assert.NoError(t, err)

	e := echo.New()

	req, err := http.NewRequest(echo.POST, "/merchants", strings.NewReader(string(j)))
        assert.NoError(t, err)

        req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
        ctx.SetPath("/merchants")

        handler := merchantHttp.NewMerchantHandler(echo.New(), mockUsecase)
        err = handler.Store(ctx)

        assert.NoError(t, err)
        assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestFailedStore(t *testing.T) {
        dummyErr := errors.New("dummy err")
        mockUsecase := new(mocks.MerchantUsecase)
        mockUsecase.On("Store", mock.Anything, mock.Anything).Return(dummyErr).Once()

        data := &domain.Merchant{
                Name: "lorem",
        }
        j, err := json.Marshal(data)
        assert.NoError(t, err)

	e := echo.New()

	req, err := http.NewRequest(echo.POST, "/merchants", strings.NewReader(string(j)))
        assert.NoError(t, err)

        req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
        ctx.SetPath("/merchants")

        handler := merchantHttp.NewMerchantHandler(echo.New(), mockUsecase)
        err = handler.Store(ctx)

        assert.NoError(t, err)
        assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestSetChild(t *testing.T) {
        mockUsecase := new(mocks.MerchantUsecase)
        mockUsecase.On("SetChild", mock.Anything, mock.Anything).Return(nil).Once()

        data := &domain.MerchantGroup{
                ParentMerchantID: 1,
                ChildMerchantID: 2,
        }
        j, err := json.Marshal(data)
        assert.NoError(t, err)

	e := echo.New()

	req, err := http.NewRequest(echo.POST, "/merchants/set_child", strings.NewReader(string(j)))
        assert.NoError(t, err)

        req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
        ctx.SetPath("/merchants/set_child")

        handler := merchantHttp.NewMerchantHandler(echo.New(), mockUsecase)
        err = handler.SetChild(ctx)

        assert.NoError(t, err)
        assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestFailedSetChild(t *testing.T) {
        dummyErr := errors.New("dummy err")
        mockUsecase := new(mocks.MerchantUsecase)
        mockUsecase.On("SetChild", mock.Anything, mock.Anything).Return(dummyErr).Once()

        data := &domain.MerchantGroup{
                ParentMerchantID: 1,
                ChildMerchantID: 2,
        }
        j, err := json.Marshal(data)
        assert.NoError(t, err)

	e := echo.New()

	req, err := http.NewRequest(echo.POST, "/merchants/set_child", strings.NewReader(string(j)))
        assert.NoError(t, err)

        req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
        ctx.SetPath("/merchants/set_child")

        handler := merchantHttp.NewMerchantHandler(echo.New(), mockUsecase)
        err = handler.SetChild(ctx)

        assert.NoError(t, err)
        assert.Equal(t, http.StatusInternalServerError, rec.Code)
}
