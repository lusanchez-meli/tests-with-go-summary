package main

type DriverService struct {
	Provider interface {
		DriverByID(DriverID) (*Driver, error)
	}
	Saver interface {
		Save(Driver) error
	}
}

func (s DriverService) DriverByID(id DriverID) (*Driver, error) {
	return s.Provider.DriverByID(id)
}

func (s DriverService) Save(d Driver) error {
	return s.Saver.Save(d)
}

type DriverID uint

type Driver struct {
	ID   DriverID
	Name string
}
