package task

import (
	"testing"
)

func TestCreateTask(t *testing.T) {
	taskName := "do a thing"
	task := CreateTask(1, taskName)

	if task == nil {
		t.Errorf("CreateTask() error")
	}
}
