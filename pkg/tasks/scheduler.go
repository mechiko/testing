package tasks

import (
	"fmt"

	"testing/internal/entity"
)

type taskSheduler struct {
	tasks []entity.Task
	app   entity.App
}

var _ entity.TaskList = &taskSheduler{}

func NewTaskScheduler(a entity.App) entity.TaskList {
	return &taskSheduler{
		app: a,
	}
}

func (tl *taskSheduler) Add(t entity.Task) error {
	for _, o := range tl.tasks {
		if o == t {
			return fmt.Errorf("task allredy added")
		}
		if o.Info().Name == t.Info().Name {
			if o.Info().State == "создана" {
				return fmt.Errorf("задача такого же типа есть в статусе создана, дождитесь завершения задачи")
			}
		}
	}
	tl.tasks = append(tl.tasks, t)
	return nil
}

func (tl *taskSheduler) Del(t entity.Task) error {
	for i, o := range tl.tasks {
		if o == t {
			tl.tasks = append(tl.tasks[:i], tl.tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task not found")
}

func (tl *taskSheduler) Clear() {
	// newtasks := []entity.Task{}
	// for _, o := range tl.tasks {
	// 	newtasks = append(newtasks, o)
	// }
	// tl.tasks = newtasks
}

func (tl *taskSheduler) List() []entity.Task {
	tl.Clear()
	return tl.tasks
}

func (tl *taskSheduler) Load() error {
	defer tl.app.GetRecovery().RecoverLog("Load()")
	// if tasks, err := tl.app.GetRepo().GetTasks().GetActive(); err != nil {
	// 	return fmt.Errorf("%w", err)
	// } else {
	// 	for i, t := range tasks.Items {
	// 		fmt.Printf("i=%v t=%+v\n", i, t)
	// 		switch t.Name {
	// 		case "TaskRequest":
	// 			fmt.Printf("Create task name=%v with t=%+v\n", t.Name, t)
	// 			// if task, err := request_rest.Load(tl.app, t); err != nil {

	// 			// } else {
	// 			// 	if err := tl.Add(task); err != nil {
	// 			// 		return fmt.Errorf("%w", err)
	// 			// 	}
	// 			// }
	// 		}
	// 	}
	// }
	return nil
}
