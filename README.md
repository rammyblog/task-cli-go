# Task manager CLI

## Introduction

This is a simple task manager CLI, I built it to hone my GOlang skills.. Just played around with it, several improvements can be done like `adding a DB`, `sending a reminder when a task is due`, `filter tasks using the date`.

## Requirements

- [Go](https://golang.org) - v1.20 above

## Installation

- Clone this repo

  ```bash
  git clone https://github.com/rammyblog/task-cli-go
  ```

- Change directory to project directory

  ```bash
  cd task-cli-go
  ```

## Usage

- Add a new task

  ```bash
  go run main.go add -t "New Task" -d 2h
  ```

- View all task

  ```bash
  go run main.go view-all
  ```

- View a task

  ```bash
  go run main.go view -i "id"
  ```

- Mark a task as completed

```bash
  go run main.go complete -i "id"
```

- Delete a task

```bash
  go run main.go delete -i "id"
```

- Edit a task

```bash
go run main.go edit -t  "I fancy GO" -i "id"
```

## Contributors

| Contributor Name                                       | Role      | Tool                                    | Language(s) |
| ------------------------------------------------------ | --------- | --------------------------------------- | ----------- |
| [Babatunde Onasanya](https://twitter.com/simply_rammy) | Developer | [VSCode](https://code.visualstudio.com) | Go          |
