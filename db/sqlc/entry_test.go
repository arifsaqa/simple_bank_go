package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/arifsaqa/simple_bank_go/utils"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	randEntry:=createRandomAccount(t);
	arg:= CreateEntryParams{
		AccountID: randEntry.ID,
		Amount: utils.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.NotZero(t, randEntry.ID)

	require.Equal(t, arg.AccountID, entry.AccountID)

	require.NotZero(t, randEntry.ID)
	require.NotZero(t, arg.Amount)

	return entry
}

func TestCreateEntry(t *testing.T){
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T)  {
	randEntry:=createRandomEntry(t)
	resEntry, err:= testQueries.GetEntry(context.Background(), randEntry.ID)
	require.NoError(t,err)
	require.NotEmpty(t, resEntry)

	require.Equal(t,randEntry.ID, resEntry.ID)
	require.Equal(t,randEntry.AccountID, resEntry.AccountID)
	require.Equal(t,randEntry.Amount, resEntry.Amount)
	require.WithinDuration(t,randEntry.CreatedAt, resEntry.CreatedAt, time.Second)
}

func TestUpdateEntry(t *testing.T)  {
	randEntry:=createRandomEntry(t)
	arg:=UpdateEntryParams{
		ID: randEntry.ID,
		Amount: utils.RandomMoney(),
	}
	resEntry, err:= testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t,err)
	require.NotEmpty(t, resEntry)

	require.Equal(t,randEntry.ID, resEntry.ID)
	require.Equal(t,randEntry.AccountID, resEntry.AccountID)
	require.Equal(t,arg.Amount, resEntry.Amount)
	require.WithinDuration(t,randEntry.CreatedAt, resEntry.CreatedAt, time.Second)
}

func TestDeleteEntry(t *testing.T)  {
	randEntry:=createRandomEntry(t)
	err:= testQueries.DeleteEntry(context.Background(), randEntry.ID)
	require.NoError(t,err)

	randEntry2, err1:= testQueries.GetEntry(context.Background(), randEntry.ID)

	require.Error(t, err1)
	require.EqualError(t, err1, sql.ErrNoRows.Error())
	require.Empty(t, randEntry2)
}

func TestListEntrys(t *testing.T)  {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)		
	}

	arg:=ListEntriesParams{Limit: 5, Offset: 5}
	entries, err:= testQueries.ListEntries(context.Background(), arg)

	require.NoError(t,err)
	for _, entry:=range entries {
		require.NotEmpty(t, entry)
	}
}
