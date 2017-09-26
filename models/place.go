package models

type Place struct {
	Place_id   int
	Place_name string
	Importance string
}

func (m *Place) Valid() error {
	return nil
}
