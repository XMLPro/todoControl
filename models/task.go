package models

type Task struct {
	Id         int    `json:"id"`
	Content    string `json:"content"`
	LimitTime  string `json:"limit_time"`
	Workload   int    `json:"workload"`
	PlaceId    int    `json:"place_id"`
	Importance int    `json:"importance"`
}
