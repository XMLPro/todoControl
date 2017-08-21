package models

type User struct {
	Id       int
	Name     string `orm:"size(64);unique"`
	Email    string `orm:"size(64);unique"`
	Password string `orm:"size(64)"`
}
