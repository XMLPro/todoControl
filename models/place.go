package models

import "errors"

type Place struct {
	PlaceId    int    `json:"place_id"`
	PlaceName  string `json:"place_name"`
	Importance int    `json:"importance"`
}

func (m *Place) Valid() error {
	if m.PlaceName == ""{
		return errors.New("Place Error: content is none ")
	}

	if m.Importance < 0 || m.Importance > 100 {
		return errors.New("Place Error: importance is out of range ")
	}

	return nil
}
