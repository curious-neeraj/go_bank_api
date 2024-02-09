package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// func apiServer() {
// 	fmt.Println("Initialize API Server!")
// }

type APIServer struct {
	listenAddress string
	store         Storage
}

func NewAPIServer(listenAddress string, store Storage) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
		store:         store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccountById))

	log.Println("API Server running on port ", s.listenAddress)
	http.ListenAndServe(s.listenAddress, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetAccount(w, r)
	case "POST":
		return s.handleCreateAccount(w, r)
	case "DELETE":
		return s.handleDeleteAccount(w, r)
	default:
		return fmt.Errorf("%s - method not allowed", r.Method)
	}
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.store.GetAccounts()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, accounts)
	// return nil
}

func (s *APIServer) handleGetAccountById(w http.ResponseWriter, r *http.Request) error {
	account := NewAccount("Neeraj", "Jain")
	return WriteJSON(w, http.StatusOK, account)

	// return nil
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	createAcReq := new(CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(createAcReq); err != nil {
		return err
	}

	account := NewAccount(createAcReq.FirstName, createAcReq.LastName)
	if err := s.store.CreateAccount(account); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// func (s *APIServer) handleTransaction(w http.ResponseWriter, r *http.Request) error {
// 	return nil
// }

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

type ApiError struct {
	Error string
}

type apiFunc func(http.ResponseWriter, *http.Request) error

// function to convert our apiFunc to Http.HandlerFunc
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			// handle error
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
