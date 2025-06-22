// Package customer provides HTTP routes for the customer domain.
package customer

import "net/http"

type CustomerHandler interface {
	CreateCustomer(w http.ResponseWriter, r *http.Request)
}

func RegisterRoutes(mux *http.ServeMux, handler CustomerHandler) {
	mux.HandleFunc("/customers", handler.CreateCustomer)
}
