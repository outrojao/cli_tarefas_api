package database

import (
	"cli_tasks_api/internal/app/task"
)

func CreateTask(taskName string) (int, error) {
	var id int
	query := `INSERT INTO tasks (name) VALUES ($1) RETURNING id;`
	err := DB.QueryRow(query, taskName).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetTaskByName(name string) (task.Task, error) {
	var t task.Task
	query := `SELECT id, name, done FROM tasks WHERE name = $1 LIMIT 1;`
	err := DB.QueryRow(query, name).Scan(&t.Id, &t.Name, &t.Done)
	if err != nil {
		return task.Task{}, err
	}
	return t, nil
}

func UpdateTaskStatus(id int, done bool) error {
	query := `UPDATE tasks SET done = $1 WHERE id = $2;`
	_, err := DB.Exec(query, done, id)
	return err
}

func DeleteTask(id int) error {
	query := `DELETE FROM tasks WHERE id = $1;`
	_, err := DB.Exec(query, id)
	return err
}

func GetAllTasks() ([]task.Task, error) {
	query := `SELECT id, name, done FROM tasks;`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []task.Task
	for rows.Next() {
		var t task.Task
		if err := rows.Scan(&t.Id, &t.Name, &t.Done); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}
