package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	Id      int
	Content string
	Done    bool
}

/*
type PlaceImp struct {
	Place      string
	Importance int
}
*/

func NewTask(content string) (*Task, error) {
	if content == "" {
		return nil, errors.New("Non value")
	}
	return &Task{Id: 0, Content: content, Done: false}, nil
}

func init() {
	orm.RegisterModel(new(Task))

	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "dev_todo.db")

	err := orm.RunSyncdb("default", true, true)
	if err != nil {
		panic(err)
	}
}
