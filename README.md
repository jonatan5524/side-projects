# side-projects-manager

Side projects manager, manages all your side projects in one place

## Requirments

Go version 1.17 or later

## Installation

Install using go

```bash
go get github.com/jonatan5524/side-projects-manager
```

Run all tests with coverage

```bash
go test ./... -cover
```

### GOPATH

Make sure your `PATH` includes the `$GOPATH/bin` directory so your commands can
be easily used:

```bash
export PATH=$PATH:$GOPATH/bin
```

## Command-Line

You can get all the possible functionaly using command line,

To see all commands available:

```bash
side-project-manager -h
```

To see all options available to a spasific command:

```bash
side-project-manager [command] --help
```

Available commands:
| command name | description |
|---|---|
| `add-dir [path]` | Adding directory of side projects. |
| `projects` | List all the side projects. |
| `dirs` | List all the directories. |
| `delete-dir [path]` | Delete directory of side projects from track (not from file system). |
| `delete-project [path]` | Delete side project from track (not from file system). |
| `project-info` | Delete directory of side projects. |
| `recent` | List recent side projects. |
