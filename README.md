# vsinit
Quick initialization of go workspaces for visual studio code (opinionated)

Very early hack, use at your own risk!

# functionality

## generate .vscode/tasks.json 
The tool generates a tasks.json that contains a "run" job that builds and runs the project from the current main folder.

Build and run was chosen over "go run ." to avoid the firewall asking for every new run under windows

