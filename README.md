# File Monitoring System

File Monitoring System is a Go application that monitors a specified directory for file changes and records information about the modified files into a SQLite database.

## Features

- Watches a directory for file creation and modification events.
- Records file information (filename and size) into a SQLite database.
- Supports concurrent file processing using worker goroutines.

## Requirements

- Go 1.16 or later
- SQLite 3.x

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/yourusername/file-monitoring.git
2. cd file-monitoring
3. go build
4. ./file-monitoring

## Configuration



