package main

import (
	"fmt"
	"orchestrator/manager"
	"orchestrator/node"
	"orchestrator/task"
	"orchestrator/worker"
	"time"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("program starting")

	task1 := task.Task{
		ID:     uuid.New(),
		Name:   "Task-1",
		State:  task.Pending,
		Image:  "Image-1",
		Memory: 1024,
		Disk:   1,
	}

	taskEvent := task.TaskEvent{
		ID:        uuid.New(),
		State:     task.Pending,
		TimeStamp: time.Now(),
		Task:      task1,
	}

	fmt.Printf("task: %v\n", task1)
	fmt.Printf("taskEvent: %v\n", taskEvent)

	worker := worker.Worker{
		Name:  "Worker1",
		Queue: *queue.New(),
		Db:    make(map[uuid.UUID]*task.Task),
	}

	fmt.Printf("worker: %v\n", worker)

	worker.CollectStats()
	worker.RunTask()
	worker.StartTask()
	worker.StopTask()

	manager := manager.Manager{
		PendingTasks: *queue.New(),
		TaskDb:       make(map[string][]*task.Task),
		EventDb:      make(map[string][]*task.TaskEvent),
		Workers:      []string{worker.Name},
	}

	fmt.Printf("manager: %v\n", manager)

	manager.SelectWorker()
	manager.UpdateTasks()
	manager.SendWork()

	node := node.Node{
		Name:   "Node-1",
		Ip:     "192.168.1.1",
		Cores:  4,
		Memory: 1024,
		Disk:   25,
		Role:   "worker",
	}

	fmt.Printf("Node: %v\n", node)

}
