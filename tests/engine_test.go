package tests

import (
	"flowgo/pkg/engine/process"
	"testing"
)

func TestEngine(t *testing.T) {
	engine := process.NewEngine()

	// 添加和启动流程定义
	engine.AddDefinition("process_1", "Sample Process 1")
	instanceID := engine.StartInstance("process_1")

	if instanceID == "" {
		t.Errorf("Failed to start process instance")
	}
}
