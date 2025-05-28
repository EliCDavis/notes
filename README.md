# Notes

Opinionated note taking.

## Install

```
go install ./cmd/notes
```

## Example

```console
foo@bar:~$ notes project new "My Example Project"
foo@bar:~$ cd MyExampleProject
foo@bar:MyExampleProject$ notes log new # Create a new daily log
foo@bar:MyExampleProject$ notes meeting new # Create meeting notes
foo@bar:MyExampleProject$ notes topic new "Some Topic" # Create a new topic
foo@bar:MyExampleProject$ notes task new "Task to Complete" # Create a new task
foo@bar:MyExampleProject$ notes task start 1 # Start a task
foo@bar:MyExampleProject$ notes task stop  1 # Stop a task
foo@bar:MyExampleProject$ notes task complete 1  # Make a task as complete
foo@bar:MyExampleProject$ notes task abandon 1 "Reason for abdonment" # Abandon a certain task
foo@bar:MyExampleProject$ notes project compile # Compile notes to single markdown
```

## Commands

```
NAME:
   notes - Manage a collection of notes

USAGE:
   notes [global options] command [command options]

AUTHOR:
   Eli C Davis

COMMANDS:
   project  Project management functionality
   log      Create and edit logs
   task     Create and edit tasks
   meeting  Create and edit meetings
   topic    Create and edit topics
   image    Manage images
   tag      Manage tags
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

### Project

#### Create

```
NAME:
   notes project new - Creates a new project

USAGE:
   notes project new [command options] [project name]

OPTIONS:
   --path value      path to place project (default: "./")
   --logs value      path to logs (default: "logs")
   --tasks value     path to tasks (default: "tasks")
   --meetings value  path to meetings (default: "meetings")
   --topics value    path to topics (default: "topics")
   --help, -h        show help
```

Example: 

```bash 
notes project new "My Project"
```

#### Compile

```
NAME:
   notes project compile - compiles project into a single file

USAGE:
   notes project compile [command options]

OPTIONS:
   --project value  Path to project.json (default: "./project.json")
   --use-includes   Whether or not to use the markdown-it-include extension (default: false)
   --help, -h       show help
```

Example: 

```bash 
notes project compile
```

### Log

#### New

```
NAME:
   notes log new - Creates a new log

USAGE:
   notes log new [command options]

OPTIONS:
   --help, -h  show help
```

Example: 

```bash 
notes log new
```

#### Open

```
NAME:
   notes log open - Opens the current log for the day, or creates one if it hasn't been created yet

USAGE:
   notes log open [command options]

OPTIONS:
   --help, -h  show help
```

Example:

```bash
notes log open
```

### Task

#### New

```
NAME:
   notes task new - Creates a new task

USAGE:
   notes task new [command options] [Task Name]

OPTIONS:
   --name value
   --help, -h    show help
```

Example: 

```bash 
notes task new "My Task"
```

#### Update

```
NAME:
   notes task update - update a task

USAGE:
   notes task update [command options]

COMMANDS:
   name     Update the name of the task
   status   Update the status of the task
   help, h  Shows a list of commands or help for one command

OPTIONS:
   --task value  ID of task to update (default: 0)
   --help, -h    show help
```

Example: 

```bash 
notes task update --task 1 name "New Task Name"
```

#### List

```
NAME:
   notes task list - lists all tasks

USAGE:
   notes task list [command options]

OPTIONS:
   --help, -h  show help
```

Example: 

```bash 
notes task list
```

#### Gantt

```
NAME:
   notes task gantt - create a gantt chart in mermaid

USAGE:
   notes task gantt [command options]

OPTIONS:
   --help, -h  show help
```

Example: 

```bash 
notes task gantt
```


### Image

Manage images to be used across your project.

#### Add

```
NAME:
   notes image add - Adds an image to the project

USAGE:
   notes image add [command options] [path to images]

OPTIONS:
   --help, -h  show help
```

Example: 

```bash 
notes image add "image 1.png" "image 2.png"
```

### Tag

#### New

```
NAME:
   notes tag new - Creates tags for the project

USAGE:
   notes tag new [command options] [name of tags]

OPTIONS:
   --help, -h  show help
```

Example

```bash
notes tag new "Setup" "Research"
```