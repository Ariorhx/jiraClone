package Tests

import (
	"JiraClone/CannotResolve/Classes"
	"testing"
)

func TestEmptyTaskPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	task := Classes.Task{}
	task.IdFromValue()
}

func TestTaskConvertValuesToIds(t *testing.T) {
	task := Classes.Task{User: "Оля", Priority: "критический"}
	task.IdFromValue()

	if task.UserId != 1 {
		t.Errorf("UserId must be 1, but it is %d\n", task.UserId)
	}
	if task.PriorityId != 4 {
		t.Errorf("PriorityId must be 4, but it is %d\n", task.PriorityId)
	}
}
