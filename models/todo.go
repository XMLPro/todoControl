package models

type Todo struct {
	User    *User
	Content string
}

type PlaceImp struct {
	User       *User
	Place      string
	Importance int
}
