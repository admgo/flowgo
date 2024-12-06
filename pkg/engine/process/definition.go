package process

// pkg/engine/process/definition.go

import (
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
}

// NewEngine 创建一个新的流程引擎
func NewEngine() *Engine {
	return &Engine{
		definitions: make(map[string]ProcessDefinition),
		instances:   make(map[string]ProcessInstance),
	}
}

// AddDefinition 添加一个新的流程定义
func (e *Engine) AddDefinition(id, name string) {
	e.definitions[id] = ProcessDefinition{ID: id, Name: name}
	fmt.Printf("Added process definition: %s\n", name)
}

// StartInstance 启动一个流程实例
func (e *Engine) StartInstance(definitionID string) string {
	instanceID := fmt.Sprintf("instance-%d", len(e.instances)+1)
	e.instances[instanceID] = ProcessInstance{ID: instanceID, Status: "running"}
	fmt.Printf("Started process instance: %s for definition ID: %s\n", instanceID, definitionID)
	return instanceID
}
