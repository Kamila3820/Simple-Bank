package db

import (
	"context"
	"simple-bank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: "secret",
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangedAt.IsZero())

	return user
}
func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)

	findUser, err := testQueries.GetUser(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, findUser)

	require.Equal(t, user.Username, findUser.Username)
	require.Equal(t, user.HashedPassword, findUser.HashedPassword)
	require.Equal(t, user.FullName, findUser.FullName)
	require.Equal(t, user.Email, findUser.Email)
	require.WithinDuration(t, user.PasswordChangedAt, findUser.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user.CreatedAt, findUser.CreatedAt, time.Second)
}
