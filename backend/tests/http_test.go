package tests

import (
	"fmt"
	"github.com/Proxypepe/wb-web/backend/cache"
	"github.com/Proxypepe/wb-web/backend/http"
	"github.com/Proxypepe/wb-web/backend/schemas"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	net "net/http"
	"net/http/httptest"
	"testing"
)

func TestGetExistsOrderByUid(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)
	teardownSuiteDB := setupSuiteDB(t)
	defer teardownSuiteDB(t)

	orderUid := "b563feb7b2b84b6test"
	order := GetExampleOrder(orderUid)
	err := cache.SaveOrder(order)
	assert.Nil(t, err)
	router := http.NewServer()
	w := httptest.NewRecorder()
	req, _ := net.NewRequest("GET", fmt.Sprintf("/order?order_uid=%s", orderUid), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var orderResponse schemas.Order
	err = json.Unmarshal(w.Body.Bytes(), &orderResponse)

	assert.Equal(t, order, orderResponse)
}

func TestGetNotExistsOrderByUid(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)
	teardownSuiteDB := setupSuiteDB(t)
	defer teardownSuiteDB(t)

	router := http.NewServer()
	w := httptest.NewRecorder()
	req, _ := net.NewRequest("GET", "/order?order_uid=b563feb7b2b84b6test1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}
