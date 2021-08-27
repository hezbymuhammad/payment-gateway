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
	transactionHttp "github.com/hezbymuhammad/payment-gateway/transaction/delivery/http"
)

func TestStore(t *testing.T) {
        mockUsecase := new(mocks.TransactionUsecase)
        mockUsecase.On("Store", mock.Anything, mock.Anything).Return(nil).Once()

        data := &domain.Transaction{
                ID: 1,
                MerchantID: 1,
                ParentMerchantID: 1,
                SettingID: 1,
                Status: true,
        }
        j, err := json.Marshal(data)
        assert.NoError(t, err)

	e := echo.New()

	req, err := http.NewRequest(echo.POST, "/transactions", strings.NewReader(string(j)))
        assert.NoError(t, err)

        req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
        ctx.SetPath("/transactions")

        handler := transactionHttp.NewTransactionHandler(echo.New(), mockUsecase)
        err = handler.Store(ctx)

        assert.NoError(t, err)
        assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestStoreInvalidParams(t *testing.T) {
        mockUsecase := new(mocks.TransactionUsecase)
        mockUsecase.On("Store", mock.Anything, mock.Anything).Return(nil).Once()

        data := &domain.Transaction{
                ID: 1,
                ParentMerchantID: 1,
                Status: true,
        }
        j, err := json.Marshal(data)
        assert.NoError(t, err)

	e := echo.New()

	req, err := http.NewRequest(echo.POST, "/transactions", strings.NewReader(string(j)))
        assert.NoError(t, err)

        req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
        ctx.SetPath("/transactions")

        handler := transactionHttp.NewTransactionHandler(echo.New(), mockUsecase)
        err = handler.Store(ctx)

        assert.NoError(t, err)
        assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestStoreUnauthorized(t *testing.T) {
        mockUsecase := new(mocks.TransactionUsecase)
        mockUsecase.On("Store", mock.Anything, mock.Anything).Return(errors.New("Unauthorized")).Once()

        data := &domain.Transaction{
                ID: 1,
                MerchantID: 1,
                ParentMerchantID: 1,
                SettingID: 1,
                Status: true,
        }
        j, err := json.Marshal(data)
        assert.NoError(t, err)

	e := echo.New()

	req, err := http.NewRequest(echo.POST, "/transactions", strings.NewReader(string(j)))
        assert.NoError(t, err)

        req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
        ctx.SetPath("/transactions")

        handler := transactionHttp.NewTransactionHandler(echo.New(), mockUsecase)
        err = handler.Store(ctx)

        assert.NoError(t, err)
        assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestFailedStore(t *testing.T) {
        dummyErr := errors.New("dummy err")
        mockUsecase := new(mocks.TransactionUsecase)
        mockUsecase.On("Store", mock.Anything, mock.Anything).Return(dummyErr).Once()

        data := &domain.Transaction{
                ID: 1,
                MerchantID: 1,
                ParentMerchantID: 1,
                SettingID: 1,
                Status: true,
        }
        j, err := json.Marshal(data)
        assert.NoError(t, err)

	e := echo.New()

	req, err := http.NewRequest(echo.POST, "/transactions", strings.NewReader(string(j)))
        assert.NoError(t, err)

        req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
        ctx.SetPath("/transactions")

        handler := transactionHttp.NewTransactionHandler(echo.New(), mockUsecase)
        err = handler.Store(ctx)

        assert.NoError(t, err)
        assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestUpdate(t *testing.T) {
        mockUsecase := new(mocks.TransactionUsecase)
        mockUsecase.On("Update", mock.Anything, mock.Anything).Return(nil).Once()

        data := &domain.Transaction{
                MerchantID: 1,
                ParentMerchantID: 1,
                SettingID: 1,
                Status: true,
        }
        j, err := json.Marshal(data)
        assert.NoError(t, err)

	e := echo.New()

	req, err := http.NewRequest(echo.PUT, "/transactions/1", strings.NewReader(string(j)))
        assert.NoError(t, err)

        req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
        ctx.SetPath("transactions/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

        handler := transactionHttp.NewTransactionHandler(echo.New(), mockUsecase)
        err = handler.Update(ctx)

        assert.NoError(t, err)
        assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateBadRequest(t *testing.T) {
        mockUsecase := new(mocks.TransactionUsecase)
        mockUsecase.On("Update", mock.Anything, mock.Anything).Return(nil).Once()

        data := &domain.Transaction{
                ParentMerchantID: 1,
                Status: true,
        }
        j, err := json.Marshal(data)
        assert.NoError(t, err)

	e := echo.New()

	req, err := http.NewRequest(echo.PUT, "/transactions/1", strings.NewReader(string(j)))
        assert.NoError(t, err)

        req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
        ctx.SetPath("transactions/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

        handler := transactionHttp.NewTransactionHandler(echo.New(), mockUsecase)
        err = handler.Update(ctx)

        assert.NoError(t, err)
        assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestUpdateMissingParams(t *testing.T) {
        mockUsecase := new(mocks.TransactionUsecase)
        mockUsecase.On("Update", mock.Anything, mock.Anything).Return(nil).Once()

        data := &domain.Transaction{
                MerchantID: 1,
                ParentMerchantID: 1,
                SettingID: 1,
                Status: true,
        }
        j, err := json.Marshal(data)
        assert.NoError(t, err)

	e := echo.New()

	req, err := http.NewRequest(echo.PUT, "/transactions/", strings.NewReader(string(j)))
        assert.NoError(t, err)

        req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
        ctx.SetPath("transactions/")

        handler := transactionHttp.NewTransactionHandler(echo.New(), mockUsecase)
        err = handler.Update(ctx)

        assert.NoError(t, err)
        assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestGetByID(t *testing.T) {
        data := domain.Transaction{
                ID: 1,
                MerchantID: 1,
                ParentMerchantID: 1,
                SettingID: 1,
                Status: true,
        }
	json_data, err := json.Marshal(data)
        assert.NoError(t, err)

        mockUsecase := new(mocks.TransactionUsecase)
        mockUsecase.On("GetByID", mock.Anything, int64(data.ID)).Return(data, nil).Once()

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/transactions/1", strings.NewReader(""))
        assert.NoError(t, err)

        req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
        ctx.SetPath("transactions/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

        handler := transactionHttp.NewTransactionHandler(echo.New(), mockUsecase)
        err = handler.GetByID(ctx)

        assert.NoError(t, err)
        assert.Equal(t, http.StatusOK, rec.Code)
        assert.Equal(t, string(json_data)+"\n", rec.Body.String())
}

func TestGetByIDInvalidParam(t *testing.T) {
        data := domain.Transaction{
                ID: 1,
                MerchantID: 1,
                ParentMerchantID: 1,
                SettingID: 1,
                Status: true,
        }

        mockUsecase := new(mocks.TransactionUsecase)
        mockUsecase.On("GetByID", mock.Anything, int64(data.ID)).Return(data, nil).Once()

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/transactions/aaa", strings.NewReader(""))
        assert.NoError(t, err)

        req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
        ctx.SetPath("transactions/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("aaa")

        handler := transactionHttp.NewTransactionHandler(echo.New(), mockUsecase)
        err = handler.GetByID(ctx)

        assert.NoError(t, err)
        assert.Equal(t, http.StatusNotFound, rec.Code)
}
