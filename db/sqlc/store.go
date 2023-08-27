package db

import (
	"context"
	"database/sql"
	"fmt"
)

// store all functions related to transtions and DB
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore Creates a new Store
func NewStore(db *sql.DB) *Store {

	return &Store{
		db:      db,
		Queries: New(db),
	}

}

// exectx function executes with a db transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx.err: %v, rb err %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()

}

// contains all input params for transfer
type TransferTxParams struct {
	FromAccountId int64 `json:"from_account_Id"`
	ToAccountId   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entery"` // after the ammount is transferred
	ToEntry     Entry    `json:"to_entry"`    // after the amount is transferred
}

// Trafer performs many action / transfer from 1 acc to another
// Create a transfer record , add account entries , add update acc balance wihtin a single DB transaction
// func (store *Store) TransactionTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	//Create a transfer record
	err := store.execTx(ctx, func(q *Queries) error {

		var err error
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountId,
			ToAccountID:   arg.ToAccountId,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		// add account entries
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountId,
			Amount:    -arg.Amount,
		})

		if err != nil {
			return err
		}

		//to account
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			Amount:    arg.Amount,
			AccountID: arg.ToAccountId,
		})
		//update Ac :TODO

		account1, err := q.GetAccount(context.Background(), arg.FromAccountId)
		if err != nil {
			fmt.Println(err)
		}

		result.FromAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
			ID:      arg.FromAccountId,
			Balance: account1.Balance - arg.Amount,
		})

		account2, err := q.GetAccount(context.Background(), arg.ToAccountId)
		if err != nil {
			fmt.Println(err)
		}

		result.ToAccount, err = q.UpdateAccount(context.Background(), UpdateAccountParams{
			ID:      arg.ToAccountId,
			Balance: account2.Balance + arg.Amount,
		})

		if err != nil {
			return err
		}

		return nil

	}) //belongs to Line : err := store.execTx(ctx, func(q *Queries) error

	return result, err
}
