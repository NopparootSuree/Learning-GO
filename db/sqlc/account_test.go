package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/NopparootSuree/Learning-GO/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    utils.RandomOwner(), // randomly generated?
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	getAccount, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, getAccount)

	require.Equal(t, account1.ID, getAccount.ID)
	require.Equal(t, account1.Owner, getAccount.Owner)
	require.Equal(t, account1.Balance, getAccount.Balance)
	require.Equal(t, account1.Currency, getAccount.Currency)
	require.WithinDuration(t, account1.CreatedAt, getAccount.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: utils.RandomMoney(),
	}

	updateAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updateAccount)

	require.Equal(t, account1.ID, updateAccount.ID)
	require.Equal(t, account1.Owner, updateAccount.Owner)
	require.Equal(t, arg.Balance, updateAccount.Balance)
	require.Equal(t, account1.Currency, updateAccount.Currency)
	require.WithinDuration(t, account1.CreatedAt, updateAccount.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	checkDeleted, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, checkDeleted)
}

func TestListAccount(t *testing.T) {
	var lastAccount Account
	//for loop CreateAccount
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	}

	arg := ListAccountParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccount(context.Background(), arg)
	fmt.Println(accounts)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}
}
