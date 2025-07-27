package db

import (
	"context"
	"simple-bank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, accountID1, accountID2 int64) Transfer {
	args := CreateTransferParams{
		FromAccountID: accountID1,
		ToAccountID:   accountID2,
		Amount:        util.RandomBalance(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, args.ToAccountID, transfer.ToAccountID)
	require.Equal(t, args.FromAccountID, transfer.FromAccountID)
	require.Equal(t, args.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	createRandomTransfer(t, account1.ID, account2.ID)
}

func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	transfer := createRandomTransfer(t, account1.ID, account2.ID)

	findTransfer, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, findTransfer)

	require.Equal(t, transfer.ID, findTransfer.ID)
	require.Equal(t, transfer.ToAccountID, findTransfer.ToAccountID)
	require.Equal(t, transfer.FromAccountID, findTransfer.FromAccountID)
	require.Equal(t, transfer.Amount, findTransfer.Amount)
	require.WithinDuration(t, transfer.CreatedAt, findTransfer.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 5; i++ {
		createRandomTransfer(t, account1.ID, account2.ID)
	}

	args := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Limit:         10,
		Offset:        0,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.True(t, transfer.FromAccountID == account1.ID)
	}
}
