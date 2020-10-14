package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"net/http/httptest"
	"os"
	fib "pex/app/fibonacci"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Remove logging from unit test
	InfoLogger.SetOutput(ioutil.Discard)
	ErrorLogger.SetOutput(ioutil.Discard)
	WarningLogger.SetOutput(ioutil.Discard)

	os.Exit(m.Run())
}

func TestGetCurrentFibHandler(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	testWriter, testRequest := setupWriterAndRequest()
	handler := GetCurrentFibSequenceHandler
	handler(testWriter, testRequest)

	fibResponse := &fib.FibResponse{}
	err := json.Unmarshal(testWriter.Body.Bytes(), fibResponse)
	code := testWriter.Code

	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, code, "expect 200 status code")
	assert.Equal(t, 1, fibResponse.FibTerm, "Inocrrect term")
	assert.Equal(t, big.NewInt(0), fibResponse.FibSequenceValue, "Incorrect sequence")
}

func TestGetNextFibHandler(t *testing.T) {
	testWriter, testRequest := setupWriterAndRequest()
	handler := GetNextFibSequenceHandler
	handler(testWriter, testRequest)

	fibResponse := &fib.FibResponse{}
	err := json.Unmarshal(testWriter.Body.Bytes(), fibResponse)
	code := testWriter.Code

	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, code, "expect 200 status code")
	assert.Equal(t, 2, fibResponse.FibTerm, "Inocrrect term")
	assert.Equal(t, big.NewInt(1), fibResponse.FibSequenceValue, "Incorrect sequence")
}

func TestGetPreviousFibHandler(t *testing.T) {
	testWriter, testRequest := setupWriterAndRequest()
	handler := GetPreviousFibSequenceHandler
	handler(testWriter, testRequest)

	fibResponse := &fib.FibResponse{}
	err := json.Unmarshal(testWriter.Body.Bytes(), fibResponse)
	code := testWriter.Code

	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, code, "expect 200 status code")
	assert.Equal(t, 1, fibResponse.FibTerm, "Inocrrect term")
	assert.Equal(t, big.NewInt(0), fibResponse.FibSequenceValue, "Incorrect sequence")
}

func TestResetFibHandler(t *testing.T) {
	testWriter, testRequest := setupWriterAndRequest()
	handler := ResetFibSequenceHandler
	handler(testWriter, testRequest)

	fibResponse := &fib.FibResponse{}
	err := json.Unmarshal(testWriter.Body.Bytes(), fibResponse)
	code := testWriter.Code

	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, code, "expect 200 status code")
	assert.Equal(t, 1, fibResponse.FibTerm, "Inocrrect term")
	assert.Equal(t, big.NewInt(0), fibResponse.FibSequenceValue, "Incorrect sequence")
}

func TestGetNextFibSequenceHandler(t *testing.T) {
	tt := []struct {
		term        string
		fibSequence int64
	}{
		{"90", 1779979416004714189},
		{"5", 3},
		{"1", 0},
		{"nine", -1}, // -1 value to keep the struct valid
		{"58as", -1},
	}

	handler := GetNthTermOfSequenceHandler

	for _, tc := range tt {
		testWriter, testRequest := setupWriterAndRequest()
		testRequest = mux.SetURLVars(testRequest, map[string]string{
			"term": tc.term,
		})
		handler(testWriter, testRequest)
		code := testWriter.Code

		if code == http.StatusBadRequest {
			assert.Equal(t, http.StatusBadRequest, code, "expect a 400 status code")
			assert.Equal(t, testWriter.Body.String(), "Invalid number: Number can only be a positive integer\n")
		} else {
			fibResponse := &fib.FibResponse{}
			err := json.Unmarshal(testWriter.Body.Bytes(), fibResponse)
			assert.NoError(t, err)

			res, err := strconv.Atoi(tc.term)
			assert.NoError(t, err)

			assert.Equal(t, http.StatusOK, code, "expect a 200 status code")
			assert.Equal(t, res, fibResponse.FibTerm, "Inocrrect term")
			assert.Equal(t, big.NewInt(tc.fibSequence), fibResponse.FibSequenceValue, "Invalid fibonacci sequence")
		}
	}

}
func setupWriterAndRequest() (*httptest.ResponseRecorder, *http.Request) {
	testWriter := httptest.NewRecorder()
	testRequest := new(http.Request)
	return testWriter, testRequest
}
