package models

type Place struct {
	Place_id   int    `json:"place_id"`
	Place_name string `json:"place_name"`
	Importance string `json:"importance"`
}

func (m *Place) Valid() error {
	return nil
}
