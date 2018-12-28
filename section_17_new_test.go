package main

// import (
// 	"fmt"
// 	"reflect"
// 	"testing"
// )

// type (
// 	fakeDriverProvider struct {
// 		driverByID func(DriverID) (*Driver, error)
// 	}

// 	fakeDriverSaver struct {
// 		save func(Driver) error
// 	}
// )

// func (m fakeDriverProvider) DriverByID(id DriverID) (*Driver, error) {
// 	if m.driverByID == nil {
// 		return nil, fmt.Errorf("mock error")
// 	}
// 	return m.driverByID(id)
// }

// func (m fakeDriverSaver) Save(d Driver) error {
// 	if m.save == nil {
// 		return fmt.Errorf("mock error")
// 	}
// 	return m.save(d)
// }

// func TestDriverService_DriverByID_success(t *testing.T) {
// 	want := &Driver{ID: 1, Name: "Jules"}
// 	service := DriverService{
// 		Provider: fakeDriverProvider{driverByID: func(DriverID) (*Driver, error) {
// 			return want, nil
// 		}},
// 		Saver: fakeDriverSaver{},
// 	}

// 	got, err := service.DriverByID(want.ID)

// 	if err != nil {
// 		t.Errorf("DriverByID(%v) error: %v", want.ID, err)
// 	}
// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("DriverByID(%v) = %v; want %v", want.ID, got, want)
// 	}
// }

// func TestDriverService_DriverByID_failsWhenDriverNotFound(t *testing.T) {
// 	id := DriverID(123)
// 	want := fmt.Errorf("driver not found")
// 	service := DriverService{
// 		Provider: fakeDriverProvider{driverByID: func(DriverID) (*Driver, error) {
// 			return nil, want
// 		}},
// 		Saver: fakeDriverSaver{},
// 	}

// 	driver, got := service.DriverByID(id)

// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("DriverByID(%v) error = %v ; want %v", id, got, want)
// 	}
// 	if driver != nil {
// 		t.Errorf("DriverByID(%v) = %v; want nil", id, driver)
// 	}
// }
