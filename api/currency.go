package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julienxr/simplebank/utils"
)

type registerCurrencyRequest struct {
	Currency string `uri:"currency" binding:"required"`
}

func (server *Server) registerCurrency(ctx *gin.Context) {
	var req registerCurrencyRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if len(req.Currency) > 3 {
		ctx.JSON(http.StatusBadRequest, "invalid input")
		return
	}

	utils.AddSupportForCurrency(req.Currency)

	ctx.JSON(http.StatusAccepted, nil)
}
