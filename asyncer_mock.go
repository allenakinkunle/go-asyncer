package asyncer

import "time"

type MockAsyncer struct {
}

func NewMockAsyncer() *MockAsyncer {
	return &MockAsyncer{}
}

func (m MockAsyncer) EnqueueTask(taskName, taskID string, payload []byte) error {
	return nil
}

func (m MockAsyncer) ScheduleTask(taskName, taskID string, payload []byte, in time.Duration) error {
	return nil
}
