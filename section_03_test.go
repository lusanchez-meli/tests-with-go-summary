package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/mercadolibre/shipping-logistics-drivers/src/api/dao"
	"github.com/mercadolibre/shipping-logistics-drivers/src/api/domain"
	"github.com/mercadolibre/shipping-logistics-drivers/src/api/rest"
	"github.com/mercadolibre/shipping-logistics-drivers/src/api/tests"
)

func TestCarrierRepository_FindCarrierByID_success(t *testing.T) {
	var mockId uint = 123
	mockStatus := "active"
	responseBytes := fmt.Sprintf(`{"id":%v,"status":"%v"}`, mockId, mockStatus)

	mockResponse := &responseMock{
		statusCode: 200,
		bytes:      []byte(responseBytes),
	}
	cr, err := dao.NewCarrierRepository(mockResponse)
	tests.FailOnError(t, err)
	wantCarrier := &domain.Carrier{
		ID:     domain.CarrierID(mockId),
		Status: domain.CarrierStatus(mockStatus),
	}

	gotCarrier, err := cr.FindCarrierByID(domain.CarrierID(123))

	tests.FailOnError(t, err)
	tests.AssertEquals(t, wantCarrier, gotCarrier)
}

func TestCarrierRepository_FindCarrierByID_successWithWant(t *testing.T) {
	want := &domain.Carrier{
		ID:     domain.CarrierID(123),
		Status: domain.CarrierStatus("active"),
	}
	repo, err := dao.NewCarrierRepository(&responseMock{
		statusCode: http.StatusOK,
		bytes: []byte(fmt.Sprintf(`{
			"id": %d,
			"status": %q
		}`, want.ID, want.Status)),
	})
	tests.FailOnError(t, err)

	got, err := repo.FindCarrierByID(want.ID)

	tests.AssertNil(t, err)
	tests.AssertEquals(t, want, got)
}

type responseMock struct {
	string     string
	bytes      []byte
	statusCode int
}

func (r *responseMock) String() string {
	return r.string
}

func (r *responseMock) Bytes() []byte {
	return r.bytes
}

func (r *responseMock) StatusCode() int {
	return r.statusCode
}

func (r *responseMock) Get(url string) rest.Response {
	return r
}
