package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/NopparootSuree/Learning-GO/utils"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	arg := CreateTransferParams{
		FromAccountID: utils.RandomInt(6, 10),
		ToAccountID:   utils.RandomInt(6, 10),
		Amount:        utils.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}
func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)

}

func TestGetTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestDeleteTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)
	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)

	checkDeleted, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, checkDeleted)
}

func TestUpdateTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)

	arg := UpdateTransferParams{
		ID:     transfer1.ID,
		Amount: utils.RandomMoney(),
	}

	updateTransfer, err := testQueries.UpdateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updateTransfer)

	require.Equal(t, transfer1.ID, updateTransfer.ID)
	require.Equal(t, transfer1.FromAccountID, updateTransfer.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, updateTransfer.ToAccountID)
	require.Equal(t, arg.Amount, updateTransfer.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, updateTransfer.CreatedAt, time.Second)
}

func TestListTransfer(t *testing.T) {
	//for loop createTransfers
	for i := 0; i < 5; i++ {
		createRandomTransfer(t)
	}

	arg := ListTransferParams{
		Limit:  5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
