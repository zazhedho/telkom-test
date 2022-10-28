package products

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/api/v1/products").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/name", ctrl.SortByName).Methods("GET")
	route.HandleFunc("", ctrl.GetAllProducts).Methods("GET")
	route.HandleFunc("", ctrl.AddProduct).Methods("POST")
	route.HandleFunc("/{id}", ctrl.UpdateProduct).Methods("PUT")
	route.HandleFunc("/{kode}", ctrl.DeleteProduct).Methods("DELETE")
}
