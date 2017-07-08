package main

// json parse用構造体
type User struct {
	primaryKey int
	uname      string //uniqueな名前
	upass      string //hash
}
