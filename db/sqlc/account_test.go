package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/arifsaqa/simple_bank_go/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg:= CreateAccountParams{
		Owner: utils.RandomOwner(),
		Balance: utils.RandomMoney(),
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

func TestCreateAccount(t *testing.T){
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T)  {
	account1:=createRandomAccount(t)
	acc, err:= testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t,err)
	require.NotEmpty(t, acc)

	require.Equal(t,account1.ID, acc.ID)
	require.Equal(t,account1.Owner, acc.Owner)
	require.Equal(t,account1.Balance, acc.Balance)
	require.Equal(t,account1.Currency, acc.Currency)
	require.WithinDuration(t,account1.CreatedAt, acc.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T)  {
	account1:=createRandomAccount(t)
	arg:=UpdateAccountParams{
		ID: account1.ID,
		Balance: utils.RandomMoney(),
	}
	acc, err:= testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t,err)
	require.NotEmpty(t, acc)

	require.Equal(t,account1.ID, acc.ID)
	require.Equal(t,account1.Owner, acc.Owner)
	require.Equal(t,arg.Balance, acc.Balance)
	require.Equal(t,account1.Currency, acc.Currency)
	require.WithinDuration(t,account1.CreatedAt, acc.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T)  {
	account1:=createRandomAccount(t)
	err:= testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t,err)

	account2, err1:= testQueries.GetAccount(context.Background(), account1.ID)

	require.Error(t, err1)
	require.EqualError(t, err1, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccounts(t *testing.T)  {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)		
	}

	arg:=ListAccountsParams{Limit: 5, Offset: 5}
	accounts, err:= testQueries.ListAccounts(context.Background(), arg)

	require.NoError(t,err)
	for _, account:=range accounts {
		require.NotEmpty(t, account)
	}

	
}
