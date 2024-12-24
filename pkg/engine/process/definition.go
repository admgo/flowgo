package process

// pkg/engine/process/definition.go

import (
	"flowgo/pkg/engine/task"
	"fmt"
)

// ProcessDefinition 定义流程结构
type ProcessDefinition struct {
	ID   string
	Name string
}

// ProcessInstance 表示流程实例
type ProcessInstance struct {
	ID     string
	Status string
}

// Engine 用于管理流程定义和实例
type Engine struct {
	definitions map[string]ProcessDefinition
	instances   map[string]ProcessInstance
	taskManager *task.TaskManager
}

// NewEngine will create a engine
func NewEngine() *Engine {
	return &Engine{
		definitions: make(map[string]ProcessDefinition),
		instances:   make(map[string]ProcessInstance),
		taskManager: task.NewTaskManager(),
	}
}

// AddDefinition add a process definition
func (e *Engine) AddDefinition(id, name string) {
	e.definitions[id] = ProcessDefinition{ID: id, Name: name}
	fmt.Printf("Added process definition: %s\n", name)
}

// StartInstance start a process instance
func (e *Engine) StartInstance(definitionID string) string {
	instanceID := fmt.Sprintf("instance-%d", len(e.instances)+1)
	e.instances[instanceID] = ProcessInstance{ID: instanceID, Status: "running"}
	fmt.Printf("Started process instance: %s for definition ID: %s\n", instanceID, definitionID)

	// Adding a new task to the instance to demonstrate task management
	taskID := fmt.Sprintf("task-%d", len(e.instances))
	e.taskManager.AddTask(taskID, "Sample Task for Instance "+instanceID)
	return instanceID
}
