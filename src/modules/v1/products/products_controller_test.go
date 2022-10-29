package products

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zazhedho/telkom-test/task6-api/src/database/orm/models"
	"github.com/zazhedho/telkom-test/task6-api/src/helpers"
)

var repos = RepoProductMock{mock.Mock{}}
var service = NewService(&repos)
var ctrl = NewCtrl(service)

func TestCtrlGetAllProduct(t *testing.T) {
	w := httptest.NewRecorder()
	mux := mux.NewRouter()

	var dataMock = models.Products{
		{ProductId: 1, Name: "test 1"},
		{ProductId: 2, Name: "test 2"},
	}

	repos.mock.On("FindAllProducts").Return(&dataMock, nil)
	mux.HandleFunc("/test/products", ctrl.GetAllProducts)

	mux.ServeHTTP(w, httptest.NewRequest("GET", "/test/products", nil))
	var product models.Products

	response := helpers.Response{
		Data: &product,
	}

	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	fmt.Println(response)

	assert.Equal(t, 200, w.Code, "status code must be 200")
	assert.False(t, response.IsError, "isError must be false")
}
