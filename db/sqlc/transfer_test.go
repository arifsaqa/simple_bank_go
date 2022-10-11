package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/arifsaqa/simple_bank_go/utils"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	fromAcc:= createRandomAccount(t)
	toAcc:= createRandomAccount(t)
	
	arg:= CreateTransferParams{
		FromAccountID: fromAcc.ID,
		ToAccountID: toAcc.ID,
		Amount: utils.RandomMoney(),
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

func TestCreateTransfer(t *testing.T){
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T)  {
	randTransfer:=createRandomTransfer(t)
	resTransfer, err:= testQueries.GetTransfer(context.Background(), randTransfer.ID)
	require.NoError(t,err)
	require.NotEmpty(t, resTransfer)

	require.Equal(t,randTransfer.ID, resTransfer.ID)
	require.Equal(t,randTransfer.FromAccountID, resTransfer.FromAccountID)
	require.Equal(t,randTransfer.ToAccountID, resTransfer.ToAccountID)
	require.Equal(t,randTransfer.Amount, resTransfer.Amount)
	require.WithinDuration(t,randTransfer.CreatedAt, resTransfer.CreatedAt, time.Second)
}

func TestUpdateTransfer(t *testing.T)  {
	randTransfer:=createRandomTransfer(t)
	arg:=UpdateTransferParams{
		ID: randTransfer.ID,
		Amount: utils.RandomMoney(),
	}
	resTransfer, err:= testQueries.UpdateTransfer(context.Background(), arg)
	require.NoError(t,err)
	require.NotEmpty(t, resTransfer)

	require.Equal(t,randTransfer.ID, resTransfer.ID)
	require.Equal(t,randTransfer.FromAccountID, resTransfer.FromAccountID)
	require.Equal(t,arg.Amount, resTransfer.Amount)
	require.WithinDuration(t,randTransfer.CreatedAt, resTransfer.CreatedAt, time.Second)
}

func TestDeleteTransfer(t *testing.T)  {
	account1:=createRandomTransfer(t)
	err:= testQueries.DeleteTransfer(context.Background(), account1.ID)
	require.NoError(t,err)

	resTransfer, err1:= testQueries.GetTransfer(context.Background(), account1.ID)

	require.Error(t, err1)
	require.EqualError(t, err1, sql.ErrNoRows.Error())
	require.Empty(t, resTransfer)
}

func TestListTransfers(t *testing.T)  {
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)		
	}

	arg:=ListTransfersParams{Limit: 5, Offset: 5}
	transfers, err:= testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t,err)
	for _, transfer:=range transfers {
		require.NotEmpty(t, transfer)
	}
}
