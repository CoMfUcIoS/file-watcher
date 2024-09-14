# File Watcher CLI Tool

This CLI tool watches a specified file or directory for changes and executes a given command whenever a change is detected. It uses the `fsnotify` package to monitor file system events and supports debouncing to prevent multiple rapid executions of the command.

## Installation

To install the tool, you need to have Go installed on your system. Then, you can build the tool using the following commands:

```sh
git clone git@github.com/comfucios/file-watcher
cd file-watcher
go build -o file-watcher
```

## Usage

The tool requires two flags: `--source` and `--action`.

- `--source`: The source file or directory to watch.
- `--action`: The CLI command to execute when a change is detected.

### Example

Watch a directory and execute a command when a change is detected:

```sh
./file-watcher --source /path/to/directory --action "echo 'Directory changed!'"
```

Watch a file and execute a command when a change is detected:

```sh
./file-watcher --source /path/to/file.txt --action "echo 'File changed!'"
```

### Notes

- The tool supports debouncing with a 500-millisecond delay to prevent multiple rapid executions of the command.
- The tool will monitor the specified directory and all its subdirectories.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

```

```
