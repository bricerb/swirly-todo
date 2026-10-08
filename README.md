# Basic Go Architecture

### File Structure

Project folder:
    - root(entire project folder)    
        - cmd
          - main.go
        - internal (src/name_of_project)
          - io (input/output)
            - input.go (where functions to read input from user, and any structures asssociated with these functions)
          - todos
            - models.go (type todo struct, type todoList struct, and any other structs)
            - todo.go (functions for adding, removing, checking off todos)
            - list.go (functions for creating new list of todos, editing list of todos, or deleting list of todos)
          - server/mux/http
            - server.go (start server)
            - routes.go (routing)
          - utils
            - math.go (addition, substraction)
            - 
          - errors
            - models.go (error struct)
            - errors.go (logError, todoError)
        - README.md
        - go.mod
        - go.*



### Github

1. Create a repository on github.com (hit the '+' icon or green 'New' button)
2. Run git commands in command line
    - git init
    - git add .
    - git commit -m "init"
    - git remote add origin <github_repo_url>
    - git push <remote> <branch_name>



### Command line commands

Make new folder(directory): mkdir <new_dir_name>
    e.g: mkdir internal/todos
Create new file: type nul > <file_location_or_name>
    e.g: type nul > cmd/main.go
Start a new Go Module: go mod init <project_name>



