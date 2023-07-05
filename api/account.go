package api

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/NopparootSuree/Learning-GO/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,currency"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		//unique_violation  foreign_key_violation
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errResponse(err))
				return
			}
		}
	}

	ctx.JSON(http.StatusOK, account)
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type listAccountParams struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listAccount(ctx *gin.Context) {
	var param listAccountParams
	if err := ctx.ShouldBindQuery(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.ListAccountParams{
		Limit:  param.PageSize,
		Offset: (param.PageID - 1) * param.PageSize,
	}

	accounts, err := server.store.ListAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}

type deleteAccountParams struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteAccount(ctx *gin.Context) {
	var param deleteAccountParams
	if err := ctx.ShouldBindUri(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, param.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	errDelete := server.store.DeleteAccount(ctx, param.ID)
	if errDelete != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(errDelete))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "deleted successfully", "data": account})
}

type updateAccountParams struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type updateAccountRequest struct {
	Balance int64 `json:"balance" binding:"required,min=0"`
}

func (server *Server) updateAccount(ctx *gin.Context) {
	var param updateAccountParams
	if err := ctx.ShouldBindUri(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	_, err := server.store.GetAccount(ctx, param.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	var req updateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	fmt.Println(req)
	arg := db.UpdateAccountParams{
		ID:      param.ID,
		Balance: req.Balance,
	}

	update, errUpdate := server.store.UpdateAccount(ctx, arg)
	if errUpdate != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(errUpdate))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "updated successfully", "data": update})
}
