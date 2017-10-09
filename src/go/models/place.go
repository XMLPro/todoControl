package models

type Place struct {
	Id         int    `json:"id"`
	PlaceName  string `json:"place_name"`
	Importance int    `json:"importance"`
}
