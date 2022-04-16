package domain


type Task struct {
	ID int `json:"ID"`
	Name string `json:"Name"`
	Content string `json:"Content"`
}

type AllTasks []Task

var Tasks = AllTasks{
	{
		ID:1,
		Name: "Task one",
		Content: "Some Content",
	},
}