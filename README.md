# goinit
Quick initialization of go workspaces (very opinionated)

Very early hack, use at your own risk!

# usage
``` goinit <module name>```

# functionality

## generate directory
uses the last part of the module name as directory , e.g. "goinit github.com/hborntraeger/goinit" will generate a directory "goinit"

## generate .vscode/tasks.json 
The tool generates a tasks.json that contains a "run" job that builds and runs the project from the current main folder.

Build and run was chosen over "go run ." to avoid the firewall asking for every new run under windows

## init module 
runs ```go mod init <module name>```

## create git repository
runs ``` git init ```