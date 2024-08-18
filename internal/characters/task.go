package characters

type TaskCode = string
type TaskType = int
type TaskProgress = int
type TaskTotal = int

const (
	TaskTypeMonsters TaskType = iota
	TaskTypeResources
	TaskTypeCrafts
)

type Task struct {
	code  TaskCode
	type_ TaskType
	total TaskTotal
}

func NewTask(code TaskCode, type_ TaskType, total TaskTotal) *Task {
	return &Task{code, type_, total}
}

func (t *Task) Code() TaskCode {
	return t.code
}

func (t *Task) Type() TaskType {
	return t.type_
}

func (t *Task) Total() TaskTotal {
	return t.total
}
