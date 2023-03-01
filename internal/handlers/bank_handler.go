package handlers

import (
	"bank/internal"
	"bank/internal/constants"
	"bank/internal/models/requests"
	"bank/internal/models/response"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func AccountDetailHandler(c *gin.Context) {
	w, _ := c.Writer, c.Request
	accNumber := c.Param(constants.URL_PARAMS_ACC_NO)
	bankManager := internal.GetBankManagerInstance()

	accDetail, err := bankManager.GetAccDetails(accNumber)
	if err != nil {
		setErrorToWriter(w, err, http.StatusInternalServerError)
		return
	}

	// Convert the user struct to a JSON response
	jsonBytes, err := json.Marshal(accDetail)
	if err != nil {
		setErrorToWriter(w, err, http.StatusInternalServerError)
		return
	}

	// Set the content type header and write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func TransferMoneyHandler(c *gin.Context) {
	w, r := c.Writer, c.Request
	bankManager := internal.GetBankManagerInstance()
	accNumberParam := c.Param(constants.URL_PARAMS_ACC_NO)
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		setErrorToWriter(w, err, http.StatusInternalServerError)
		return
	}

	// Decode the request body into a struct
	var requestBody requests.TransferAmountRequest
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		setErrorToWriter(w, err, http.StatusInternalServerError)
		return
	}

	if requestBody.FromAccNo != accNumberParam || requestBody.FromAccNo == requestBody.ToAccNo {
		err = errors.New(constants.ERROR_INVALID_OPERATION)
		setErrorToWriter(w, err, http.StatusForbidden)
		return
	}

	transactionResponse, err := bankManager.TransferMoney(requestBody)
	if err != nil {
		setErrorToWriter(w, err, http.StatusInternalServerError)
		return
	}

	// Convert the user struct to a JSON response
	jsonBytes, err := json.Marshal(transactionResponse)
	if err != nil {
		setErrorToWriter(w, err, http.StatusInternalServerError)
		return
	}

	// Set the content type header and write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func setErrorToWriter(w gin.ResponseWriter, err error, errorCode int) {
	w.WriteHeader(errorCode)
	errorBytes := ErrorResponseInBytes(errorCode, err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(errorBytes)
}

func ErrorResponseInBytes(errorCode int, err error) []byte {
	error := response.ErrorResponse{ErrorCode: errorCode, Message: err.Error()}
	jsonBytes, err := json.Marshal(error)
	return jsonBytes
}
