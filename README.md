# Go-Task-Manager

The exercise from https://gophercises.com/exercises/

```
$ go-task-manager
go-task-manager is a CLI for managing your TODOs.

Usage:
  go-task-manager [command]

Available Commands:
  add         Add a new task to your TODO list
  do          Mark a task on your TODO list as complete
  list        List all of your incomplete tasks

Use "go-task-manager [command] --help" for more information about a command.

$ go-task-manager add review talk proposal
Added "review talk proposal" to your task list.

$ go-task-manager add clean dishes
Added "clean dishes" to your task list.

$ go-task-manager list
You have the following tasks:
1. review talk proposal
2. some task description

$ go-task-manager do 1
You have completed the "review talk proposal" task.

$ go-task-manager list
You have the following tasks:
1. some task description
```