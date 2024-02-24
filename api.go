package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error{
	
}
type apiFunc func(http.ResponseWriter,*http.Request) error

func makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		if err := f(w,r); 
		err != nil {
			// handle the error
		}
	}
}
type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}


func (s *APIServer) Run(){
	router := mux.NewRouter()
	router.HandleFunc("/account", s.handleAccount).Methods("POST")

}
func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Response) error{
return nil
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Response) error{
return nil
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Response) error{
return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Response) error{
return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Response) error{
return nil
}