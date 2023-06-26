package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/NopparootSuree/Learning-GO/utils"
	"github.com/stretchr/testify/require"
)

func createRandomEntrie(t *testing.T) Entry {
	arg := CreateEntrieParams{
		AccountID: utils.RandomInt(6, 10),
		Amount:    utils.RandomMoney(),
	}

	entrie, err := testQueries.CreateEntrie(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entrie)

	require.Equal(t, arg.AccountID, entrie.AccountID)
	require.Equal(t, arg.Amount, entrie.Amount)

	require.NotZero(t, entrie.ID)
	require.NotZero(t, entrie.CreatedAt)

	return entrie
}

func TestCreateEntrie(t *testing.T) {
	createRandomEntrie(t)
}

func TestGetEntrie(t *testing.T) {
	entrie := createRandomEntrie(t)
	getEntrie, err := testQueries.GetEntrie(context.Background(), entrie.ID)
	require.NoError(t, err)
	require.NotEmpty(t, getEntrie)

	require.Equal(t, entrie.ID, getEntrie.ID)
	require.Equal(t, entrie.AccountID, getEntrie.AccountID)
	require.Equal(t, entrie.Amount, getEntrie.Amount)
	require.WithinDuration(t, entrie.CreatedAt, getEntrie.CreatedAt, time.Second)
}

func TestUpdateEntrie(t *testing.T) {
	entrie := createRandomEntrie(t)

	arg := UpdateEntrieParams{
		ID:     entrie.ID,
		Amount: utils.RandomMoney(),
	}

	updateEntrie, err := testQueries.UpdateEntrie(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updateEntrie)

	require.Equal(t, entrie.ID, updateEntrie.ID)
	require.Equal(t, entrie.AccountID, updateEntrie.AccountID)
	require.Equal(t, arg.Amount, updateEntrie.Amount)
	require.WithinDuration(t, entrie.CreatedAt, updateEntrie.CreatedAt, time.Second)
}

func TestDeleteEntrie(t *testing.T) {
	entrie := createRandomEntrie(t)
	err := testQueries.DeleteEntrie(context.Background(), entrie.ID)
	require.NoError(t, err)

	checkDeleted, err := testQueries.GetEntrie(context.Background(), entrie.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, checkDeleted)
}

func TestListEntrie(t *testing.T) {
	//for loop create entry
	for i := 0; i < 5; i++ {
		createRandomEntrie(t)
	}

	arg := ListEntrieParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntrie(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entrie := range entries {
		require.NotEmpty(t, entrie)
	}
}
