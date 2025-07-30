package api

import (
	"net/http"
	db "simple-bank/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createAccountParams struct {
	Owner       string `json:"owner" binding:"required"`
	Concurrency string `json:"concurrency" binding:"required, oneof=USD EUR THB"`
}

func (s *Server) createAccount(ctx *gin.Context) {
	var req createAccountParams
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Balance:  0,
		Currency: req.Concurrency,
	}

	account, err := s.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, account)
}
