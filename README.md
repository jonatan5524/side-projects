# side-projects-manager

Side projects manager, manages all your side projects in one place

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
| `add-dir` | Adding directory of side projects. |
| `projects` | List all the side projects. |
| `dirs` | List all the directories. |
| `delete-dir` | Delete directory of side projects. |
