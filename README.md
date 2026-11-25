# Basic Go Architecture

### File Structure

```
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
```


### Github

1. Create a repository on github.com (hit the '+' icon or green 'New' button)
2. Run git commands in command line
    - git init (this command is to setup a git repo locally, only run once then skip this step)
    - git add .
    - git commit -m "init" (only use "init" as the inital commit message, any messages after should be more descriptive. Can include notes on what progress has been made)
    - git remote add origin <github_repo_url> (this command is to setup the remote, only run once then skip this step)
    - git push <remote> <branch_name>
       e.g: git push origin main 



### Command line commands

```
Make new folder(directory): mkdir <new_dir_name>
    e.g: mkdir internal/todos
Create new file: type nul > <file_location_or_name>
    e.g: type nul > cmd/main.go
Start a new Go Module: go mod init <project_name>
```

-----------

### PROMPT



You are a Go software developer. Please build a command-line interface (CLI) Todo application written in Go.

The application should be interactive and meet the following requirements:

1. Initial Startup & User Greeting

When the application starts, it should prompt the user to enter their name (e.g., "Welcome! Please enter your name:").

Once entered, greet the user by name and start the main interactive menu.

2. Data Structure & Features

Multiple Lists: A user should be able to create, view, and manage multiple todo lists (e.g., "Work", "Personal", "Groceries").

Todo Items: Each list contains multiple todo items. Every todo item must have:

A description (string).

A status indicating whether it has been completed or not (boolean).

3. CLI Navigation & Commands

After starting, all inputs must be taken dynamically through the command line (using bufio.Scanner or fmt.Scan).

Provide a clear menu loop with options such as:

Create a new todo list.

View all todo lists.

Select a list to view/manage its todos.

Add a new todo to a specific list.

Mark a todo as done/undone.

Delete a todo or a list.

Exit the application.

4. Code Quality & Best Practices

Use clean, idiomatic Go code (proper structs, slices, packages if necessary, and error handling).

Structure the code into logical files (e.g., main.go, models.go, app.go) to keep it maintainable.

Include clear comments and a README.md explaining how to run the application.