package http

import (
        "fmt"
	"net/http"
        "strconv"
        "log"

	"github.com/labstack/echo"

	"github.com/hezbymuhammad/payment-gateway/domain"
)

type ResponseError struct {
	Message string `json:"message"`
}

type TransactionHandler struct {
        Usecase domain.TransactionUsecase
}

func NewTransactionHandler(e *echo.Echo, u domain.TransactionUsecase) *TransactionHandler {
        handler := &TransactionHandler{
                Usecase: u,
        }

        e.POST("/transactions", handler.Store)
        e.PUT("/transactions/:id", handler.Update)
        e.GET("/transactions/:id", handler.GetByID)

        return handler
}

func (h *TransactionHandler) Store(c echo.Context) error {
	ctx := c.Request().Context()
        var data domain.Transaction
        c.Bind(&data)
	if data.MerchantID == 0 || data.SettingID == 0 {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "Bad request param"})
	}
        err := h.Usecase.Store(ctx, &data)
	if err != nil && fmt.Sprint(err) == "Unauthorized" {
		return c.JSON(http.StatusUnauthorized, ResponseError{Message: "Unauthorized"})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: "Failed to proceed"})
	}

        return c.NoContent(http.StatusCreated)
}

func (h *TransactionHandler) Update(c echo.Context) error {
        idP, err := strconv.Atoi(c.Param("id"))
        if err != nil {
		return c.JSON(http.StatusNotFound, ResponseError{Message: "Not found"})
	}
        id := int64(idP)

	ctx := c.Request().Context()
        var data domain.Transaction
        c.Bind(&data)
        data.ID = id
	if data.MerchantID == 0 || data.SettingID == 0 || data.ID == 0 {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "Bad request param"})
	}
        err = h.Usecase.Update(ctx, &data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: "Failed to proceed"})
	}

        return c.NoContent(http.StatusOK)
}

func (h *TransactionHandler) GetByID(c echo.Context) error {
        idP, err := strconv.Atoi(c.Param("id"))
        log.Println(c.Param("id"))
        if err != nil {
		return c.JSON(http.StatusNotFound, ResponseError{Message: "Not found"})
	}
        id := int64(idP)

	ctx := c.Request().Context()
        res, err := h.Usecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: "Failed to proceed"})
	}

        return c.JSON(http.StatusOK, res)
}
