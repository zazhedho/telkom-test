package products

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/zazhedho/telkom-test/task6-api/src/database/orm/models"
	"github.com/zazhedho/telkom-test/task6-api/src/helpers"
	"github.com/zazhedho/telkom-test/task6-api/src/interfaces"
)

type product_ctrl struct {
	svc interfaces.ProductService
}

func NewCtrl(ctrl interfaces.ProductService) *product_ctrl {
	return &product_ctrl{ctrl}
}

func (c *product_ctrl) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	c.svc.GetAllProducts().Send(w)
}

func (c *product_ctrl) AddProduct(w http.ResponseWriter, r *http.Request) {

	var data models.Product
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		helpers.New(err, 500, true)
		return
	}
	c.svc.AddProduct(&data).Send(w)
}

func (c *product_ctrl) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	var data models.Product
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		helpers.New(err, 500, true)
		return
	}

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		helpers.New(err, 500, true)
		return
	}

	c.svc.UpdateProduct(id, &data).Send(w)
}

func (c *product_ctrl) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var data models.Product

	params := mux.Vars(r)
	kode := strings.ToLower(params["kode"])

	c.svc.DeleteProduct(kode, &data).Send(w)
}

func (c *product_ctrl) SortByName(w http.ResponseWriter, r *http.Request) {

	var data models.Products
	name := strings.ToLower(r.URL.Query().Get("name"))
	c.svc.SortByName(name, &data).Send(w)
}
