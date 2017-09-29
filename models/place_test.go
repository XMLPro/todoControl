package models

import "testing"

func TestPlace_Valid(t *testing.T) {
	place := Place{PlaceId:0 , PlaceName: "name",Importance:0}
	if err := place.Valid(); err != nil {
		t.Error(err)
	}

	//non name
	place = Place{PlaceId:1, PlaceName: "", Importance:0 }
	if err := place.Valid(); err == nil{
		t.Error(err)
	}

	place = Place{PlaceId:2, PlaceName: "name2", Importance:-1}
	if err := place.Valid(); err == nil{
		t.Error(err)
	}
}
