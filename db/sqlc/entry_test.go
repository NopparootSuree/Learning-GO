package db

import (
	"context"
	"testing"
	"time"

	"github.com/NopparootSuree/Learning-GO/utils"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    utils.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry := createRandomEntry(t, account)
	getEntry, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, getEntry)

	require.Equal(t, entry.ID, getEntry.ID)
	require.Equal(t, entry.AccountID, getEntry.AccountID)
	require.Equal(t, entry.Amount, getEntry.Amount)
	require.WithinDuration(t, entry.CreatedAt, getEntry.CreatedAt, time.Second)
}

func TestListEntry(t *testing.T) {
	account := createRandomAccount(t)
	//for loop create entry
	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}

	arg := ListEntryParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}

	entrys, err := testQueries.ListEntry(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entrys, 5)

	for _, entry := range entrys {
		require.NotEmpty(t, entry)
	}
}
