package main

import (
	"errors"
	"fmt"
	"go_test/com/todo/domain"
	"go_test/com/todo/utiils/io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	isFirst := true
	for {
		if isFirst {
			greet()
			isFirst = false
		}
		err := executeIntent(getIntent())
		if err != nil {
			log.Fatal(err)
		}
	}
}

func greet() {
	var currTime = time.Now()
	var hour = currTime.Hour()
	if hour >= 4 && hour < 12 {
		fmt.Println("Hello!, Good Morning")
	} else if hour >= 12 && hour < 16 {
		fmt.Println("Hello!, Good Afternoon")
	} else if hour >= 16 && hour < 21 {
		fmt.Println("Hello!, Good Evening")
	} else if hour >= 21 || hour < 4 {
		fmt.Println("Hello!, Good Morning")
	}
	fmt.Println("How may I help you today?")
}

func getIntent() uint8 {
	var input uint8
	fmt.Println("Enter: ")
	fmt.Println("'1' to add new Task")
	fmt.Println("'2' to edit a Task")
	fmt.Println("'3' to delete a Task")
	fmt.Println("'4' to see upcoming tasks")
	fmt.Println("'0' to exit")
	_, _ = fmt.Scanln(&input)
	return input
}

func executeIntent(intent uint8) error {
	switch intent {
	case 0:
		os.Exit(0)
	case 1:
		err := createNewTask()
		if err != nil {
			return err
		}
		clearTerminal()
	case 2:
		domain.EditTask()
	case 3:
	case 4:
		domain.PrintTasks()
	default:
		return errors.New("unknown intent")
	}
	return nil
}

func createNewTask() error {
	taskTypesString := make([]string, 0)
	taskTypes := domain.GetTaskTypes()
	for i := range taskTypes {
		taskTypesString = append(taskTypesString, taskTypes[i].GetName())
	}
	taskTypeIndex, err := io.ChoiceInput("Task Type", taskTypesString)
	if err != nil {
		return err
	}
	var selectedTaskType = domain.GetTaskTypes()[*taskTypeIndex]
	description, err := io.StringInputOfLength("Description", true, 100)
	if err != nil {
		return err
	}
	repetitionStatus, err := io.BoolInput("Repetition Status")
	if err != nil {
		return err
	}
	desiredTime, err := io.TimeInput("Desired Time", io.MilitaryFormat)
	if err != nil {
		return err
	}
	deadline, err := io.TimeInput("Deadline Time", io.MilitaryFormat)
	if err != nil {
		return err
	}
	reminderTime, err := io.TimeInput("Reminder Time", io.MilitaryFormat)
	if err != nil {
		return err
	}
	domain.CreateNewTask(selectedTaskType, *description, *repetitionStatus, desiredTime, deadline, reminderTime)
	return nil
}

func clearTerminal() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux", "darwin": // Unix-based OS like Linux or macOS
		cmd = exec.Command("clear")
	case "windows": // Windows OS
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		fmt.Println("Unsupported platform")
		return
	}

	// Attach the command's output to the terminal
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}
