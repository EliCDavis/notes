# Notes

Structured notes

## Commands

### Project

#### Create

```bash
NAME:
   notes project new - Creates a new project

USAGE:
   notes project new [command options] [project name]

OPTIONS:
   --path value  path to place project (default: "./")
   --help, -h    show help
```

Example: 

```bash 
notes project new "My Project"
```

#### Compile

```bash
NAME:
   notes project compile - compiles project into a single file

USAGE:
   notes project compile [command options]

OPTIONS:
   --project value  Path to project.json (default: "./project.json")
   --help, -h       show help
```

Example: 

```bash 
notes project compile
```

### Log

#### New

```bash
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

### Task

#### New

```bash
NAME:
   notes task new - Creates a new task

USAGE:
   notes task new [command options]

OPTIONS:
   --name value
   --help, -h    show help
```

Example: 

```bash 
notes task new --name "My Task"
```