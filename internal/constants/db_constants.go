package constants

const SUBTRACT_AMOUNT_TO_ACCOUBT_QUERY = "UPDATE Account SET balance = balance - ? WHERE id = ?"
const ROW_LOCK_ACCOUNT_AND_GET_QUERY = "SELECT * FROM Account WHERE id = ? FOR UPDATE"
const ADD_AMOUNT_TO_ACCOUBT_QUERY = "UPDATE Account SET balance = balance + ? WHERE id = ?"
const INSERT_ONE_TRANSACTION_DETAIL_QUERY = "INSERT INTO Transaction (transaction_id, from_acc_no, to_acc_no, transferred_amount, created_datetime)VALUES (?, ?, ?, ?, ?)"
const FETCH_ACCOUNT_DETAIL_QUERY = "SELECT * FROM Account WHERE id = ?"
