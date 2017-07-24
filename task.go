package main

type Task struct {
	//primaryKey int
	Content    string `json:"content"`
	Deadline   string `json:"deadline"`
	From       string `json:"from"`
	Workload   string `json:"workload"`
	Importance string `json:"importance"`
}
