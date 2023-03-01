package internal

import (
	"bank/internal/constants"
	"bank/internal/db"
	"bank/internal/db/sqlite"
	"bank/internal/models/requests"
	"bank/internal/models/response"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type BankManger struct {
	executor db.BankExecutor
}

var BankManagerInstance *BankManger

func init() {
	InitializeBankManager()
}

func InitializeBankManager() {
	// Initialize the singleton instance of the BankService struct
	dbInstance, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/bank_db")
	if err != nil {
		panic(fmt.Sprintf(constants.ERROR_DB_CONNECTION, err))
	}
	executor := sqlite.BankTransactionDao{}
	executor.InitializeDb(dbInstance)
	BankManagerInstance = &BankManger{executor: &executor}
	fmt.Println("Db initialized . . . ")
}

func GetBankManagerInstance() *BankManger {
	if BankManagerInstance == nil {
		InitializeBankManager()
	}
	return BankManagerInstance
}

func (b *BankManger) GetAccDetails(accNumber string) (*response.AccountDetail, error) {
	return b.executor.GetDetailByAccNumber(accNumber)
}

func (b *BankManger) TransferMoney(requestBody requests.TransferAmountRequest) (*response.TransferAmountResponse, error) {
	return b.executor.TransferMoney(requestBody.FromAccNo, requestBody.ToAccNo, requestBody.Amount)
	//if err != nil {
	//	return nil, err
	//}
	//return response.ConvertDTOToResponseModel(*transactionDto)
}
