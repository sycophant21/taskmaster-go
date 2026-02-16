# taskmaster-go

A command-line task management and reminder application written in Go.

## Overview

taskmaster-go is a terminal-based to-do list application that helps you organize and track tasks. It provides an interactive menu system for creating, editing, viewing, and managing tasks with time-based reminders.

## Features

### Task Management
- **Create Tasks**: Add new tasks with detailed properties
- **Edit Tasks**: Modify task descriptions
- **View Tasks**: Display all tasks with their details
- **Delete Tasks**: Remove completed or unwanted tasks

### Task Properties
Each task includes:
- **Task Type**: Choose from 20 predefined categories (Buy, Sell, Do, Cook, Clean, Read, Write, Drive, Walk, Shop, Exercise, Pay, Sleep, Wake, Plan, Watch, Call, Text, Organize, Repair)
- **Description**: Brief text description (up to 100 characters)
- **Repetition Status**: Mark tasks as recurring (Yes/No)
- **Desired Time**: Preferred time to complete the task
- **Deadline**: Latest time the task must be completed
- **Reminder Time**: When to receive a reminder
- **Metadata**: Automatic tracking of creation and modification timestamps

### User Interface
- Time-of-day greetings (Good Morning/Afternoon/Evening)
- Interactive menu with numbered options
- Multiple time input formats:
  - 12-hour format (e.g., 11:59AM, 05:36PM)
  - 24-hour format (e.g., 11:59, 17:36)
  - Military time format (e.g., 1159, 1736)

## Project Structure

```
com/todo/
├── main.go                    # Application entry point, menu loop
├── domain/
│   └── task.go               # Task data structures and task types
├── metadata/
│   └── metadata.go           # Timestamp tracking for tasks
├── collection/
│   └── blocking_queue.go     # (Stub - not yet implemented)
└── utiils/
    ├── io/
    │   └── io_utils.go       # Input utilities for user interaction
    └── calendar/
        └── calendar.go       # (Stub - not yet implemented)
```

## Running the Application

```bash
go run com/todo/main.go
```

## Usage

Upon starting, you'll see a greeting based on the current time and a menu:

```
Hello!, Good Morning
How may I help you today?
Enter:
'1' to add new Task
'2' to edit a Task
'3' to delete a Task
'4' to see upcoming tasks
'0' to exit
```

### Creating a Task

1. Select option `1`
2. Choose a task type from the list (1-20)
3. Enter a description (optional, max 100 characters)
4. Specify if the task is repetitive (Y/N)
5. Enter desired time in your preferred format
6. Enter deadline time
7. Enter reminder time

### Viewing Tasks

Select option `4` to see all tasks. Each task displays:
- Task ID and type
- Description
- Repetition status
- Desired time and deadline
- Reminder time

### Editing Tasks

1. Select option `2`
2. View the list of tasks
3. Enter the task ID to edit (or 0 to go back)
4. Update the task description

## Input Utilities

The application provides helpful input functions:

- **Boolean Input**: Y/N prompts with validation
- **String Input**: Text input with optional character limit
- **Choice Input**: Select from numbered options
- **Time Input**: Parse time in multiple formats
- **Integer Input**: Numeric input with type safety

Special escape sequence: Type `\e` to exit string input early.

## Current State

The application currently supports:
- Creating tasks with full property sets
- Viewing all tasks in the task list
- Editing task descriptions
- Interactive menu navigation with input validation

**In Development**:
- Task deletion functionality (menu option exists but not implemented)
- Blocking queue for task scheduling (`blocking_queue.go` is empty)
- Calendar utilities (`calendar.go` is empty)
- Advanced repetition patterns (commented code suggests future custom repetition support: decade/year/month/week/day/hour/minute selection)

## Implementation Details

- Tasks are stored in memory using maps and slices
- Task IDs are auto-incremented starting from 1
- Task types are predefined as constants
- Metadata automatically tracks creation and modification times
- Time parsing supports three standard formats (12-hour, 24-hour, military)

## Requirements

- Go 1.23.1 or later (as specified in go.mod)
- No external dependencies (uses only Go standard library)

## Module Name

The project uses module name `go_test` (from go.mod).
