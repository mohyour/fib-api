package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	fib "pex/app/fibonacci"
	"pex/config"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var (
	// Session instatiation
	key   = []byte(config.EnvVariable("SECRET"))
	store = sessions.NewCookieStore(key)

	// First number in fibonacci sequence - f(0)
	fibNumber int = 0

	// WarningLogger ...
	WarningLogger *log.Logger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	// InfoLogger ...
	InfoLogger *log.Logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	// ErrorLogger ...
	ErrorLogger *log.Logger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)

// GetCurrentFibSequenceHandler - get current number in the fibonacci sequence
func GetCurrentFibSequenceHandler(w http.ResponseWriter, r *http.Request) {
	setResponseHeader(w)
	log.Println("Getting current fibonacci sequence...")
	session, _ := store.Get(r, "fib-cookie")

	// Get value from user session
	_, ok := session.Values["started"].(bool)
	fibNumber, _ := session.Values["current-fibnum"].(int)

	if !ok {
		fibNumber = 0
		session.Values["started"] = true
	}

	writeResponse(fibNumber+1, fibNumber, session, w, r)
}

// GetNextFibSequenceHandler - get next number in the fibonacci sequence
func GetNextFibSequenceHandler(w http.ResponseWriter, r *http.Request) {
	setResponseHeader(w)
	log.Println("Getting next fibonacci sequence...")
	// Get value from user session
	session, _ := store.Get(r, "fib-cookie")
	fibNumber, _ := session.Values["current-fibnum"].(int)

	fibNumber++

	writeResponse(fibNumber+1, fibNumber, session, w, r)
}

// GetPreviousFibSequenceHandler - get previous number in the fibonacci sequence
func GetPreviousFibSequenceHandler(w http.ResponseWriter, r *http.Request) {
	setResponseHeader(w)
	log.Println("Getting previous fibonacci sequence...")
	// Get value from user session
	session, _ := store.Get(r, "fib-cookie")
	fibNumber, _ := session.Values["current-fibnum"].(int)

	if fibNumber < 1 {
		fibNumber = 0
		WarningLogger.Println("Fibonacci term cannot be lower than 1 - f(0)")
	} else {
		fibNumber--
	}

	writeResponse(fibNumber+1, fibNumber, session, w, r)
}

// GetNthTermOfSequenceHandler - get nth term of fibonacci sequence
// first term = f(0), second term = f(1)
func GetNthTermOfSequenceHandler(w http.ResponseWriter, r *http.Request) {
	setResponseHeader(w)
	log.Println("Getting nth term fibonacci sequence...")
	// Get value from user session
	session, _ := store.Get(r, "fib-cookie")

	vars := mux.Vars(r)
	term, err := strconv.Atoi(vars["term"])
	if err != nil || term < 1 {
		http.Error(w, "Invalid number: Number can only be a positive integer", http.StatusBadRequest)
		ErrorLogger.Printf("Invalid term: %s", vars["term"])
		return
	}

	writeResponse(term, term-1, session, w, r)
}

// ResetFibSequenceHandler - reset fibonacci sequence to the 1st term
func ResetFibSequenceHandler(w http.ResponseWriter, r *http.Request) {
	setResponseHeader(w)
	log.Println("Resetting fibonacci sequence...")
	session, _ := store.Get(r, "fib-cookie")

	writeResponse(1, 0, session, w, r)
}

// writeResponse - marshal response to JSON
func writeResponse(term, num int, session *sessions.Session, w http.ResponseWriter, r *http.Request) {
	response := fib.FibResponse{
		FibTerm:          term,
		FibSequenceValue: fib.Fibonacci(num),
	}
	InfoLogger.Printf("{Term: %d, FibSequence: %s}", term, fib.Fibonacci(num))

	// Update term and fibnumber to user session
	updateFibTerm(num, session, w, r)

	// Format response to Json
	res, err := json.Marshal(response)
	if err != nil {
		ErrorLogger.Println(err)
		http.Error(w, "Json marshal error", http.StatusInternalServerError)
	}
	// Write response to http
	_, err = w.Write(res)
}

// updateFibTerm - update values in session
func updateFibTerm(fibNumber int, session *sessions.Session, w http.ResponseWriter, r *http.Request) {
	session.Values["current-fibnum"] = fibNumber
	session.Values["started"] = true
	session.Save(r, w)
}

// setResponseHeader - response header to allow CORS
func setResponseHeader(w http.ResponseWriter) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Content-Type", "application/json")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
