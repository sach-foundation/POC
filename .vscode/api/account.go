package api

import (
	db "GOLANG/db/sqlc"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateAccountRequest struct {
	Owner string `json:"owner" binding:"required"`

	Currency string `json:"currency" binding:"required,oneof=INR USD CAD EUR"`
}

func (server *Server) CreateAccount(ctx *gin.Context) {

	var req CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		// ctx.JSON(http.StatusBadRequest, err)
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Balance:  0,
		Currency: req.Currency,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		// ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, account)
}

type GetAccountByIdReq struct {
	ID int64 `uri:"id" binding:"required,min=1"` //here min means , minimum value  should be 1
}

func (server *Server) GetAccountById(ctx *gin.Context) {

	var req GetAccountByIdReq
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, req.ID)

	if err != nil {
		if err == sql.ErrNoRows {

			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type ListAccountReq struct {
	PageId   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) ListAccount(ctx *gin.Context) {
	var req ListAccountReq

	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageId - 1) * req.PageSize,
	}

	accounts, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))

	}

	ctx.JSON(http.StatusOK, accounts)

}

type updateAccountReq struct {
	ID      int64 `json:"id" binding:"required"`
	Balance int64 `json:"balance" binding:"required"`
}

func (server *Server) UpdateAccount(ctx *gin.Context) {

	var req updateAccountReq

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.UpdateAccountParams{
		ID:      req.ID,
		Balance: req.Balance,
	}
	account, err := server.store.UpdateAccount(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)

}
