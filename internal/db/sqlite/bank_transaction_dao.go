package sqlite

import (
	"bank/internal/constants"
	"bank/internal/models/response"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"time"
)

type BankTransactionDao struct {
	db *sql.DB
}

func (b *BankTransactionDao) InitializeDb(db *sql.DB) {
	b.db = db
}

func (b *BankTransactionDao) GetDetailByAccNumber(accountNumber string) (*response.AccountDetail, error) {
	var accountDetail response.AccountDetail
	err := b.db.QueryRow(constants.FETCH_ACCOUNT_DETAIL_QUERY, accountNumber).Scan(&accountDetail.AccNumber, &accountDetail.Balance)
	if err != nil {
		return nil, err
	}
	//Todo: Also add transaction list details also
	return &accountDetail, nil
}

func (b *BankTransactionDao) TransferMoney(fromAccNo string, toAccNo string, transferAmount float64) (*response.TransferAmountResponse, error) {

	tx, err := b.db.Begin()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	var fromAccount response.AccountDetail
	var toAccount response.AccountDetail

	//Locking row for FromAccount in Account table
	rows, err := tx.Query(constants.ROW_LOCK_ACCOUNT_AND_GET_QUERY, fromAccNo)

	defer rows.Close()

	fromAccountFound := false

	for rows.Next() {
		if fromAccountFound {
			return nil, errors.New(constants.ERROR_INTERNAL_SERVER)
		}
		if err := rows.Scan(&fromAccount.AccNumber, &fromAccount.Balance); err != nil {
			return nil, errors.New(constants.ERROR_INTERNAL_SERVER)
		}
		fromAccountFound = true
	}

	if fromAccountFound == false {
		// no rows found
		return nil, errors.New(constants.ERROR_FROM_ACCOUNT_NOT_EXISTS)
	}

	if fromAccount.Balance-transferAmount < 0 {
		tx.Rollback()
		return nil, errors.New(constants.ERROR_INSUFFICIENT_BALANCE)
	}

	// Subtract the amount from the fromAccount balance
	_, err = tx.Exec(constants.SUBTRACT_AMOUNT_TO_ACCOUBT_QUERY, transferAmount, fromAccNo)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	//Locking row for FromAccount in Account table
	rows, err = tx.Query(constants.ROW_LOCK_ACCOUNT_AND_GET_QUERY, toAccNo)

	defer rows.Close()

	toAccountFound := false

	for rows.Next() {
		if toAccountFound {
			return nil, errors.New(constants.ERROR_INTERNAL_SERVER)
		}
		if err := rows.Scan(&toAccount.AccNumber, &toAccount.Balance); err != nil {
			return nil, errors.New(constants.ERROR_INTERNAL_SERVER)
		}
		toAccountFound = true
	}

	if toAccountFound == false {
		// no rows found
		return nil, errors.New(constants.ERROR_TO_ACCOUNT_NOT_EXISTS)
	}

	// Add the amount to the toAccount balance
	_, err = tx.Exec(constants.ADD_AMOUNT_TO_ACCOUBT_QUERY, transferAmount, toAccNo)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Save the transaction details in the transactions table, sequence for insert
	//(transaction_id, from_acc_no, to_acc_no,transferred_amount, created_datetime)VALUES (?, ?, ?, ?, ?)
	now := time.Now()
	timestamp := now.Format(constants.DATE_FORMAT)
	transactionId := uuid.New().String()
	_, err = tx.Exec(constants.INSERT_ONE_TRANSACTION_DETAIL_QUERY, transactionId, fromAccNo, toAccNo, transferAmount, timestamp)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	var transactionDto = response.TransferAmountResponse{
		Id:          transactionId,
		Transferred: transferAmount,
		From: response.AccountDetailResponse{
			Id:      fromAccNo,
			Balance: fromAccount.Balance - transferAmount,
		},
		To: response.AccountDetailResponse{
			Id:      toAccNo,
			Balance: toAccount.Balance + transferAmount,
		},
		CreatedDatetime: now,
	}

	return &transactionDto, nil
}
