package db

import (
	"GOLANG/db/util"
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	_ "github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)
	// arg := CreateAccountParams{
	// 	Owner:    util.RandomOwner(),
	// 	Balance:  util.RandomMoney(),
	// 	Currency: util.RandomCurrency(),
	// }
	// account, err := testQueries.CreateAccount(context.Background(), arg)
	// require.NoError(t, err)
	// require.NotEmpty(t, account)
	// require.Equal(t, arg.Owner, account.Owner)
	// require.Equal(t, arg.Balance, account.Balance)
	// require.Equal(t, arg.Currency, account.Currency)

	// require.NotZero(t, account.ID)
	// require.NotZero(t, account.CreatedAt)

}

func CreateRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
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

func TestGetAccount(t *testing.T) {
	// var id int64
	account1 := CreateRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmptyf(t, account1, "Data Doest not Exit")
	require.Equalf(t, account2.Balance, account1.Balance, "Balance Not Equal")
	require.Equalf(t, account2.Currency, account1.Currency, "Currency Not Equal")
	require.Equalf(t, account2.Owner, account1.Owner, "owner not equal")
	require.Equalf(t, account2.ID, account2.ID, "Id not found")
	require.WithinDurationf(t, account1.CreatedAt, account2.CreatedAt, time.Second, "Duration not Correct")

}

func TestDeleteAccount(t *testing.T) {
	fmt.Print("Delete Test")
	account1 := CreateRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)

}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomAccount(t)
	}

	arg := ListAccountsParams{
		Owner:  "USMAN",
		Limit:  5,
		Offset: 0,
	}

	// testQueries.ListAccounts()
	accounts, err := testQueries.ListAccounts(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, accounts, 2)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}

}
