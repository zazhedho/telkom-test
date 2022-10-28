package routers

import (
	"errors"

	"github.com/gorilla/mux"
	"github.com/zazhedho/telkom-test/task6-api/src/database/orm"
	"github.com/zazhedho/telkom-test/task6-api/src/modules/v1/products"
)

func New() (*mux.Router, error) {
	mainRoute := mux.NewRouter()
	db, err := orm.New()
	if err != nil {
		return nil, errors.New("gagal init database")
	}

	products.New(mainRoute, db)

	return mainRoute, nil
}
