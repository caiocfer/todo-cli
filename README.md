# todo-cli

A small command-line TODO list written in Go. Tasks are stored in a local JSON file (`todolist.json`). The project exposes simple commands to add, list, complete and delete tasks using a Cobra-based CLI.

## Features

- Add tasks
- List tasks in a table view
- Mark tasks as completed
- Delete tasks
- Stores tasks in `todolist.json` in the repository root

## Requirements

- Go 1.18+ (tested with Go 1.25.3 in development)

## Quick start

Clone the repo and build:

```bash
git clone https://github.com/caiocfer/todo-cli.git
cd todo-cli
go build ./...
```

Run the CLI without building:

```bash
go run main.go <command>
```

Examples

- Add a task (the `add` command will prompt for a task name):

```bash
go run main.go add
```

- List tasks:

```bash
go run main.go list
```

- Mark a task completed:

```bash
go run main.go complete
```

- Delete a task:

```bash
go run main.go delete
```

- Show version:

```bash
go run main.go version
```

## Storage

Tasks are saved to a JSON file named `todolist.json` in the repository root by default (see `constants/constants.go`). Each task has this shape:

```json
{
  "id": 1,
  "name": "Some task",
  "completed": false
}
```

IDs are assigned when tasks are added; the code matches operations by task `Id` (not slice index), so IDs remain stable even after deletes.
