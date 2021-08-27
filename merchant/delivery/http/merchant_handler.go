package http

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/hezbymuhammad/payment-gateway/domain"
)

type ResponseError struct {
	Message string `json:"message"`
}

type MerchantHandler struct {
        Usecase domain.MerchantUsecase
}

func NewMerchantHandler(e *echo.Echo, u domain.MerchantUsecase) *MerchantHandler {
        handler := &MerchantHandler{
                Usecase: u,
        }

        e.POST("/merchants", handler.Store)
        e.POST("/merchants/set_child", handler.SetChild)

        return handler
}

func (h *MerchantHandler) Store(c echo.Context) error {
	ctx := c.Request().Context()
        var data domain.Merchant
        c.Bind(&data)
	if data.Name == "" {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "Bad request param"})
	}
        err := h.Usecase.Store(ctx, &data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: "Failed to proceed"})
	}

        return c.NoContent(http.StatusCreated)
}

func (h *MerchantHandler) SetChild(c echo.Context) error {
	ctx := c.Request().Context()

        var data domain.MerchantGroup
        c.Bind(&data)
	if data.ParentMerchantID == 0 || data.ChildMerchantID == 0 {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: "Bad request param"})
	}

        err := h.Usecase.SetChild(ctx, &data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: "Failed to proceed"})
	}

        return c.NoContent(http.StatusCreated)
}
