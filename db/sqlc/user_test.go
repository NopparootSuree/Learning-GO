package db

import (
	"context"
	"testing"
	"time"

	"github.com/NopparootSuree/Learning-GO/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := utils.HashPassword(utils.RandomString(6))
	require.NoError(t, err)
	arg := CreateUserParams{
		Username:       utils.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       utils.RandomOwner(),
		Email:          utils.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	getUser, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, getAccount)

	require.Equal(t, user1.Username, getUser.Username)
	require.Equal(t, user1.HashedPassword, getUser.HashedPassword)
	require.Equal(t, user1.FullName, getUser.FullName)
	require.Equal(t, user1.Email, getUser.Email)
	require.WithinDuration(t, user1.PasswordChangedAt, getUser.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, getUser.CreatedAt, time.Second)
}
