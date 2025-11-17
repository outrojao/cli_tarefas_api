package task

type Task struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

func CreateTask(id int, taskName string) *Task {
	return &Task{
		Id:   id,
		Name: taskName,
		Done: false,
	}
}
