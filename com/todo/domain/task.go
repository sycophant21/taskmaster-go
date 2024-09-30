package domain

import (
	"fmt"
	"go_test/com/todo/metadata"
	"strconv"
	"time"
)

const (
	Hourly TaskRepetitionType = iota
	Daily
	Weekly
	Monthly
	Yearly
	Custom
	// While choosing Custom TaskRepetitionType the user should be able to do one of the following
	// The user should be able to choose
	//which decade(s) of this century then
	//which year(s) of those decade(s) then
	//which month(s) of those year(s) then
	//which date(s)/week(s) of those months then
	//which day(s) of the week(s) then
	//which hour(s) of those day(s) then
	//which minute(s) of those hour(s)
	// they want the task to be repeated

)

var (
	Buy      = TaskType{index: 1, name: "Buy"}
	Sell     = TaskType{index: 2, name: "Sell"}
	Do       = TaskType{index: 3, name: "Do"}
	Cook     = TaskType{index: 4, name: "Cook"}
	Clean    = TaskType{index: 5, name: "Clean"}
	Read     = TaskType{index: 6, name: "Read"}
	Write    = TaskType{index: 7, name: "Write"}
	Drive    = TaskType{index: 8, name: "Drive"}
	Walk     = TaskType{index: 9, name: "Walk"}
	Shop     = TaskType{index: 10, name: "Shop"}
	Exercise = TaskType{index: 11, name: "Exercise"}
	Pay      = TaskType{index: 12, name: "Pay"}
	Sleep    = TaskType{index: 13, name: "Sleep"}
	Wake     = TaskType{index: 14, name: "Wake"}
	Plan     = TaskType{index: 15, name: "Plan"}
	Watch    = TaskType{index: 16, name: "Watch"}
	Call     = TaskType{index: 17, name: "Call"}
	Text     = TaskType{index: 18, name: "Text"}
	Organize = TaskType{index: 19, name: "Organize"}
	Repair   = TaskType{index: 20, name: "Repair"}
)
var (
	taskMap     map[int]Task
	taskTypeMap map[TaskType][]Task
	taskList    []Task
	lastTaskId  int
	taskTypes   []TaskType
)

func init() {
	taskMap = make(map[int]Task)
	taskTypeMap = make(map[TaskType][]Task)
	taskList = make([]Task, 0)
	lastTaskId = 1
	taskTypes = []TaskType{Buy, Sell, Do, Cook, Clean, Read, Write, Drive, Walk, Shop, Exercise, Pay, Sleep, Wake, Plan, Watch, Call, Text, Organize, Repair}
}

type Task struct {
	id           int
	taskType     TaskType
	description  string
	isCompleted  bool
	repetitive   bool
	desiredTime  *time.Time
	deadline     *time.Time
	metadata     metadata.Metadata
	reminderTime *time.Time
}

type TaskType struct {
	index int
	name  string
}

func (t TaskType) GetName() string {
	return t.name
}

func Empty() Task {
	return Task{}
}

func CreateNewTask(taskType TaskType, description string, repetitive bool, desiredTime *time.Time, deadline *time.Time, reminderTime *time.Time) Task {
	task := Task{
		id:           lastTaskId,
		taskType:     taskType,
		description:  description,
		isCompleted:  false,
		repetitive:   repetitive,
		desiredTime:  desiredTime,
		deadline:     deadline,
		metadata:     metadata.CreateNewMetadata(),
		reminderTime: reminderTime,
	}
	taskMap[lastTaskId] = task
	taskList = append(taskList, task)
	lastTaskId++
	return task
}

func GetTaskTypes() []TaskType {
	sliceCopy := make([]TaskType, len(taskTypes))
	copy(sliceCopy, taskTypes)
	return sliceCopy
}

func GetTasks() []Task {
	sliceCopy := make([]Task, len(taskList))
	copy(sliceCopy, taskList)
	return sliceCopy
}

func PrintTasks() {
	for i := 1; i <= len(taskList); i++ {
		taskList[i-1].print()
	}
}

func (t Task) print() {
	fmt.Println(strconv.Itoa(t.id)+".", t.taskType.name, t.description, "on", t.repetitive, "at", t.desiredTime, "or by", t.deadline, ".\nGet a reminder at", t.reminderTime)
}

func printTaskTypes() {
	fmt.Println("Select Task Type: ")
	for i := 0; i < len(taskTypes); i++ {
		fmt.Println(taskTypes[i].index, ". ", taskTypes[i].name)
	}
}

func EditTask() {
	PrintTasks()
	fmt.Println("Enter Task index to edit.")
	fmt.Println("Enter 0 to go back.")
	var input int8
	_, _ = fmt.Scanln(&input)
	for input < 0 || int(input) > len(taskList) {
		fmt.Println("Invalid input. Please enter a valid input.")
		_, _ = fmt.Scanln(&input)
	}
	if input == 0 {
		return
	} else {
		taskList[input-1].edit()
	}
}

func (t *Task) edit() {
	var changed bool
	t.description, changed = updateVal("Description", t.description)
	defer t.updateMetadata(changed, time.Now())
	/*	t.repetitive, changed = updateVal("Repetitive", t.repetitive)
		defer t.updateMetadata(changed, time.Now())
		var desiredTimeString string
		desiredTimeString, changed = updateVal("Desired Time", desiredTimeString)
		defer t.updateMetadata(changed, time.Now())
		t.deadline, changed = updateVal("Deadline", t.deadline)
		defer t.updateMetadata(changed, time.Now())
		t.reminderTime, changed = updateVal("Reminder Time", t.reminderTime)
		defer t.updateMetadata(changed, time.Now())*/
}

func updateVal[T taskFields](varName string, varVal T) (T, bool) {
	fmt.Println(varName, "Current value is", varVal)
	fmt.Println("Enter the new value ")
	var newVal T
	_, _ = fmt.Scanln(&newVal)
	return newVal, varVal != newVal
}

type taskFields interface {
	string | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func (t *Task) updateMetadata(change bool, newTime time.Time) {
	if change {
		t.metadata = metadata.CreateNewMetadataWithDetails(t.metadata.GetCreatedAt(), newTime)
	}
}

type TaskRepetitionType int

/*func SelectTaskType() TaskType {
	var selectedTaskType int8
	printTaskTypes()
	_, _ = fmt.Scanln(&selectedTaskType)
	for selectedTaskType < 1 || selectedTaskType >= int8(len(taskTypes)) {
		fmt.Println("Invalid Task Type. Please select a valid type.")
		_, _ = fmt.Scanln(&selectedTaskType)
	}
	return taskTypes[selectedTaskType-1]
}

func GetDescription() string {
	fmt.Println("Add Description (optional, Press Enter to skip):")
	des, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.TrimSpace(des)
}

func GetRepetition() bool {
	var isRepetitiveString string
	fmt.Println("Repeat Task (Enter 'Y' for Yes, 'N' for No.)?")
	_, _ = fmt.Scanln(&isRepetitiveString)
	for isRepetitiveString != "Y" && isRepetitiveString != "N" {
		fmt.Println("Invalid repetition status. Please enter a valid status.")
		_, _ = fmt.Scanln(&isRepetitiveString)
	}
	return isRepetitiveString == "Y"
}*/
